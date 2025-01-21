package registry

import (
	"fmt"
)

func NewReader(_ RootKey, _ string) (Reader, error) {
	return nil, fmt.Errorf("windows registry are not supported on Darwin")
}
