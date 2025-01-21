package cli

type CommandLineExecutorMock struct {
	stdout string
	stderr string
	err    error
}

var _ CommandLineExecutor = (*CommandLineExecutorMock)(nil)

func CreateNewCliExecutorMock(
	stdout string,
	stderr string,
	err error,
) CommandLineExecutor {
	return &CommandLineExecutorMock{
		stdout: stdout,
		stderr: stderr,
		err:    err,
	}
}

// ExecuteCommand implements CommandLineExecutor.
func (c *CommandLineExecutorMock) ExecuteCommand(
	...string,
) (stdOut string, stdErr string, error error) {
	return c.stdout, c.stderr, c.err
}
