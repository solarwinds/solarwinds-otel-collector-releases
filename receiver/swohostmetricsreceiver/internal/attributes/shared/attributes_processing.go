package shared

import "sync"

// Process incoming attributes from a particular channel.
// In case channel is close the channel pointer is set to nil
// and waitGroup gets the signal that this goroutine is done.
func ProcessReceivedAttributes(
	generatedAttrs Attributes,
	resultingAttrs Attributes,
	opened bool,
	ch *AttributesChannel,
	wg *sync.WaitGroup,
) {
	if !opened {
		*ch = nil
		wg.Done()
	} else {
		mergeAttributes(generatedAttrs, resultingAttrs)
	}
}

// Takes already existing attributes and new ones. It merges them.
func mergeAttributes(
	increment Attributes,
	result Attributes,
) {
	for k, v := range increment {
		result[k] = v
	}
}
