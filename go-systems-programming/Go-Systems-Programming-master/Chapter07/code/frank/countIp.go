package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	minusCOL := flag.Int("COL", 1, "Column")
	flag.Parse()
	flags := flag.Args()

	if len(flags) == 0 {
		fmt.Printf("usage: %s <file1> [<file2> [...<fileN>]]\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	column := *minusCOL
	if column < 0 {
		fmt.Println("Invalid Column number")
		os.Exit(1)
	}

	myIps := make(map[string]int)
	for _, filename := range flags {
		fmt.Println("\t\t", filename)
		f, err := os.Open(filename)
		if err != nil {
			fmt.Printf("error opening file %s\n", err)
			continue
		}
		defer f.Close()

		r := bufio.NewReader(f)
		for {
			line, err := r.ReadString('\n')

			if err == io.EOF {
				break
			} else if err != nil {
				fmt.Printf("error reading file %s", err)
				continue
			}

			data := strings.Fields(line)
			ip := data[column-1]
			trail := net.ParseIP(ip)
			if trail.To4() == nil {
				continue
			}

			_, ok := myIps[ip]
			if ok {
				myIps[ip] = myIps[ip] + 1
			} else {
				myIps[ip] = 1
			}
		}

	}

	for key, _ := range myIps {
		fmt.Printf("%s %d\n", key, myIps[key])
	}
}
