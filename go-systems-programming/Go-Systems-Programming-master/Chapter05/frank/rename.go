package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	overWrite := flag.Bool("overwrite", false, "overwrite")

	flag.Parse()
	flags := flag.Args()

	if len(flags) < 2 {
		fmt.Println("Please provide two argument")
		os.Exit(1)
	}

	source := flags[0]
	destination := flags[1]
	fileInfo, err := os.Stat(source)
	if err == nil {
		mode := fileInfo.Mode()
		if mode.IsRegular() == false {
			fmt.Println("Sorry, we only support regular files")
			os.Exit(1)
		}
	} else {
		fmt.Println("Error reading:", source)
		os.Exit(1)
	}

	newDestination := destination
	destInfo, err := os.Stat(destination)
	if err == nil {
		mode := destInfo.Mode()
		if mode.IsDir() {
			justThename := filepath.Base(source)
			newDestination = destination + "/" + justThename
		}
	}

	destination = newDestination
	destInfo, err = os.Stat(destination)
	if err == nil {
		if *overWrite == false {
			fmt.Println("Destination file already exists")
			os.Exit(1)
		}
	}

	err = os.Rename(source, destination)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
