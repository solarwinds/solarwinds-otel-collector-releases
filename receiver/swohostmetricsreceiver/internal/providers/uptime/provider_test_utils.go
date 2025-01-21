package uptime

type succeedingUptimeWrapper struct {
	uptime uint64
}

var _ Wrapper = (*succeedingUptimeWrapper)(nil)

func CreateSuceedingUptimeWrapper(
	uptime uint64,
) Wrapper {
	return &succeedingUptimeWrapper{
		uptime: uptime,
	}
}

// GetUptime implements UptimeWrapper.
func (w *succeedingUptimeWrapper) GetUptime() (uint64, error) {
	return w.uptime, nil
}

type failingUptimeWrapper struct {
	err error
}

var _ Wrapper = (*failingUptimeWrapper)(nil)

func CreateFailingUptimeWrapper(
	err error,
) Wrapper {
	return &failingUptimeWrapper{
		err: err,
	}
}

// GetUptime implements UptimeWrapper.
func (w *failingUptimeWrapper) GetUptime() (uint64, error) {
	return 0, w.err
}
