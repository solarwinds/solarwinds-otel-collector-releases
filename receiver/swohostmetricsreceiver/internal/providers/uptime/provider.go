package uptime

type Uptime struct {
	Uptime uint64
	Error  error
}

// Provider for host uptime.
type Provider interface {
	// Provides channel capable of delivering uptime.
	GetUptime() <-chan Uptime
}

type hostProvider struct {
	Wrapper
}

var _ Provider = (*hostProvider)(nil)

func CreateUptimeProvider(
	w Wrapper,
) Provider {
	return &hostProvider{
		Wrapper: w,
	}
}

// GetUptime implements Provider.
func (p *hostProvider) GetUptime() <-chan Uptime {
	ch := make(chan Uptime)
	go p.getUptimeInternal(ch)
	return ch
}

func (p *hostProvider) getUptimeInternal(ch chan Uptime) {
	defer close(ch)

	uptime, err := p.Wrapper.GetUptime()
	d := Uptime{
		Uptime: uptime,
		Error:  err,
	}

	ch <- d
}
