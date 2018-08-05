package ns

import (
	"encoding/json"
	"os"
	"path"
	"sync"
	"time"

	etcd "github.com/coreos/etcd/client"
	"github.com/glorydeath/nux"
	log "github.com/sirupsen/logrus"
	"github.com/wusuopubupt/go-slides/lessons/lesson06/agent/metric"
	"github.com/wusuopubupt/go-slides/lessons/lesson06/agent/utils"
	"golang.org/x/net/context"
)

const (
	machinePrefix      = "/agent"
	machineStateSuffix = "state"
	GB                 = 1024 * 1024 * 1024
	InstanceCPU        = 1
	InstanceMem        = 8
	InstanceDisk       = 40
)

type Capacity struct {
	CpuCores     int     `json:"cpuCores"`
	Memory       float64 `json:"memory"`
	DiskCapacity float64 `json:"diskCapacity"`
}

// MachineState represents a point-in-time snapshot of the
// state of the local host.
type MachineState struct {
	Total          Capacity `json:"total"`
	Available      Capacity `json:"available"`
	AvailableSlot  Capacity `json:"availableSlot"`
	Hostname       string   `json:"hostname"`
	InstanceRegIP  string   `json:"instanceRegIP"`
	OS             string   `json:"os"`
	PublicIP       string   `json:"publicIP"`
	Version        string   `json:"version"`
	LastUpdateTime string   `json:"lastUpdateTime"`
}

type Machine struct {
	sync.RWMutex
	staticState  MachineState
	dynamicState *MachineState
	reg          EtcdRegistry
}

func (r *EtcdRegistry) Machines() (machines []MachineState, err error) {
	key := r.prefixed(machinePrefix)
	opts := &etcd.GetOptions{
		Sort:      true,
		Recursive: true,
	}

	resp, err := r.kAPI.Get(context.Background(), key, opts)
	if err != nil {
		if isEtcdError(err, etcd.ErrorCodeKeyNotFound) {
			err = nil
		}
		return
	}

	for _, node := range resp.Node.Nodes {
		var mach MachineState
		mach, err = readMachineState(node)
		if err != nil {
			return
		}

		if mach.Hostname != "" {
			machines = append(machines, mach)
		}
	}

	return
}

func (r *EtcdRegistry) CreateMachineState(ms MachineState, ttl time.Duration) (uint64, error) {
	ms.LastUpdateTime = utils.GetCurrentTimeStr()
	val, err := json.Marshal(ms)
	if err != nil {
		return uint64(0), err
	}

	key := r.prefixed(machinePrefix, ms.InstanceRegIP, machineStateSuffix)
	opts := &etcd.SetOptions{
		PrevExist: etcd.PrevNoExist,
		TTL:       ttl,
	}
	//log.Debug("key: ", key)
	resp, err := r.kAPI.Set(context.Background(), key, string(val), opts)
	if err != nil {
		return uint64(0), err
	}

	return resp.Node.ModifiedIndex, nil
}

func (r *EtcdRegistry) SetMachineState(ms MachineState, ttl time.Duration) (uint64, error) {
	ms.LastUpdateTime = utils.GetCurrentTimeStr()
	val, err := json.Marshal(ms)
	if err != nil {
		return uint64(0), err
	}

	key := r.prefixed(machinePrefix, ms.InstanceRegIP, machineStateSuffix)
	opts := &etcd.SetOptions{
		PrevExist: etcd.PrevExist,
		TTL:       ttl,
	}
	resp, err := r.kAPI.Set(context.Background(), key, string(val), opts)
	if err == nil {
		return resp.Node.ModifiedIndex, nil
	}

	// If state was not present, explicitly create it so the other members
	// in the cluster know this is a new member
	opts.PrevExist = etcd.PrevNoExist

	resp, err = r.kAPI.Set(context.Background(), key, string(val), opts)
	if err != nil {
		return uint64(0), err
	}

	return resp.Node.ModifiedIndex, nil
}

func (r *EtcdRegistry) MachineState(machID string) (MachineState, error) {
	key := path.Join(r.keyPrefix, machinePrefix, machID)
	opts := &etcd.GetOptions{
		Recursive: true,
		Sort:      true,
	}

	resp, err := r.kAPI.Get(context.Background(), key, opts)
	if err != nil {
		return MachineState{}, err
	}

	return readMachineState(resp.Node)
}

func (r *EtcdRegistry) RemoveMachineState(machID string) error {
	key := r.prefixed(machinePrefix, machID, machineStateSuffix)
	_, err := r.kAPI.Delete(context.Background(), key, nil)
	if isEtcdError(err, etcd.ErrorCodeKeyNotFound) {
		err = nil
	}
	return err
}

