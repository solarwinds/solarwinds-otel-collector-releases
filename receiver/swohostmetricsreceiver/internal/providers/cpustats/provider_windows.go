package cpustats

import (
	"context"
	"fmt"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/solarwinds/solarwinds-otel-collector/receiver/swohostmetricsreceiver/internal/providers"
)

const (
	cpuTotal   = "cpu-total"
	factorToUs = 1000000.0
)

type cpuStatsProvider interface {
	TimesWithContext(ctx context.Context, percpu bool) ([]cpu.TimesStat, error)
}

type gopsutilProvider struct{}

func (g *gopsutilProvider) TimesWithContext(ctx context.Context, percpu bool) ([]cpu.TimesStat, error) {
	return cpu.TimesWithContext(ctx, percpu)
}

type provider struct {
	cpuStatsProvider cpuStatsProvider
}

func CreateProvider() providers.Provider[Container] {
	return &provider{
		cpuStatsProvider: &gopsutilProvider{},
	}
}

// Provide implements providers.Provider.
func (p *provider) Provide() <-chan Container {
	ch := make(chan Container)

	go func() {
		defer close(ch)
		cpuStats := Container{
			WorkDetails: make(map[string][]WorkDetail),
		}

		timesStats, err := p.cpuStatsProvider.TimesWithContext(context.Background(), false)
		if err != nil {
			ch <- Container{
				Error: fmt.Errorf("get cpu times: %w", err),
			}
		}

		for _, timesStat := range timesStats {
			if timesStat.CPU != cpuTotal {
				continue
			}

			totalTime := timesStat.User + timesStat.System

			cpuStats.WorkDetails[FieldTypeCPUTime] = append(
				cpuStats.WorkDetails[FieldTypeCPUTime],
				WorkDetail{
					AttrName:  "mode",
					AttrValue: UserMode,
					Value:     timesStat.User * factorToUs,
				},
			)
			cpuStats.WorkDetails[FieldTypeCPUTime] = append(
				cpuStats.WorkDetails[FieldTypeCPUTime],
				WorkDetail{
					AttrName:  "mode",
					AttrValue: IdleState,
					Value:     timesStat.Idle * factorToUs,
				},
			)
			cpuStats.WorkDetails[FieldTypeCPUTime] = append(
				cpuStats.WorkDetails[FieldTypeCPUTime],
				WorkDetail{
					AttrName:  "mode",
					AttrValue: SystemMode,
					Value:     timesStat.System * factorToUs,
				},
			)
			cpuStats.WorkDetails[FieldTypeCPUTime] = append(
				cpuStats.WorkDetails[FieldTypeCPUTime],
				WorkDetail{
					AttrName:  "mode",
					AttrValue: TotalTime,
					Value:     totalTime * factorToUs,
				},
			)
		}

		ch <- cpuStats
	}()

	return ch
}
