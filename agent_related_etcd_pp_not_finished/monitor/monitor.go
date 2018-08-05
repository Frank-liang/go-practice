package server

import (
	"time"

	heart "github.com/Frank-liang/go/agent_related_etcd_pp_not_finished/heartbeat"

	"log"
)

func NewMonitor(ttl time.Duration, interval time.Duration) *Monitor {
	return &Monitor{ttl, interval}
}

type Monitor struct {
	TTL      time.Duration
	interval time.Duration
}

// AgentBeat periodically checks the given Heart to make sure it
// beats successfully. If the heartbeat check fails for any
// reason, an error is returned. If the supplied channel is
// closed, Monitor returns true and a nil error.
func (m *Monitor) AgentBeat(hrt heart.Heart, sdc <-chan struct{}) (bool, error) {
	next := time.After(0)
	for {
		select {
		case <-sdc:
			return true, nil
		case <-next:
			_, err := hrt.Beat(m.TTL)
			if err != nil {
				log.Print("Monitor heartbeat function returned err, retrying in %v: %v", m.interval, err)
			}
			next = time.After(m.interval)
		}
	}
}
