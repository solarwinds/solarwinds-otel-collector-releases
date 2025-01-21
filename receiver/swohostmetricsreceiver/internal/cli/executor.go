package cli

import (
	"bytes"
	"os/exec"
)

type CommandLineExecutor interface {
	ExecuteCommand(command ...string) (stdOut, stdErr string, error error)
}

// executes commands given by method arguments and returns tuple of stringified standard output,
// with stringified standard error output and error if command fails.
func executeCommand(
	command string,
	args ...string,
) (stdOut, stdErr string, error error) {
	cmd := exec.Command(command, args...)

	var stdout bytes.Buffer
	var stderr bytes.Buffer

	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	return stdout.String(), stderr.String(), err
}
