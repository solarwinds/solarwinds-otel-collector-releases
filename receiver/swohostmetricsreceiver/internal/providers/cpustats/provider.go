package cpustats

const UserHz = 100

const (
	UserMode      = "user"
	NiceProc      = "nice"
	SystemMode    = "system"
	IdleState     = "idle"
	IOWait        = "io_waits"
	IRQ           = "irq"
	SoftIRQ       = "softirq"
	StealTime     = "steal"
	GuestTime     = "guest"
	GuestNiceProc = "guest_nice"
	TotalTime     = "total"

	FieldTypeCPUTime      = "cpu_time"
	FieldTypeProcesses    = "processes"
	FieldTypeCurrentProcs = "current_procs"
	FieldTypeIntr         = "intr"
	FieldTypeCtxt         = "ctxt"
	FieldTypeNumCores     = "numcores"
)

type WorkDetail struct {
	AttrName  string
	AttrValue string
	Value     float64
}

type Container struct {
	WorkDetails  map[string][]WorkDetail
	Error        error
	totalCPUTime float64
	numCPUs      float64
}
