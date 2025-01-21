package wmi

import (
	"reflect"
)

type ExecutorMock struct {
	errors  map[string]error
	results map[string]interface{}
}

func CreateWmiExecutorMock(results []interface{}, errors map[interface{}]error) Executor {
	mock := &ExecutorMock{
		errors:  make(map[string]error),
		results: make(map[string]interface{}),
	}

	for _, result := range results {
		mock.SetupResult(result)
	}

	for i, err := range errors {
		mock.SetupError(i, err)
	}

	return mock
}

func (m *ExecutorMock) SetupResult(i interface{}) {
	m.results[getImplementationName(i)] = i
}

func (m *ExecutorMock) SetupError(i interface{}, error error) {
	m.errors[getImplementationName(i)] = error
}

func (*ExecutorMock) CreateQuery(_ interface{}, _ string) string { return "not relevant" }

func (m *ExecutorMock) Query(_ string, dst interface{}) (interface{}, error) {
	err := m.errors[getImplementationName(dst)]
	if err != nil {
		return nil, err
	}

	result := m.results[getImplementationName(dst)]
	return result, nil
}

func getImplementationName(i interface{}) string {
	return reflect.TypeOf(i).String()
}
