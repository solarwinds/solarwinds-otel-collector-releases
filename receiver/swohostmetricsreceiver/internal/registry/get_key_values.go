package registry

// TODO: Create proper wrapper around registry access https://swicloud.atlassian.net/browse/NH-67748

type GetKeyValuesTypeFunc func(rootKey RootKey, parent, keyName string, names []string) (map[string]string, error)

func GetKeyValues(rootKey RootKey, parent, keyName string, names []string) (map[string]string, error) {
	regReader, err := NewReader(rootKey, parent)
	if err != nil {
		return nil, err
	}

	return regReader.GetKeyValues(keyName, names)
}

type GetKeyUIntValueTypeFunc func(rootKey RootKey, parent, keyName string, name string) (uint64, error)

func GetKeyUIntValue(rootKey RootKey, parent, keyName string, name string) (uint64, error) {
	regReader, err := NewReader(rootKey, parent)
	if err != nil {
		return 0, err
	}

	return regReader.GetKeyUIntValue(keyName, name)
}
