package cli

import "fmt"

type PowershellCli struct{}

var _ (CommandLineExecutor) = (*PowershellCli)(nil)

func (*PowershellCli) ExecuteCommand(command ...string) (stdOut, strErr string, error error) {
	stdOut, stdErr, err := executeCommand(
		"powershell",
		command...,
	)
	if err != nil {
		message := fmt.Sprintf("Powershell command %s failed.", command)
		return "", "", fmt.Errorf(message, err)
	}

	return stdOut, stdErr, nil
}
