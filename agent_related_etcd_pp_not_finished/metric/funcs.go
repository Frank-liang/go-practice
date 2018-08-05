package metric

import (
	"os"
	"time"
)

type FuncsAndInterval struct {
	Fs       []func() []*MetricValue
	Interval int
}

var Mappers []FuncsAndInterval

func BuildMappers() {
	// todo interval config
	interval := 30
	Mappers = []FuncsAndInterval{
		FuncsAndInterval{
			Fs: []func() []*MetricValue{
				CpuMetrics,
				KernelMetrics,
				LoadAvgMetrics,
				MemMetrics,
				DiskIOMetrics,
				IOStatsMetrics,
				NetstatMetrics,
				UdpMetrics,
			},
			Interval: interval,
		},
		FuncsAndInterval{
			Fs: []func() []*MetricValue{
				DeviceMetrics,
			},
			Interval: interval,
		},
		FuncsAndInterval{
			Fs: []func() []*MetricValue{
				DuMetrics,
			},
			Interval: interval,
		},
	}
}

func RefreshDeltaData() {
	for {
		UpdateCpuStat()
		UpdateDiskStats()
		time.Sleep(time.Second)
	}
}

func collect(sec int64, fns []func() []*MetricValue) (mvs []*MetricValue) {

	hostname, err := os.Hostname()
	if err != nil {
		return
	}

	mvs = []*MetricValue{}

	for _, fn := range fns {
		items := fn()
		if items == nil {
			continue
		}

		if len(items) == 0 {
			continue
		}

		for _, mv := range items {
			mvs = append(mvs, mv)
		}
	}

	now := time.Now().Unix()
	for j := 0; j < len(mvs); j++ {
		mvs[j].Step = sec
		mvs[j].Endpoint = hostname
		mvs[j].Timestamp = now
	}
	return
}
