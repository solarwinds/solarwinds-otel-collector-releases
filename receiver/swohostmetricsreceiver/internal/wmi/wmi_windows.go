package wmi

import (
	"github.com/yusufpapurcu/wmi"
)

type executor struct{}

func NewExecutor() Executor {
	return &executor{}
}

// CreateQuery creates string WMI query based on the input.
// Note: Implementation of WMI always requires pointer to array of src.
func (*executor) CreateQuery(src interface{}, where string) string {
	return wmi.CreateQuery(src, where)
}

// Query executes WMI query and modifies dst with its result.
// Returns modified object and error in case query fails.
// Note: Implementation of WMI always requires pointer to array of dst.
func (*executor) Query(query string, dst interface{}) (interface{}, error) {
	err := wmi.Query(query, dst)
	return dst, err
}
