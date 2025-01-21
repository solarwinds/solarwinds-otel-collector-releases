package scope

import "github.com/solarwinds-cloud/uams-otel-collector-plugin/test"

type EmitterMock struct {
	emitResult *Result
	initResult error
	name       string
	EmitCC     test.CallsCounter
	InitCC     test.CallsCounter
	NameCC     test.CallsCounter
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
		EmitCC:     test.CallsCounter{},
		InitCC:     test.CallsCounter{},
		NameCC:     test.CallsCounter{},
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
