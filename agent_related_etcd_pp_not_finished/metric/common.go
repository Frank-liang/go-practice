package metric

import (
	"strings"
)

type MetricValue struct {
	Endpoint  string      `json:"endpoint"`
	Metric    string      `json:"metric"`
	Value     interface{} `json:"value"`
	Step      int64       `json:"step"`
	Type      string      `json:"counterType"`
	Tags      string      `json:"tags"`
	Timestamp int64       `json:"timestamp"`
}

func NewMetricValue(metric string, val interface{}, dataType string, tags ...string) *MetricValue {
	mv := MetricValue{
		Metric: metric,
		Value:  val,
		Type:   dataType,
	}

	size := len(tags)

	if size > 0 {
		mv.Tags = strings.Join(tags, ",")
	}

	return &mv
}

func GaugeValue(metric string, val interface{}, tags ...string) *MetricValue {
	return NewMetricValue(metric, val, "GAUGE", tags...)
}

func CounterValue(metric string, val interface{}, tags ...string) *MetricValue {
	return NewMetricValue(metric, val, "COUNTER", tags...)
}
