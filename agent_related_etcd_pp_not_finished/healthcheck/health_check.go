package healthCheck

import (
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
	"syscall"

	log "github.com/sirupsen/logrus"
	"github.com/wusuopubupt/go-slides/lessons/lesson06/agent/config"
)

type addrResult struct {
	addr   string
	opened bool
}

func IsAppHealthy(dirName string) bool {
	stateCommandPath := filepath.Join(config.Cfg.HomeDirectory, dirName, "bin/state.sh")
	// first check if the file exist, if not, return false
	if _, err := os.Stat(stateCommandPath); os.IsNotExist(err) {
		return false
	}

	exec.Command("chmod", "755", stateCommandPath).Run()

	output, _ := exec.Command(stateCommandPath).Output()
	state := strings.TrimSpace(string(output[:]))
	log.Debugf("Current service %s app state is %s ", dirName, state)

	return strings.EqualFold(state, "serving") || strings.EqualFold(state, "loading")
}

func PortsActive(addrs []string) map[string]bool {
	var waitGroup sync.WaitGroup
	c := make(chan addrResult, len(addrs))
	results := make(map[string]bool)

	for _, addr := range addrs {
		waitGroup.Add(1)
		go tcpConnection(addr, c, &waitGroup)
	}

	waitGroup.Wait()
	for range addrs {
		result := <-c
		results[result.addr] = result.opened
	}

	return results
}

func PortActive(addr string) bool {
	addrs := []string{addr}
	result := PortsActive(addrs)
	return result[addr]
}

func tcpConnection(addr string, c chan addrResult, waitGroup *sync.WaitGroup) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		c <- addrResult{addr, false}
		waitGroup.Done()
		return
	}
	defer conn.Close()

	c <- addrResult{addr, true}
	waitGroup.Done()
}

func ProcessActive(pid int) bool {
	process, err := os.FindProcess(pid)
	if err != nil {
		return false
	}
	err = process.Signal(syscall.Signal(0))
	// if err is nil, then the process is running
	if err != nil {
		return false
	}
	return true
}
