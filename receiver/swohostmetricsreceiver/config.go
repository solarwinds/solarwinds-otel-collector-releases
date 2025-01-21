package swohostmetricsreceiver

import (
	"fmt"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/confmap"
	"go.opentelemetry.io/collector/receiver/scraperhelper"
	"go.uber.org/zap"
)

/*
Expected config example

swohostmetrics:
	collection_interval: <duration>
	scrapers:
		hostinfo:
			...
		another:
			...
*/

// ReceiverConfig defines SWO host metrics configuration.
type ReceiverConfig struct {
	// common receiver settings.
	scraperhelper.ControllerConfig `mapstructure:",squash"`

	// available scrapers for receiver.
	Scrapers map[string]component.Config `mapstructure:"-"`
}

var (
	_ component.Config    = (*ReceiverConfig)(nil) // Type check against Config
	_ confmap.Unmarshaler = (*ReceiverConfig)(nil) // Type check against Unmarshaller
)

// Unmarshal implements confmap.Unmarshaler.
func (receiverConfig *ReceiverConfig) Unmarshal(rawConfig *confmap.Conf) error {
	if receiverConfig == nil {
		return fmt.Errorf("receiverConfig function receiver is nil")
	}

	if rawConfig == nil {
		return fmt.Errorf("raw configuration object is nil")
	}

	const logErrorInclude = ": %w"
	// try to unmarshall raw config into receiver config
	err := rawConfig.Unmarshal(receiverConfig, confmap.WithIgnoreUnused())
	if err != nil {
		message := "Config unmarshalling failed"
		zap.L().Error(message, zap.Error(err))
		return fmt.Errorf(message+logErrorInclude, err)
	}

	// loading scrapers config section
	scrapersSectionConfigMap, err := rawConfig.Sub("scrapers")
	if err != nil {
		message := "Failed to fetch scrapers section from config"
		zap.L().Error(message, zap.Error(err))
		return fmt.Errorf(message+logErrorInclude, err)
	}

	// processing scrapers
	scraperMap := scrapersSectionConfigMap.ToStringMap()
	receiverConfig.Scrapers = make(map[string]component.Config, len(scraperMap))
	for scraperName := range scraperMap {
		scraperFactory, err := GetScraperFactory(scraperName)
		if err != nil {
			message := fmt.Sprintf("Scraper factory for scraper %s was not found", scraperName)
			zap.L().Error(message, zap.Error(err))
			return fmt.Errorf(message+logErrorInclude, err)
		}

		// loads scraper config with default values
		scraperConfig := scraperFactory.CreateDefaultConfig()
		// extracting scraper config from configuration map
		scraperSectionConfigMap, err := scrapersSectionConfigMap.Sub(scraperName)
		if err != nil {
			message := fmt.Sprintf("Scraper configuration for scraper %s can not be fetched", scraperName)
			zap.L().Error(message, zap.Error(err))
			return fmt.Errorf(message+logErrorInclude, err)
		}

		// unmarshal it into scraper configuration struct
		err = scraperSectionConfigMap.Unmarshal(scraperConfig, confmap.WithIgnoreUnused())
		if err != nil {
			message := fmt.Sprintf("Umarshalling config for scraper %s failed", scraperName)
			zap.L().Error(message, zap.Error(err))
			return fmt.Errorf(message+logErrorInclude, err)
		}

		// set up unmarshalled config for given scraper
		receiverConfig.Scrapers[scraperName] = scraperConfig
	}

	return nil
}
