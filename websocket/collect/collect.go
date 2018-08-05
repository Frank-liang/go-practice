package collect

import (
	"bufio"
	"fmt"
	"io"
	"os/exec"
	"strings"
)

func GetPS() ([]byte, error) {
	cmd := exec.Command("ps", "aux")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	cmd.Start()
	reader := bufio.NewReader(stdout)
	arr := make([]string, 0)
	for {
		line, err2 := reader.ReadString('\n')
		if err2 != nil || io.EOF == err2 {
			break
		}
		arr = append(arr, line)
	}
	s := strings.Join(arr, "<br/>")
	return []byte(s), nil
}
