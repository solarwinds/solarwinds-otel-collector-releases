package metrics

import (
	"time"

	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
)

type MetricMetadata struct {
	Name        string
	Description string
	Unit        string
}

// Create new metric slice with one metric. The metric slice is returned.
func ConstructMetricBase(metadata MetricMetadata) pmetric.MetricSlice {
	slice := pmetric.NewMetricSlice()
	slice.EnsureCapacity(1)

	metric := slice.AppendEmpty()
	metric.SetDescription(metadata.Description)
	metric.SetName(metadata.Name)
	metric.SetUnit(metadata.Unit)

	return slice
}

/*
Set empty, monotonic sum to the metric and allocate capacity for data points.
The allocated data points slice is returned.
*/
func PrepareEmptySum(metric pmetric.Metric, capacity int) pmetric.NumberDataPointSlice {
	sum := metric.SetEmptySum()
	sum.SetIsMonotonic(true)
	sum.SetAggregationTemporality(pmetric.AggregationTemporalityUnspecified)

	dataPoints := sum.DataPoints()
	dataPoints.EnsureCapacity(capacity)

	return dataPoints
}

/*
Append new number data point to the number data points slice.
The new number data point with 0 value is returned.

The `startTimestampNano` cannot be negative.
*/
func AppendNumberDataPoint(dataPoints pmetric.NumberDataPointSlice, startTimestampNano time.Time) pmetric.NumberDataPoint {
	dataPoint := dataPoints.AppendEmpty()

	dataPoint.SetIntValue(0)
	dataPoint.SetFlags(pmetric.DefaultDataPointFlags)
	dataPoint.SetStartTimestamp(pcommon.NewTimestampFromTime(startTimestampNano))
	now := time.Now()
	dataPoint.SetTimestamp(pcommon.NewTimestampFromTime(now))

	return dataPoint
}
