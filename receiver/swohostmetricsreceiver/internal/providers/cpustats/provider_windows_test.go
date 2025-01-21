package cpustats

import (
	"context"
	"testing"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/stretchr/testify/assert"
)

type mockCpuStatsProvider struct{}

func (m *mockCpuStatsProvider) TimesWithContext(ctx context.Context, percpu bool) ([]cpu.TimesStat, error) {
	return []cpu.TimesStat{
		{
			CPU:    "cpu-total",
			System: 3.0,
			Idle:   0.2,
			User:   4.0,
		},
	}, nil
}

func Test_provide_Success(t *testing.T) {
	p := &provider{cpuStatsProvider: &mockCpuStatsProvider{}}
	container := <-p.Provide()

	assert.EqualValues(
		t,
		[]WorkDetail{
			{AttrName: "mode", AttrValue: "user", Value: 4000000},
			{AttrName: "mode", AttrValue: "system", Value: 3000000},
			{AttrName: "mode", AttrValue: "idle", Value: 200000},
			{AttrName: "mode", AttrValue: "total", Value: 7000000},
		},
		container.WorkDetails["cpu_time"],
	)
}
