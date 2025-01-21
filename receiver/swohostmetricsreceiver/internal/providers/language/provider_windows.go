package language

import (
	"fmt"

	"github.com/solarwinds/solarwinds-otel-collector/receiver/swohostmetricsreceiver/internal/providers"
	"github.com/solarwinds/solarwinds-otel-collector/receiver/swohostmetricsreceiver/internal/registry"
	"go.uber.org/zap"
)

const (
	controlPanelKey  = "Control Panel"
	internationalKey = "International"
	localeValueName  = "LocaleName"
)

type provider struct {
	displayNamesMap   map[string]string
	getRegistryValues registry.GetKeyValuesTypeFunc
	lcidProvider      LCIDProvider
}

var _ providers.Provider[Language] = (*provider)(nil)

func CreateLanguageProvider() providers.Provider[Language] {
	return &provider{
		displayNamesMap:   getDisplayLanguages(),
		getRegistryValues: registry.GetKeyValues,
		lcidProvider:      NewWindowsLCIDProvider(),
	}
}

// Provide implements Provider.
func (p *provider) Provide() <-chan Language {
	ch := make(chan Language)
	go func() {
		defer close(ch)
		lang, err := p.getLanguageInfo()
		if err != nil {
			zap.L().Error("failed to get language info", zap.Error(err))
		}

		if err == nil {
			ch <- *lang
		}
	}()
	return ch
}

func (p *provider) getLanguageInfo() (*Language, error) {
	names := []string{localeValueName}
	values, err := p.getRegistryValues(registry.CurrentUserKey, controlPanelKey, internationalKey, names)
	if err != nil {
		return nil, fmt.Errorf("failed to get current user locale info: %w", err)
	}

	shortName := values[localeValueName]
	displayName, ok := p.displayNamesMap[shortName]
	if !ok {
		return nil, fmt.Errorf("display name not found for language '%s'", shortName)
	}

	lcid := p.lcidProvider.GetUserDefaultLCID()
	lang := &Language{
		LCID:        int(lcid),
		Name:        shortName,
		DisplayName: displayName,
	}

	return lang, nil
}
