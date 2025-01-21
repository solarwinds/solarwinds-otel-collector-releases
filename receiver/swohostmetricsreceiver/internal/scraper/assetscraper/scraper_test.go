package assetscraper

import (
	"context"
	"fmt"
	"testing"

	"github.com/solarwinds/solarwinds-otel-collector/receiver/swohostmetricsreceiver/internal/scraper/assetscraper/metrics/installedsoftware"
	"github.com/solarwinds/solarwinds-otel-collector/receiver/swohostmetricsreceiver/internal/scraper/assetscraper/metrics/installedupdates"
	"github.com/solarwinds/solarwinds-otel-collector/receiver/swohostmetricsreceiver/internal/types"
	"github.com/stretchr/testify/assert"
)

func Test_Functional(t *testing.T) {
	t.Skip("This test must to be run manually only")

	sc := Config{
		ScraperConfig: types.ScraperConfig{
			Metrics: map[string]types.MetricSettingsConfig{
				installedsoftware.Name: {Enabled: true},
				installedupdates.Name:  {Enabled: true},
			},
		},
	}

	s, err := NewAssetScraper(&sc)
	assert.NoError(t, err, "scraper creation must not fail")
	err = s.Start(context.TODO(), nil)
	assert.NoError(t, err)
	m, err := s.Scrape(context.TODO())
	assert.NoError(t, err)
	fmt.Printf("Result: %+v\n", m)
}
