package main

import (
	"time"

	"github.com/51reboot/golang-01-homework/lesson12/binggan/monitor/common"
)

type MetricFunc func() []*common.Metric

type Sched struct {
	ch chan *common.Metric
}

func NewSched(ch chan *common.Metric) *Sched {
	return &Sched{
		ch: ch,
	}
}

func (s *Sched) AddMetric(collecter MetricFunc, step time.Duration) {
	go func() {
		ticker := time.NewTicker(step)
		for range ticker.C {
			metrics := collecter()
			for _, metric := range metrics {
				s.ch <- metric
			}
		}
	}()
}
