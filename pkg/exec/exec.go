package exec

import (
	"bytes"
	"io"
	"os"
	"os/exec"
)

//go:generate mockgen -destination=../../mocks/mock_exec.go -package=mocks github.com/yardbirdsax/ensure-tfenv-versions/pkg/exec Exec
type Exec interface {
	ExecCommand(string, bool, ...string) (string, error)
}

type Executor struct{}

func NewExecutor() *Executor {
	return &Executor{}
}

func (*Executor) ExecCommand(command string, writeToConsole bool, args ...string) (output string, err error) {
	var buffer bytes.Buffer
	stdOutWriters := []io.Writer{&buffer}
	stdErrWriters := []io.Writer{&buffer}
	if writeToConsole {
		stdErrWriters = append(stdErrWriters, os.Stderr)
	}
	stdOutWriters = append(stdOutWriters, os.Stdout)
	stdOutW := io.MultiWriter(stdOutWriters...)
	stdErrW := io.MultiWriter(stdErrWriters...)

	cmd := exec.Command(command, args...)
	cmd.Stdout = stdOutW
	cmd.Stdin = os.Stdin
	cmd.Stderr = stdErrW

	err = cmd.Run()

	output = buffer.String()
	return
}
