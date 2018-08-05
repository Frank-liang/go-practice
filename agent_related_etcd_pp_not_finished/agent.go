package agent

import (
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"

	etcd "github.com/coreos/etcd/client"
	log "github.com/sirupsen/logrus"
	"github.com/wusuopubupt/go-slides/lessons/lesson06/agent/config"
	"github.com/wusuopubupt/go-slides/lessons/lesson06/agent/metric"
	monitor "github.com/wusuopubupt/go-slides/lessons/lesson06/agent/monitor"
)

const (
	machineStateRefreshInterval   = 5 * time.Second
	taskManagerGetInterval        = 3 * time.Second
	healthcheckManagerGetInterval = 5 * time.Second
)

type Agent struct {
	mach   *ns.Machine
	hrt    heart.Heart
	mon    *monitor.Monitor
	tm     *taskManager.TaskManager
	hcm    *healthCheck.HealthCheckManager
	reg    ns.EtcdRegistry
	killc  chan struct{}  // used to signal monitor to shutdown Agent
	stopc  chan struct{}  // used to terminate deploy and undeploy
	mstopc chan struct{}  // used to terminate health check and machine
	wg     sync.WaitGroup // used to co-ordinate shutdown
}

func NewMachineFromConfig(reg ns.EtcdRegistry, cfg config.Config) (*ns.Machine, error) {
	state := ns.MachineState{
		PublicIP:      cfg.AgentConfig.PublicIP,
		InstanceRegIP: cfg.InstanceRegIP,
		Total: ns.Capacity{
			CpuCores:     cfg.AgentConfig.CpuCores,
			Memory:       cfg.AgentConfig.Memory,
			DiskCapacity: cfg.AgentConfig.DiskCapacity,
		},
		Version: config.VersionStr,
	}

	mach := ns.NewMachine(reg, state)
	mach.Refresh()
	return mach, nil
}

func New(cfg config.Config) (*Agent, error) {
	agentTTL, err := time.ParseDuration(cfg.AgentConfig.AgentTTL)
	if err != nil {
		return nil, err
	}

	heartBeatInterval, err := time.ParseDuration(cfg.AgentConfig.HeartBeatInterval)
	if err != nil {
		return nil, err
	}

	eCfg := etcd.Config{
		Endpoints:               strings.Split(cfg.EtcdServers, ","),
		HeaderTimeoutPerRequest: (time.Duration(cfg.EtcdRequestTimeout*1000) * time.Millisecond),
		Username:                cfg.AgentConfig.EtcdUsername,
		Password:                cfg.AgentConfig.EtcdPassword,
	}

	eClient, err := etcd.New(eCfg)
	if err != nil {
		return nil, err
	}

	kAPI := etcd.NewKeysAPI(eClient)

	reg := ns.NewEtcdRegistry(kAPI, cfg.EtcdKeyPrefix)

	machine, err := NewMachineFromConfig(*reg, cfg)
	if err != nil {
		return nil, err
	}

	// heart beat
	hrt := heart.New(*reg, machine)
	// monitor
	mon := monitor.NewMonitor(agentTTL, heartBeatInterval)
	// task manager
	tm := taskManager.NewTaskManager(*reg, machine, cfg.AgentConfig.PkgSite, cfg.EtcdServers)
	// app health check manager
	hcm := healthCheck.NewHealthCheckManager(*reg, machine)

	srv := Agent{
		mach:   machine,
		hrt:    hrt,
		mon:    mon,
		tm:     tm,
		hcm:    hcm,
		reg:    *reg,
		killc:  make(chan struct{}),
		stopc:  make(chan struct{}),
		mstopc: make(chan struct{}),
	}

	return &srv, nil
}

func (s *Agent) Run() {
	log.Infof("Establishing etcd connectivity")

	var err error
	sleep := 2 * time.Second
	metric.BuildMappers()
	metric.UpdateCpuStat()
	metric.UpdateDiskStats()
	time.Sleep(sleep)
	metric.UpdateCpuStat()
	metric.UpdateDiskStats()

	for {
		_, err = s.hrt.Register(s.mon.TTL)
		if err == nil {
			log.Infof("hrt.Register() success")
			break
		}
		log.Warningf("Agent register machine failed: %v, retrying in %d sec.", err, sleep/time.Second)
		time.Sleep(sleep)
	}

	log.Infof("Starting Agent components")
	s.wg = sync.WaitGroup{}

	components := []func(){
		func() { s.mach.PeriodicRefresh(machineStateRefreshInterval, s.mstopc) },
		func() { s.tm.PeriodicSchedule(taskManagerGetInterval, s.stopc, s.mstopc) },
		func() { s.hcm.PeriodicHealthCheck(healthcheckManagerGetInterval, s.mstopc) },
	}

	for _, f := range components {
		f := f
		s.wg.Add(1)
		go func() {
			f()
			s.wg.Done()
		}()
	}

	go s.Supervise()
	s.ListenForSignals()
}

func (s *Agent) Supervise() {
	sd, err := s.mon.AgentBeat(s.hrt, s.killc)
	if sd {
		log.Infof("Agent monitor triggered: told to shut down")
	} else {
		log.Errorf("Agent monitor triggered: %v", err)
	}
	close(s.stopc)
	s.wg.Wait()
	os.Exit(0)
}

// Kill is used to gracefully terminate the Agent by triggering the Monitor to shut down
func (s *Agent) Kill() {
	log.Infof("Gracefully shutting down")
	close(s.killc)
}

func (s *Agent) ListenForSignals() {
	shutdown := func() {
		s.Kill()
	}

	sigchan := make(chan os.Signal, 1)
	signal_map := map[os.Signal]func(){
		syscall.SIGTERM: shutdown,
		syscall.SIGINT:  shutdown,
	}

	for k := range signal_map {
		signal.Notify(sigchan, k)
	}

	for true {
		sig := <-sigchan
		handler, ok := signal_map[sig]
		if ok {
			handler()
		}
	}
}
