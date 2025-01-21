package installedsoftware

import (
	"fmt"
	"time"

	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.uber.org/zap"

	"github.com/solarwinds/solarwinds-otel-collector/receiver/swohostmetricsreceiver/internal/providers/installedsoftware"
	metricshelper "github.com/solarwinds/solarwinds-otel-collector/receiver/swohostmetricsreceiver/internal/scraper/assetscraper/metrics"
	"github.com/solarwinds/solarwinds-otel-collector/receiver/swohostmetricsreceiver/internal/scraper/framework/metric"
)

const (
	Name        = "swo.asset.installedsoftware"
	description = "carries attributes describing installed software"
	unit        = ""
)

type Emitter struct {
	provider      installedsoftware.Provider
	startTimeNano time.Time
}

var _ metric.Emitter = (*Emitter)(nil)

func NewEmitter() metric.Emitter {
	return createInstalledSoftwareEmitter(
		installedsoftware.NewInstalledSoftwareProvider(),
	)
}

func createInstalledSoftwareEmitter(
	provider installedsoftware.Provider,
) metric.Emitter {
	return &Emitter{
		provider: provider,
	}
}

// Emit implements metric.Emitter.
func (emitter *Emitter) Emit() *metric.Result {
	ms, err := emitter.populateMetric()
	if err != nil {
		message := fmt.Sprintf("metric %s population failed", Name)
		zap.L().Error(message, zap.Error(err))

		return &metric.Result{
			Data:  pmetric.NewMetricSlice(),
			Error: fmt.Errorf(message+": %w", err),
		}
	}

	return &metric.Result{
		Data:  ms,
		Error: nil,
	}
}

// Init implements metric.Emitter.
func (emitter *Emitter) Init() error {
	emitter.startTimeNano = time.Now()

	return nil
}

// Name implements metric.Emitter.
func (emitter *Emitter) Name() string {
	return Name
}

func (emitter *Emitter) populateMetric() (pmetric.MetricSlice, error) {
	isCollection, err := emitter.provider.GetSoftware()
	if err != nil {
		message := "obtaining installed software failed"
		zap.L().Error(message, zap.Error(err))
		return pmetric.NewMetricSlice(), fmt.Errorf(message+"%w", err)
	}

	// Nothing was obtained, so nothing can be sent outside. No error behavior.
	if len(isCollection) == 0 {
		zap.L().Debug("no software was obtained from installed sofftware metric emitter")
		return pmetric.NewMetricSlice(), nil
	}

	metadata := metricshelper.MetricMetadata{
		Name:        Name,
		Description: description,
		Unit:        unit,
	}
	metricSlice := metricshelper.ConstructMetricBase(metadata)
	metric := metricSlice.At(0)

	dataPoints := metricshelper.PrepareEmptySum(metric, len(isCollection))
	for _, is := range isCollection {
		dataPoint := metricshelper.AppendNumberDataPoint(dataPoints, emitter.startTimeNano)

		rawAttributes := getAttributes(is)
		_ = dataPoint.Attributes().FromRaw(rawAttributes)
	}

	return metricSlice, nil
}
