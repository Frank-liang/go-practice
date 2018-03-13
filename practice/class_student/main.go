package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/51reboot/golang-01-homework/lesson6/jungle85gopy/stuServer/grade"
)

func main() {
	var grd grade.Grade
	fmt.Println("start:", grd)

	actionMap := map[string]func([]string) error{
		"create": grd.Create,
		"select": grd.Change,

		"add":    grd.Add,
		"list":   grd.List,
		"save":   grd.Save,
		"load":   grd.Load,
		"update": grd.Update,
		"delete": grd.Delete,
		"exit":   exit,
	}
	// saved = true
	fid := bufio.NewReader(os.Stdin)
	for {
		curClass := grd.GetcurClass()
		fmt.Printf("class: [%s] > ", curClass)

		cmd, args := parseCmd(fid)
		if cmd == "" {
			continue
		}
		fmt.Println("\targs:", args)
		actionFunc, ok := actionMap[cmd]
		if !ok {
			usage()
			continue
		}
		err := actionFunc(args)
		if err != nil {
			grade.PrintInfo(fmt.Sprintf("execute action [%s] error: %s", cmd, err))
			continue
		}
	}
	// return
}

func parseCmd(f *bufio.Reader) (string, []string) {
	// grade.printInfo(fmt.Sprintf("saved value: %v", saved))

	line, _ := f.ReadString('\n')
	line = strings.TrimSpace(line)
	argsAll := strings.Fields(line)
	if len(argsAll) == 0 {
		return "", nil
	} else if len(argsAll) == 1 {
		return argsAll[0], nil
	}
	return argsAll[0], argsAll[1:]
}

func usage() {
	grade.PrintInfo("cli usage:")

	fmt.Println("  + create name -- add class info")
	fmt.Println("  + select name -- change to class or create a new class")
	fmt.Println("  + add name id -- add student info")
	fmt.Println("  + list \t-- list student info")

	fmt.Println("  + update name id -- update student name by id")
	fmt.Println("  + delete name id -- delete student")
	fmt.Println("  + load file \t-- load student from file")
	fmt.Println("  + save file \t-- save student info file")
	fmt.Println("  + exit \t-- exit the cli without save")
}

func exit(args []string) error {
	// if saved || checkYes("exit with saving stu info") {
	os.Exit(0)
	// }
	return nil
}
