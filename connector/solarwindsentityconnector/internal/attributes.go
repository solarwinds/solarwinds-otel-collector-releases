package internal

import (
	"github.com/solarwinds/solarwinds-otel-collector-releases/connector/solarwindsentityconnector/config"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"strings"
)

type Attributes struct {
	data map[string][]Attribute
}

type Attribute struct {
	Value     pcommon.Value
	Direction config.Direction
}

func NewAttributes(resourceAttrs pcommon.Map, sourcePrefix, destPrefix string) Attributes {
	result := make(map[string][]Attribute, resourceAttrs.Len())

	appendAttribute := func(key string, value pcommon.Value, direction config.Direction) {
		result[key] = append(result[key], Attribute{
			Value:     value,
			Direction: direction,
		})
	}

	resourceAttrs.Range(func(key string, value pcommon.Value) bool {
		if strings.HasPrefix(key, sourcePrefix) {
			trimmedKey := strings.TrimPrefix(key, sourcePrefix)
			appendAttribute(trimmedKey, value, config.Source)
			return true
		}

		if strings.HasPrefix(key, destPrefix) {
			trimmedKey := strings.TrimPrefix(key, destPrefix)
			appendAttribute(trimmedKey, value, config.Destination)
			return true
		}

		appendAttribute(key, value, config.None)
		return true // Continue iteration
	})

	return Attributes{
		data: result,
	}
}

func (a *Attributes) findAttributeValue(key string, dir config.Direction) (pcommon.Value, bool) {
	attrVal, exists := a.data[key]

	if !exists {
		return pcommon.NewValueStr(""), false
	}

	for _, v := range attrVal {
		if v.Direction == dir {
			return v.Value, true
		}
	}

	return pcommon.NewValueStr(""), false
}
