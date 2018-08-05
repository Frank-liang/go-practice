package config

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"path"
	"strings"

	log "github.com/sirupsen/logrus"
)

var (
	VersionStr = "unknown"
	Cfg        Config
)

const AGNET_PATH = "agent"
const ETCD_PATH = "etcd"

type Config struct {
	IsDebug                 bool
	EtcdServers             string
	EtcdKeyPrefix           string
	EtcdRequestTimeout      float64
	EngineReconcileInterval float64
	Verbosity               int
	RawMetadata             string
	UnitsDirectory          string
	Hostname                string
	ConfigServerURI         string
	InstanceRegIP           string
	HomeDirectory           string
	AgentConfig             AgentConfig
}

type AgentConfig struct {
	AgentTTL          string  `json:"agent-ttl"`
	CpuCores          int     `json:"cpu"`
	DiskCapacity      float64 `json:"disk"`
	Memory            float64 `json:"memory"`
	EtcdPassword      string  `json:"etcd_password"`
	EtcdUsername      string  `json:"etcd_username"`
	HeartBeatInterval string  `json:"heartbeat_inverval"`
	PkgSite           string  `json:"pkg_site"`
	ListenPort        int     `json:"port"`
	PublicIP          string  `json:"public_ip"`
}

func (c *Config) Metadata() map[string]string {
	meta := make(map[string]string, 0)

	for _, pair := range strings.Split(c.RawMetadata, ",") {
		parts := strings.SplitN(pair, "=", 2)
		if len(parts) != 2 {
			continue
		}

		key := strings.TrimSpace(parts[0])
		val := strings.TrimSpace(parts[1])

		meta[key] = val
	}

	return meta
}

func (c *Config) GetCfgCenterConfig() (err error) {
	Cfg = *c
	cfg_env := "/home/work/"
	agent_cfg_path := path.Join(cfg_env, AGNET_PATH, c.InstanceRegIP)
	etcd_cfg_path := path.Join(cfg_env, ETCD_PATH)

	agent_cfg_url := c.ConfigServerURI + agent_cfg_path
	res, err := http.Get(agent_cfg_url)
	if err != nil {
		log.Errorf("HTTP request %s failed: %v", agent_cfg_url, err)
	}
	if res.Body != nil {
		defer res.Body.Close()
	}
	conf_str, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Errorf("HTTP read %s failed: %v", agent_cfg_url, err)
	}

	etcd_cfg_url := c.ConfigServerURI + etcd_cfg_path
	res, err = http.Get(etcd_cfg_url)
	if err != nil {
		log.Errorf("HTTP request %s failed: %v", etcd_cfg_url, err)
	}
	if res.Body != nil {
		defer res.Body.Close()
	}
	etcd_conf_str, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Errorf("HTTP read %s failed: %v", etcd_cfg_url, err)
	}

	AgentConfig := AgentConfig{}
	err = json.Unmarshal(conf_str, &AgentConfig)

	if c.AgentConfig.AgentTTL == "20s" && AgentConfig.AgentTTL != "" {
		c.AgentConfig.AgentTTL = AgentConfig.AgentTTL
	}
	if c.AgentConfig.PublicIP == "" && AgentConfig.PublicIP != "" {
		c.AgentConfig.PublicIP = AgentConfig.PublicIP
	}
	if c.AgentConfig.CpuCores == 0 && AgentConfig.CpuCores != 0 {
		c.AgentConfig.CpuCores = AgentConfig.CpuCores
	}
	if c.AgentConfig.DiskCapacity == 0 && AgentConfig.DiskCapacity != 0 {
		c.AgentConfig.DiskCapacity = AgentConfig.DiskCapacity
	}
	if c.AgentConfig.Memory == 0 && AgentConfig.Memory != 0 {
		c.AgentConfig.Memory = AgentConfig.Memory
	}
	if c.AgentConfig.EtcdPassword == "" && AgentConfig.EtcdPassword != "" {
		c.AgentConfig.EtcdPassword = AgentConfig.EtcdPassword
	}
	if c.AgentConfig.EtcdUsername == "" && AgentConfig.EtcdPassword != "" {
		c.AgentConfig.EtcdUsername = AgentConfig.EtcdUsername
	}
	if c.AgentConfig.HeartBeatInterval == "5s" && AgentConfig.HeartBeatInterval != "" {
		c.AgentConfig.HeartBeatInterval = AgentConfig.HeartBeatInterval
	}
	if c.AgentConfig.ListenPort == 2110 && AgentConfig.ListenPort != 0 {
		c.AgentConfig.ListenPort = AgentConfig.ListenPort
	}

	if c.EtcdServers == "" && etcd_conf_str != nil {
		err = json.Unmarshal(etcd_conf_str, &c.EtcdServers)
	}

	return nil
}
