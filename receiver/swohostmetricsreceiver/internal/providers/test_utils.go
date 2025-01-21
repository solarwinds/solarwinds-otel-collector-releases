package providers

import (
	"strings"

	"github.com/solarwinds/solarwinds-otel-collector/receiver/swohostmetricsreceiver/internal/cli"
)

type cliExecutorMock struct {
	stdout string
	stderr string
	err    error
}

var _ cli.CommandLineExecutor = (*cliExecutorMock)(nil)

func CreateCommandLineExecutorMock(
	stdout string,
	stderr string,
	err error,
) cli.CommandLineExecutor {
	return &cliExecutorMock{
		stdout: stdout,
		stderr: stderr,
		err:    err,
	}
}

// ExecuteCommand implements cli.CommandLineExecutor.
func (e *cliExecutorMock) ExecuteCommand(_ ...string) (stdOut string, stdErr string, error error) {
	return e.stdout, e.stderr, e.err
}

type CommandResult struct {
	Stdout string
	Stderr string
	Err    error
}

type cliMultiCommandExecutorMock struct {
	commandSet map[string]CommandResult
}

var _ cli.CommandLineExecutor = (*cliMultiCommandExecutorMock)(nil)

func CreateMultiCommandExecutorMock(commandsWithResults map[string]CommandResult) cli.CommandLineExecutor {
	return &cliMultiCommandExecutorMock{
		commandSet: commandsWithResults,
	}
}

// ExecuteCommand implements cli.CommandLineExecutor.
func (m *cliMultiCommandExecutorMock) ExecuteCommand(command ...string) (stdout string, stderr string, error error) {
	combinedCommand := strings.Join(command, " ")
	if v, found := m.commandSet[combinedCommand]; found {
		return v.Stdout, v.Stderr, v.Err
	}
	return "", "", nil
}
