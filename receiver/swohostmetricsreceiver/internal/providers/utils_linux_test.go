package providers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ParseKeyValues_ReturnsExpectedResult(t *testing.T) {
	expectedResult := map[string]string{
		"Time zone": "Expected value",
	}
	text := "Time zone: Expected value"
	actualResult := ParseKeyValue(text, ": ", []string{"Time zone"})

	assert.Equal(t, expectedResult, actualResult)
}

func Test_ParseKeyValues_ReturnsEmptyMapForInvalidString(t *testing.T) {
	expectedResult := map[string]string{}
	text := "Time zone: Invalid: value"
	actualResult := ParseKeyValue(text, ": ", []string{"Time zone"})

	assert.Equal(t, expectedResult, actualResult)
}

func Test_ParseKeyValues_ReturnsEmptyMapForEmptyInput(t *testing.T) {
	expectedResult := map[string]string{}
	text := ""
	actualResult := ParseKeyValue(text, ": ", []string{"Time zone"})

	assert.Equal(t, expectedResult, actualResult)
}

func Test_ParseKeyValues_ReturnsEmptyMapForInvalidSeparator(t *testing.T) {
	expectedResult := map[string]string{}
	text := "Time zone: Invalid: value"
	actualResult := ParseKeyValue(text, "invalid separator", []string{"Time zone"})

	assert.Equal(t, expectedResult, actualResult)
}

func Test_ParseKeyValues_ReturnsExpectedMapForAllAttributes(t *testing.T) {
	expectedResult := map[string]string{
		"First param":  "value",
		"Second param": "value2",
		"Third param":  "value3",
	}
	text := `First param: value
Second param: value2
Third param: value3`
	actualResult := ParseKeyValue(text, ": ", []string{"First param", "Second param", "Third param"})

	assert.Equal(t, expectedResult, actualResult)
}

func Test_ParseKeyValues_FiltersInvalidLinesAndReturnsExpectedMap(t *testing.T) {
	expectedResult := map[string]string{
		"First param":  "value",
		"Second param": "value2",
		"Third param":  "value3",
	}
	text := `First param: value

Second param: value2
Invalid param=HELLO
Third param: value3`
	actualResult := ParseKeyValue(text, ": ", []string{"First param", "Second param", "Third param"})

	assert.Equal(t, expectedResult, actualResult)
}
