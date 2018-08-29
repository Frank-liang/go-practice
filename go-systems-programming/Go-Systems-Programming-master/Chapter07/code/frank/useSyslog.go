package main

import (
	"fmt"
	"log"
	"log/syslog"
	"os"
	"path/filepath"
)

func main() {
	programName := filepath.Base(os.Args[0])
	sysLog, e := syslog.New(syslog.LOG_INFO|syslog.LOG_LOCAL7, programName)
	if e != nil {
		log.Fatal(e)
	}
	sysLog.Crit("Crit: Logging in Go")
	fmt.Fprintf(sysLog, "log:Print: Logging in Go!")
}
