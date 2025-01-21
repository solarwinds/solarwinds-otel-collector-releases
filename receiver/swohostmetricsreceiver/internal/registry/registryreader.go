package registry

// TODO: Create proper wrapper around registry access https://swicloud.atlassian.net/browse/NH-67748

type Reader interface {
	GetSubKeys() ([]string, error)
	GetKeyValues(keyName string, names []string) (map[string]string, error)
	GetKeyUIntValue(keyName string, name string) (uint64, error)
}

type RootKey int32

const (
	LocalMachineKey RootKey = 0
	CurrentUserKey  RootKey = 1
)
