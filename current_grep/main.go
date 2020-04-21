package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
)

type Job struct {
	filename string
	results  chan<- Result
}

type Result struct {
	filename string
	lino     int
	line     string
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	if len(os.Args) < 3 || os.Args[1] == "-h" || os.Args[1] == "--help" {
		fmt.Printf("Usage: %s <regrxp> <file>\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}
	if lineRx, err := regexp.Compile(os.Args[1]); err != nil {
		log.Fatalf("invalid regexp: %s\n", err)
	} else {
		//grep(lineRx, commandLineFiles(os.Args[2:]))
	}

}

func commandLineFiles(files []string) []string {
	args := make([]string, 0, len(files))
	for _, name := range files {
		if matches, err := filepath.Glob(name); err != nil {
			args = append(args, name)
		} else if matches != nil {
			args = append(args, matches...)
		}
		return args
	}
	return files
}

var workers = runtime.NumCPU()

/*func grep(lineRx *regexp.Regexp, filenames []string){
	jobs := make(chan Job, workers)
	results := make(chan Result, mininum(1000, len(filenames)))


}
*/

func mininum(first int, rest ...int) int {
	for _, x := range rest {
		if x < first {
			first = x
		}
	}
	return first
}
