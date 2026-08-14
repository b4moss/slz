package cli

import (
	"errors"
	"testing"

	"github.com/b4moss/slz/internal/mac"
)

type failStartCmd struct{}

func (failStartCmd) Start() error { return errors.New("caffeinate failed") }
func (failStartCmd) Run() error   { return nil }
func (failStartCmd) Wait() error  { return nil }

type failStartRunner struct{}

func (failStartRunner) Command(name string, args ...string) mac.Cmd {
	return failStartCmd{}
}

func TestMacDimSuccess(t *testing.T) {
	r := &countingRunner{}
	_, _, err := runCLI(t, Options{GOOS: "darwin", Runner: r}, "mac", "dim")
	if err != nil {
		t.Fatalf("expected success, got %v", err)
	}
	if r.n == 0 {
		t.Fatal("expected caffeinate/pmset to be invoked")
	}
}

func TestMacDimNonDarwin(t *testing.T) {
	r := &countingRunner{}
	_, _, err := runCLI(t, Options{GOOS: "linux", Runner: r}, "mac", "dim")
	if err == nil {
		t.Fatal("expected non-zero exit on non-darwin")
	}
	if r.n != 0 {
		t.Fatalf("non-darwin should not exec, got %d calls", r.n)
	}
}

func TestMacDimCaffeinateFailure(t *testing.T) {
	_, _, err := runCLI(t, Options{GOOS: "darwin", Runner: failStartRunner{}}, "mac", "dim")
	if err == nil {
		t.Fatal("expected non-zero exit when caffeinate fails")
	}
}

func TestMacDimRejectsExtraArgs(t *testing.T) {
	r := &countingRunner{}
	_, _, err := runCLI(t, Options{GOOS: "darwin", Runner: r}, "mac", "dim", "extra")
	if err == nil {
		t.Fatal("expected non-zero exit for extra args")
	}
}

func TestMacUnknownSubcommand(t *testing.T) {
	_, _, err := runCLI(t, Options{GOOS: "darwin", Runner: &countingRunner{}}, "mac", "nosuch")
	if err == nil {
		t.Fatal("expected non-zero exit for unknown mac subcommand")
	}
}
