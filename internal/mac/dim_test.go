package mac

import (
	"errors"
	"strings"
	"sync"
	"testing"
	"time"
)

type recCmd struct {
	r    *recRunner
	name string
	args []string
}

func (c *recCmd) label(method string) string {
	parts := append([]string{method, c.name}, c.args...)
	return strings.Join(parts, " ")
}

func (c *recCmd) Start() error {
	c.r.append(c.label("Start"))
	return c.r.startErr
}

func (c *recCmd) Run() error {
	c.r.append(c.label("Run"))
	return c.r.runErr
}

func (c *recCmd) Wait() error {
	c.r.append(c.label("Wait"))
	if c.r.enteredWait != nil {
		select {
		case <-c.r.enteredWait:
		default:
			close(c.r.enteredWait)
		}
	}
	if c.r.blockWait != nil {
		<-c.r.blockWait
	}
	return c.r.waitErr
}

type recRunner struct {
	mu          sync.Mutex
	ops         []string
	startErr    error
	runErr      error
	waitErr     error
	blockWait   chan struct{}
	enteredWait chan struct{}
}

func (r *recRunner) Command(name string, args ...string) Cmd {
	return &recCmd{r: r, name: name, args: args}
}

func (r *recRunner) append(op string) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.ops = append(r.ops, op)
}

func (r *recRunner) snapshot() []string {
	r.mu.Lock()
	defer r.mu.Unlock()
	out := make([]string, len(r.ops))
	copy(out, r.ops)
	return out
}

func TestDimCallsCaffeinateThenPmset(t *testing.T) {
	r := &recRunner{}
	if err := Dim("darwin", r); err != nil {
		t.Fatalf("expected success, got %v", err)
	}
	got := strings.Join(r.snapshot(), " | ")
	want := "Start caffeinate | Run pmset displaysleepnow | Wait caffeinate"
	if got != want {
		t.Fatalf("ops = %q, want %q", got, want)
	}
}

func TestDimSuccess(t *testing.T) {
	if err := Dim("darwin", &recRunner{}); err != nil {
		t.Fatalf("expected success, got %v", err)
	}
}

func TestDimWaitsForCaffeinate(t *testing.T) {
	r := &recRunner{
		blockWait:   make(chan struct{}),
		enteredWait: make(chan struct{}),
	}
	done := make(chan error, 1)
	go func() {
		done <- Dim("darwin", r)
	}()

	select {
	case <-r.enteredWait:
	case <-time.After(time.Second):
		t.Fatal("Dim did not wait for caffeinate")
	}

	select {
	case err := <-done:
		t.Fatalf("Dim returned before caffeinate finished: %v", err)
	case <-time.After(30 * time.Millisecond):
	}

	close(r.blockWait)
	if err := <-done; err != nil {
		t.Fatalf("expected success, got %v", err)
	}
}

func TestDimNonDarwinDoesNotExec(t *testing.T) {
	r := &recRunner{}
	if err := Dim("linux", r); err == nil {
		t.Fatal("expected error on non-darwin")
	}
	if ops := r.snapshot(); len(ops) != 0 {
		t.Fatalf("non-darwin should not exec, got %v", ops)
	}
}

func TestDimCaffeinateFailure(t *testing.T) {
	r := &recRunner{startErr: errors.New("caffeinate failed")}
	if err := Dim("darwin", r); err == nil {
		t.Fatal("expected error when caffeinate fails")
	}
	for _, op := range r.snapshot() {
		if strings.Contains(op, "pmset") {
			t.Fatalf("pmset should not run after caffeinate failure, got %v", r.snapshot())
		}
	}
}
