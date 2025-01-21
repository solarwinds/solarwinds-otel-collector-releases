package installedupdates

import (
	"fmt"
	"strings"

	"github.com/solarwinds/solarwinds-otel-collector/receiver/swohostmetricsreceiver/internal/wmi"
	"go.uber.org/zap"
)

type windowsProvider struct {
	wmi wmi.Executor
}

var _ (Provider) = (*windowsProvider)(nil)

func NewProvider() Provider {
	return createWindowsProvider(
		wmi.NewExecutor(),
	)
}

func createWindowsProvider(
	wmi wmi.Executor,
) Provider {
	return &windowsProvider{
		wmi: wmi,
	}
}

// Win32_QuickFixEngineering represents actual Hot Fix WMI Object
// with subset of fields required for scraping.
type Win32_QuickFixEngineering struct {
	Caption     string
	Description string
	HotFixID    string
	InstalledBy string
	InstalledOn string
}

func (provider *windowsProvider) GetUpdates() ([]InstalledUpdate, error) {
	result, err := wmi.QueryResult[[]Win32_QuickFixEngineering](provider.wmi)
	if err != nil {
		message := "Invalid installed updates output."
		zap.L().Error(message, zap.Error(err))

		return []InstalledUpdate{}, fmt.Errorf("%s %w", message, err)
	}

	var updates []InstalledUpdate
	for _, update := range result {
		updates = append(updates, InstalledUpdate{
			Caption:     update.Caption,
			HotFixID:    update.HotFixID,
			InstalledOn: update.InstalledOn,
			InstalledBy: update.InstalledBy,
			Description: update.Description,
		})
	}

	return formatDates(updates), nil
}

// Format date from mm/dd/yyyy to yyyy-mm-dd for all updates.
func formatDates(updates []InstalledUpdate) []InstalledUpdate {
	formattedUpdates := []InstalledUpdate{}
	for _, update := range updates {
		formattedUpdate := InstalledUpdate{
			Caption:     update.Caption,
			Description: update.Description,
			HotFixID:    update.HotFixID,
			InstalledOn: formatDate(update.InstalledOn),
			InstalledBy: update.InstalledBy,
		}
		formattedUpdates = append(formattedUpdates, formattedUpdate)
	}

	return formattedUpdates
}

// format date from mm/dd/yyyy to yyyy-mm-dd.
func formatDate(inputDate string) string {
	outputDate := ""
	dateParts := strings.Split(inputDate, "/")
	if len(dateParts) == 3 {
		outputDate = fmt.Sprintf("%s-%s-%s",
			dateParts[2],
			dateParts[0],
			dateParts[1])
	}

	return outputDate
}
