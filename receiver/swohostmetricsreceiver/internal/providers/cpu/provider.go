package cpu

type Processor struct {
	Name         string
	Manufacturer string
	Speed        float64
	Cores        uint32
	Threads      uint32
	Model        string
	Stepping     string
}

type Container struct {
	Processors []Processor
	Error      error
}
