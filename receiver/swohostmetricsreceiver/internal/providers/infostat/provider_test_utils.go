package infostat

import "github.com/shirou/gopsutil/v3/host"

type infoExecutorMock struct {
	infoStat *host.InfoStat
	err      error
}

// Getinfo implements InfoStatExecutor.
func (iem *infoExecutorMock) Getinfo() (*host.InfoStat, error) {
	return iem.infoStat, iem.err
}

var _ infoStatExecutor = (*infoExecutorMock)(nil)

func CreateInfoStatProviderMock(infoStat *host.InfoStat, err error) infoStatExecutor {
	return &infoExecutorMock{
		infoStat: infoStat,
		err:      err,
	}
}
