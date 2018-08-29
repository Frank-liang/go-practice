package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func excludeName(name string, exclude string) bool {
	if exclude == "" {
		return false
	}

	if filepath.Base(name) == exclude {
		return true
	}
	return false
}

func main() {
	mS := flag.Bool("s", false, "Sockets")
	mP := flag.Bool("p", false, "Pipes")
	mSL := flag.Bool("sl", false, "Symbolic Links")
	mD := flag.Bool("d", false, "Directories")
	mF := flag.Bool("f", false, "files")
	mX := flag.String("x", "", "Files")

	flag.Parse()
	flags := flag.Args()

	printAll := false
	if *mS && *mP && *mSL && *mD && *mF {
		printAll = true
	}

	if !(*mS || *mP || *mSL || *mD || *mF) {
		printAll = true
	}

	if len(flags) == 0 {
		fmt.Println("Not enough arguments ")
		os.Exit(1)
	}

	Path := flags[0]

	walkFunction := func(path string, info os.FileInfo, err error) error {
		fileInfo, err := os.Stat(path)
		if err != nil {
			return err
		}

		if excludeName(Path, *mX) {
			return nil
		}

		if printAll {
			fmt.Println(path)
			return nil
		}

		mode := fileInfo.Mode()
		if mode.IsRegular() && *mF {
			fmt.Println(path)
			return nil
		}

		if mode.IsDir() && *mD {
			fmt.Println(path)
			return nil
		}

		fileInfo, _ = os.Lstat(path)

		if fileInfo.Mode()&os.ModeSymlink != 0 {
			if *mSL {
				fmt.Println(path)
				return nil
			}
		}

		if fileInfo.Mode()&os.ModeNamedPipe != 0 {
			if *mP {
				fmt.Println(path)
				return nil
			}
		}

		if fileInfo.Mode()&os.ModeSocket != 0 {
			if *mS {
				fmt.Println(path)
				return nil
			}
		}

		return nil
	}

	err := filepath.Walk(Path, walkFunction)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
