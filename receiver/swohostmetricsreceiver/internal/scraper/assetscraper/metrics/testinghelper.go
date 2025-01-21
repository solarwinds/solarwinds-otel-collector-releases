package metrics

import (
	"testing"

	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/pdata/pcommon"
)

func RequireAttribute(t *testing.T, attributes pcommon.Map, attributeKey string, expectedValue string) {
	value, exists := attributes.Get(attributeKey)
	require.Truef(t, exists, "The %s attribute should exists", attributeKey)
	require.Equalf(t, expectedValue, value.AsString(), "The %s attribute should be %s", attributeKey, expectedValue)
}
