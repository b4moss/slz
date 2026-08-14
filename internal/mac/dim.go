package mac

import (
	"fmt"
	"os"
	"os/exec"
)

type Cmd interface {
	Start() error
	Run() error
	Wait() error
}

type Runner interface {
	Command(name string, args ...string) Cmd
}

type execRunner struct{}

func NewRunner() Runner {
	return execRunner{}
}

type execCmd struct {
	cmd *exec.Cmd
}

func (execRunner) Command(name string, args ...string) Cmd {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	return execCmd{cmd: cmd}
}

func (c execCmd) Start() error { return c.cmd.Start() }
func (c execCmd) Run() error   { return c.cmd.Run() }
func (c execCmd) Wait() error  { return c.cmd.Wait() }

func Dim(goos string, runner Runner) error {
	if goos != "darwin" {
		return fmt.Errorf("mac dim is only supported on macOS")
	}
	caffeinate := runner.Command("caffeinate")
	if err := caffeinate.Start(); err != nil {
		return err
	}
	if err := runner.Command("pmset", "displaysleepnow").Run(); err != nil {
		return err
	}
	return caffeinate.Wait()
}
