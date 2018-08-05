package metric

import (
	"log"

	"github.com/glorydeath/nux"
)

func LoadAvgMetrics() []*MetricValue {
	load, err := nux.LoadAvg()
	if err != nil {
		log.Println(err)
		return nil
	}

	return []*MetricValue{
		GaugeValue("load.1min", load.Avg1min),
		GaugeValue("load.5min", load.Avg5min),
		GaugeValue("load.15min", load.Avg15min),
	}

}
