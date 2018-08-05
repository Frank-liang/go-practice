package healthCheck

import (
	"log"
	"time"
)

type HealthCheckManager struct {
	reg  ns.EtcdRegistry
	mach *ns.Machine
}

func NewHealthCheckManager(reg ns.EtcdRegistry, mach *ns.Machine) *HealthCheckManager {
	return &HealthCheckManager{
		reg:  reg,
		mach: mach,
	}
}

func (hcm *HealthCheckManager) PeriodicHealthCheck(interval time.Duration, stop <-chan struct{}) {
	ticker := time.NewTicker(interval)
	for {
		select {
		case <-stop:
			log.Print("Stop!")
			ticker.Stop()
			return
		case <-ticker.C:
			hcm.healthCheck()
		}
	}
}

func (hcm *HealthCheckManager) healthCheck() (err error) {
	hcm.instancesHealthCheck()
	return
}

func (hcm *HealthCheckManager) instancesHealthCheck() (err error) {
	return
}
