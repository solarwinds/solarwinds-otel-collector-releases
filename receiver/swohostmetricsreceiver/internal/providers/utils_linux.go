package providers

import "strings"

// ParseKeyValue takes text, splits it to separate lines, then
// is trying to find keys in lines split by the passed separator.
func ParseKeyValue(text, separator string, keys []string) map[string]string {
	attributesMap := make(map[string]string)
	for _, line := range strings.Split(text, "\n") {
		splitAttribute := strings.Split(strings.TrimSpace(line), separator)
		// correct result should consist of key and value property
		if len(splitAttribute) != 2 {
			continue
		}

		// Add wanted keys only to the map
		for _, key := range keys {
			if key == splitAttribute[0] {
				attributesMap[key] = splitAttribute[1]
				break
			}
		}
	}
	return attributesMap
}
