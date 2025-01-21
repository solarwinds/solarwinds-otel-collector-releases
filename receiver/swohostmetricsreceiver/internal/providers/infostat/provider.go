package infostat

import (
	"fmt"

	"github.com/solarwinds/solarwinds-otel-collector/receiver/swohostmetricsreceiver/internal/providers"

	"github.com/shirou/gopsutil/v3/host"
	"go.uber.org/zap"
)

type InfoStat struct {
	Hostname             string
	BootTime             uint64
	Os                   string
	Platform             string
	PlatformFamily       string
	PlatformVersion      string
	KernelVersion        string
	KernelArchitecture   string
	VirtualizationSystem string
	VirtualizationRole   string
	HostID               string
}

type provider struct {
	internalExecutor infoStatExecutor
}

var _ providers.Provider[InfoStat] = (*provider)(nil)

func CreateInfoStatProvider() providers.Provider[InfoStat] {
	return &provider{
		internalExecutor: &executor{},
	}
}

// Wrapper for host.Info() implementation.
type infoStatExecutor interface {
	Getinfo() (*host.InfoStat, error)
}
type executor struct{}

var _ infoStatExecutor = (*executor)(nil)

// Getinfo implements infoStatExecutor.
func (*executor) Getinfo() (*host.InfoStat, error) {
	return host.Info()
}

// Provide implements Provider.
func (is *provider) Provide() <-chan InfoStat {
	ch := make(chan InfoStat)
	go is.provideInternal(ch)
	return ch
}

func (is *provider) provideInternal(ch chan<- InfoStat) {
	defer close(ch)

	infoStat, err := is.internalExecutor.Getinfo()
	if err != nil {
		zap.L().Error("InfoStat command execution failed", zap.Error(err))
		return
	}

	infoStatDetails := InfoStat{
		Hostname:             infoStat.Hostname,
		BootTime:             infoStat.BootTime,
		Os:                   infoStat.OS,
		Platform:             infoStat.Platform,
		PlatformFamily:       infoStat.PlatformFamily,
		PlatformVersion:      infoStat.PlatformVersion,
		KernelVersion:        infoStat.KernelVersion,
		KernelArchitecture:   infoStat.KernelArch,
		VirtualizationSystem: infoStat.VirtualizationSystem,
		VirtualizationRole:   infoStat.VirtualizationRole,
		HostID:               infoStat.HostID,
	}

	zap.L().Debug(fmt.Sprintf("InfoStat provided %+v", infoStatDetails))
	ch <- infoStatDetails
}
