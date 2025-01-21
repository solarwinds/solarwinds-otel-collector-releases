package example

import (
	"context"
	"testing"

	"github.com/solarwinds/solarwinds-otel-collector/receiver/swohostmetricsreceiver/internal/types"
	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/collector/receiver"
)

func Test_CreateScraperExplicitly_ScraperIsProvided(t *testing.T) {
	config := &ScraperConfig{
		DelayedProcessingConfig: types.DelayedProcessingConfig{},
		ScraperConfig: types.ScraperConfig{
			Metrics: map[string]types.MetricSettingsConfig{
				scope1metric1: {Enabled: true},
			},
		},
	}

	sut, err := CreateScraperExplicitly(
		context.TODO(),
		receiver.Settings{},
		config,
	)

	assert.NoError(t, err, "creation must not fail")
	assert.NotNil(t, sut, "scraper must be provided")
}

func Test_CreateScraperExplicitly_FailsOnEmptyConfig(t *testing.T) {
	config := &ScraperConfig{
		ScraperConfig: types.ScraperConfig{},
	}

	sut, err := CreateScraperExplicitly(
		context.TODO(),
		receiver.Settings{},
		config,
	)

	assert.Error(t, err, "creation must fail")
	assert.Nil(t, sut, "scraper must not be provided")
}

func Test_CreateScraperImplicitly_ScraperIsProvided(t *testing.T) {
	config := &ScraperConfig{
		DelayedProcessingConfig: types.DelayedProcessingConfig{},
		ScraperConfig: types.ScraperConfig{
			Metrics: map[string]types.MetricSettingsConfig{
				scope1metric1: {Enabled: true},
			},
		},
	}

	sut, err := CreateScraperImplicitly(
		context.TODO(),
		receiver.Settings{},
		config,
	)

	assert.NoError(t, err, "creation must not fail")
	assert.NotNil(t, sut, "scraper must be provided")
}

func Test_CreateScraperImplicitly_FailsOnEmptyConfig(t *testing.T) {
	config := &ScraperConfig{
		ScraperConfig: types.ScraperConfig{},
	}

	sut, err := CreateScraperImplicitly(
		context.TODO(),
		receiver.Settings{},
		config,
	)

	assert.Error(t, err, "creation must fail")
	assert.Nil(t, sut, "scraper must not be provided")
}
