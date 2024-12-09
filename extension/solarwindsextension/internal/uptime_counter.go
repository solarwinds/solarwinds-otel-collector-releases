package internal

import "time"

type uptimeCounter struct {
	startTimeUnixNano int64
}

func newUptimeCounter() *uptimeCounter {
	return &uptimeCounter{startTimeUnixNano: time.Now().UnixNano()}
}

func (u *uptimeCounter) Get() float64 {
	// Borrowed from Collector's processor uptime metric
	return float64(time.Now().UnixNano()-u.startTimeUnixNano) / 1e9
}
