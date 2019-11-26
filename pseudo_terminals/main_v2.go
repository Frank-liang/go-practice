package main

import (
	"fmt"
	"io"
	"math/rand"
	"os"
)

var cmdFunc func(w io.Writer, args []string) (exit bool)

func main() {

	switch cmd {
	case "exit":
		cmdFunc = exitCmd
	}

	if cmdFunc == nil {
		fmt.Fprintf(w, "%q not found\n", cmd)
		continue
	}

	if cmdFunc(w, args) {
		return
	}

}

func exitCmd(w io.Writer, args []string) bool {
	fmt.Fprintf(w, "Goodbye! :)")
	return true
}

func shuffle(w io.Writer, args ...string) bool {
	rand.Shuffle(len(args), func(i, j int) {
		args[i], args[j] = args[j], args[i]
	})

	for i := range args {
		if i > 0 {
			fmt.Fprint(w, "")
		}

		fmt.Fprintf(W, "%s", args[i])
	}
	fmt.Fprintln(w)
	return false
}

func print(w io.Writer, args ...string) bool {
	if len(args) != 1 {
		fmt.Fprintln(w, "Please specify one file!")
		return false
	}
	f, err := os.Open(args[0])
	if err != nil {
		fmt.Fprintf(w, "Cannot open %s: %s\n", args[0], err)
	}
	fmt.Fprintln(w)
	return false
}
