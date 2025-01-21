package uptime

import "github.com/shirou/gopsutil/v3/host"

type Wrapper interface {
	GetUptime() (uint64, error)
}

type wrapper struct{}

func CreateUptimeWrapper() Wrapper {
	return &wrapper{}
}

// GetUptime implements UptimeWrapper.
func (*wrapper) GetUptime() (uint64, error) {
	return host.Uptime()
}

var _ Wrapper = (*wrapper)(nil)
