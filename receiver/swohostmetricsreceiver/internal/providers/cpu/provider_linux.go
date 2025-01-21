package cpu

import (
	"strconv"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/solarwinds/solarwinds-otel-collector/receiver/swohostmetricsreceiver/internal/providers"
)

type cpuExecutor interface {
	Info() ([]cpu.InfoStat, error)
}

type executor struct{}

func (e *executor) Info() ([]cpu.InfoStat, error) {
	return cpu.Info()
}

var _ cpuExecutor = (*executor)(nil)

type provider struct {
	executor cpuExecutor
}

func CreateProvider() providers.Provider[Container] {
	return &provider{
		executor: &executor{},
	}
}

// Provide implements providers.Provider.
func (p *provider) Provide() <-chan Container {
	ch := make(chan Container)
	go func() {
		defer close(ch)

		infoStats, err := p.executor.Info()
		if err != nil {
			ch <- Container{Error: err}
			return
		}

		processors := make(map[string]Processor, len(infoStats))
		cores := make(map[CoreKey]bool, len(infoStats))
		for _, infoStat := range infoStats {
			processor, exists := processors[infoStat.PhysicalID]
			if !exists {
				processor = Processor{
					Name:         infoStat.ModelName,
					Manufacturer: infoStat.VendorID,
					Speed:        infoStat.Mhz,
					Cores:        0,
					Threads:      0,
					Model:        infoStat.Model,
					Stepping:     strconv.FormatInt(int64(infoStat.Stepping), 10),
				}
			}

			// InfoStat for Linux is per virtual CPU (thread), so cores and threads have to be counted manually
			processor.Threads++
			coreKey := CoreKey{infoStat.PhysicalID, infoStat.CoreID}
			if _, exists := cores[coreKey]; !exists {
				cores[coreKey] = true
				processor.Cores++
			}

			processors[infoStat.PhysicalID] = processor
		}

		ch <- Container{Processors: toSlice(processors)}
	}()
	return ch
}

func toSlice[T any](m map[string]T) []T {
	list := make([]T, 0, len(m))
	for _, v := range m {
		list = append(list, v)
	}
	return list
}

type CoreKey struct {
	ProcessorID string
	CoreID      string
}
