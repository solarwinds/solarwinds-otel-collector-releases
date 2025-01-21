package cli

import (
	"fmt"
)

type BashCli struct{}

var _ (CommandLineExecutor) = (*BashCli)(nil)

func NewBashCliExecutor() CommandLineExecutor {
	return new(BashCli)
}

func (*BashCli) ExecuteCommand(command ...string) (stdOut, stdErr string, error error) {
	internalCommands := append([]string{"-c"}, command...)
	stdOut, stdErr, err := executeCommand(
		"bash",
		internalCommands...)
	if err != nil {
		message := fmt.Sprintf("Bash command %s failed.", command)
		return "", "", fmt.Errorf(message, err)
	}

	return stdOut, stdErr, nil
}
