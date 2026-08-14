package cli

import (
	"bytes"
	"testing"

	"github.com/b4moss/slz/internal/mac"
)

func runCLI(t *testing.T, opts Options, args ...string) (stdout, stderr string, err error) {
	t.Helper()
	cmd := NewRoot(opts)
	outBuf := new(bytes.Buffer)
	errBuf := new(bytes.Buffer)
	cmd.SetOut(outBuf)
	cmd.SetErr(errBuf)
	cmd.SetArgs(args)
	err = cmd.Execute()
	return outBuf.String(), errBuf.String(), err
}

type stubCmd struct{}

func (stubCmd) Start() error { return nil }
func (stubCmd) Run() error   { return nil }
func (stubCmd) Wait() error  { return nil }

type countingRunner struct {
	n int
}

func (r *countingRunner) Command(name string, args ...string) mac.Cmd {
	r.n++
	return stubCmd{}
}
