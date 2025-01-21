package scope

import "github.com/solarwinds/solarwinds-otel-collector/pkg/testutil"

type EmitterMock struct {
	emitResult *Result
	initResult error
	name       string
	EmitCC     testutil.CallsCounter
	InitCC     testutil.CallsCounter
	NameCC     testutil.CallsCounter
}

var _ Emitter = (*EmitterMock)(nil)

func CreateEmitterMock(
	emitResult *Result,
	initResult error,
	name string,
) *EmitterMock {
	return &EmitterMock{
		emitResult: emitResult,
		initResult: initResult,
		name:       name,
		EmitCC:     testutil.CallsCounter{},
		InitCC:     testutil.CallsCounter{},
		NameCC:     testutil.CallsCounter{},
	}
}

// Emit implements Emitter.
func (e *EmitterMock) Emit() *Result {
	e.EmitCC.IncrementCount()
	return e.emitResult
}

// Init implements Emitter.
func (e *EmitterMock) Init() error {
	e.InitCC.IncrementCount()
	return e.initResult
}

// Name implements Emitter.
func (e *EmitterMock) Name() string {
	e.NameCC.IncrementCount()
	return e.name
}
