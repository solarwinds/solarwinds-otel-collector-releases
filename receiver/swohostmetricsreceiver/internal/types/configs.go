package types

/*
Expected config example

swohostmetrics:
	collection_interval: <duration>
	root_path: <string>
	scrapers:
		hostinfo:
			metrics:
				swo.hostinfo.uptime:
					enabled: true
				swo.hostinfo.whatever:
					enabled: false
		another:
			metrics:
				swo.another.whatever:
					enabled:true
		...
*/

// Scraper metrics configuration.
type MetricSettingsConfig struct {
	// flag which indicates if metric is enabled or disabled for scraping
	Enabled bool `mapstructure:"enabled"`
}

// Scraper configuration.
type ScraperConfig struct {
	// scraper provides metrics scraping.
	Metrics map[string]MetricSettingsConfig `mapstructure:"metrics"`
}

// CreateDefaultScraperConfig creates scraper config usable as general
// configuration struct for common/usual use.
func CreateDefaultScraperConfig() *ScraperConfig {
	return &ScraperConfig{
		Metrics: map[string]MetricSettingsConfig{},
	}
}
