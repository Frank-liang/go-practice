package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	w := os.Stdout
	fmt.Fprint(w, "welcome  \n")
	for {
		s.Scan()
		msg := string(s.Bytes())
		if msg == "exit" {
			return
		}
		fmt.Fprintf(w, "You wrote %q\n", msg)
	}
}