// readMachineState reads machine state from an etcd node
func readMachineState(node *etcd.Node) (mach MachineState, err error) {
	for _, obj := range node.Nodes {
		err = json.Unmarshal([]byte(obj.Value), &mach)
		if err != nil {
			log.Errorf("Parse machine state %s, error: %s", obj.Value, err)
			return
		}
	}
	return
}

// Refresh updates the current state of the Machine.
func (m *Machine) Refresh() {
	m.RLock()
	defer m.RUnlock()

	cs := m.currentState()
	if cs == nil {
		log.Fatal("Unable to refresh machine state")
	} else {
		log.Print("Machine current state: %s", cs)
		m.dynamicState = cs
	}
}

// PeriodicRefresh updates the current state of the Machine at the
// interval indicated. Operation ceases when the provided channel is closed.
func (m *Machine) PeriodicRefresh(interval time.Duration, stop <-chan struct{}) {
	// keep maintain CPU&Disk stat in background
	ticker := time.NewTicker(interval)
	for {
		select {
		case <-stop:
			log.Print("Halting Machine.PeriodicRefresh")
			ticker.Stop()
			return
		case <-ticker.C:
			metric.UpdateCpuStat()
			metric.UpdateDiskStats()
			m.Refresh()
		}
	}
}

func (m *Machine) resourceUsage() (cpu int, mem float64, disk float64) {
	return
}

func (m *Machine) GetSlotCapacity() (cap Capacity) {
	cpuUsed, memUsed, diskUsed := m.resourceUsage()
	cap = Capacity{
		CpuCores:     m.State().Total.CpuCores - cpuUsed,
		Memory:       m.State().Total.Memory - memUsed,
		DiskCapacity: m.State().Total.DiskCapacity - diskUsed,
	}
	return
}

// currentState generates a MachineState object with the values read from
// the local system
func (m *Machine) currentState() *MachineState {
	hostname, err := os.Hostname()
	if err != nil {
		log.Errorf("Error retrieving hostname: %v\n", err)
		return nil
	}

	var diskCapacityTotal float64
	var diskCapacityFree float64

	df_map := metric.GetDirDf("/home")
	if df_map == nil {
		df_map = metric.GetDirDf("/")
	}

	if df_map == nil {
		log.Error("Error retriveing Disk Cap")
		diskCapacityFree = 0
		diskCapacityTotal = 0
	} else {
		diskCapacityFree = float64((*df_map)["free"] / GB)
		diskCapacityTotal = float64((*df_map)["total"]) / GB
	}

	mem, err := nux.MemInfo()

	if err != nil {
		log.Errorf("Error retrieving MemInfo: %v\n", err)
	}

	if mem == nil {
		mem = &nux.Mem{0, 0, 0, 0, 0, 0, 0}
	}

	publicIP := ""

	return &MachineState{
		Total: Capacity{
			CpuCores:     nux.NumCpu(),
			Memory:       float64(mem.MemTotal) / GB,
			DiskCapacity: diskCapacityTotal,
		},
		Available: Capacity{
			CpuCores:     int(float64(nux.NumCpu()) * metric.CpuIdle() / 100.0),
			Memory:       float64((mem.MemFree + mem.Buffers + mem.Cached) / GB),
			DiskCapacity: diskCapacityFree,
		},
		AvailableSlot: m.GetSlotCapacity(),
		Hostname:      hostname,
		PublicIP:      publicIP,
	}
}

func NewMachine(reg EtcdRegistry, static MachineState) *Machine {
	log.Print("Created Machine with static state %s", static)
	m := &Machine{
		staticState: static,
		reg:         reg,
	}
	return m
}

func (m *Machine) String() string {
	return m.State().Hostname
}

// stackState is used to merge two MachineStates. Values configured on the top
// MachineState always take precedence over those on the bottom.
func stackState(top, bottom MachineState) MachineState {
	state := MachineState(bottom)

	if top.Hostname != "" {
		state.Hostname = top.Hostname
	}

	if top.PublicIP != "" {
		state.PublicIP = top.PublicIP
	}

	if top.InstanceRegIP != "" {
		state.InstanceRegIP = top.InstanceRegIP
	}

	if top.OS != "" {
		state.OS = top.OS
	}

	if top.Version != "" {
		state.Version = top.Version
	}

	if top.Total.CpuCores != 0 {
		state.Total.CpuCores = top.Total.CpuCores
	}

	if top.Total.Memory != 0 {
		state.Total.Memory = top.Total.Memory
	}

	if top.Total.DiskCapacity != 0 {
		state.Total.DiskCapacity = top.Total.DiskCapacity
	}

	return state
}

// State returns a MachineState object representing the Machine's
// static state overlaid on its dynamic state at the time of execution.
func (m *Machine) State() (state MachineState) {
	m.RLock()
	defer m.RUnlock()

	if m.dynamicState == nil {
		state = MachineState(m.staticState)
	} else {
		state = stackState(m.staticState, *m.dynamicState)
	}

	return
}
