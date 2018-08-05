package heart

import (
	"time"

	ns "github.com/wusuopubupt/go-slides/lessons/lesson06/agent/nameservice"
)

type Heart interface {
	Beat(time.Duration) (uint64, error)
	Clear() error
	Register(time.Duration) (uint64, error)
}

func New(reg ns.EtcdRegistry, mach *ns.Machine) Heart {
	return &machineHeart{reg, mach}
}

type machineHeart struct {
	reg  ns.EtcdRegistry
	mach *ns.Machine
}

func (h *machineHeart) Register(ttl time.Duration) (uint64, error) {
	return h.reg.CreateMachineState(h.mach.State(), ttl)
}

func (h *machineHeart) Beat(ttl time.Duration) (uint64, error) {
	return h.reg.SetMachineState(h.mach.State(), ttl)
}

func (h *machineHeart) Clear() error {
	return h.reg.RemoveMachineState(h.mach.State().InstanceRegIP)
}
