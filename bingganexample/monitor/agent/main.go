package main

import (
	"bufio"
	"flag"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/51reboot/golang-01-homework/lesson12/binggan/monitor/common"
	"github.com/BurntSushi/toml"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/load"
)

func NewMetric(metric string, value float64) *common.Metric {
	hostname, _ := os.Hostname()
	return &common.Metric{
		Metric:    metric,
		Endpoint:  hostname,
		Value:     value,
		Tag:       []string{runtime.GOOS},
		Timestamp: time.Now().Unix(),
	}
}

func CpuMetric() []*common.Metric {
	var ret []*common.Metric
	cpus, err := cpu.Percent(time.Second, false)
	if err != nil {
		panic(err)
	}
	metric := NewMetric("cpu.usage", cpus[0])
	ret = append(ret, metric)

	cpuload, err := load.Avg()
	if err == nil {
		metric = NewMetric("cpu.load1", cpuload.Load1)
		ret = append(ret, metric)
		metric = NewMetric("cpu.load5", cpuload.Load5)
		ret = append(ret, metric)
	}

	return ret
}

func NewUserMetric(cmdstr string) MetricFunc {
	return func() []*common.Metric {
		metrics, err := getUserMetrics(cmdstr)
		if err != nil {
			log.Print(err)
			return []*common.Metric{}
		}
		return metrics
	}
}

func getUserMetrics(cmdstr string) ([]*common.Metric, error) {
	// 构建命令
	// 获取标准输出
	// 按行解析
	// 获取key ,value
	// 包装成common.Metric
	var ret []*common.Metric
	cmd := exec.Command("bash", "-c", cmdstr)
	stdout, _ := cmd.StdoutPipe()

	err := cmd.Start()
	if err != nil {
		return nil, err
	}

	r := bufio.NewReader(stdout)
	for {
		line, err := r.ReadString('\n')
		if err != nil {
			break
		}
		line = strings.TrimSpace(line)
		fields := strings.Fields(line)
		if len(fields) != 2 {
			continue
		}
		key, value := fields[0], fields[1]
		n, err := strconv.ParseFloat(value, 64)
		if err != nil {
			log.Print(err)
			continue
		}
		metric := NewMetric(key, n)
		ret = append(ret, metric)
	}

	return ret, nil
}

func main() {
	flag.Parse()

	_, err := toml.DecodeFile(*configPath, &gcfg)
	if err != nil {
		log.Fatal(err)
	}

	sender := NewSender(gcfg.Sender.TransAddr)
	ch := sender.Channel()

	sched := NewSched(ch)

	sched.AddMetric(CpuMetric, 3*time.Second)
	for _, ucfg := range gcfg.UserScript {
		sched.AddMetric(NewUserMetric(ucfg.Path), time.Duration(ucfg.Step)*time.Second)
	}
	// memory, time.Second * 3
	// disk, time.Minute
	sender.Start()
}
