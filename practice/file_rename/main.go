package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

var re = regexp.MustCompile("^(.+?) ([0-9]{4}) [(]([0-9]+) of ([0-9]+)[)][.](.+?)$")
var replaceString = "$2 - $1 - $3 of $4.$5"

func main() {
	// 1: Just for test
	//	fileName := "birthday_001.txt"
	//	// => Birthday - 1 of 4.txt
	//	newName, err := match(fileName, 4)
	//	if err != nil {
	//		fmt.Println("No match")
	//		os.Exit(1)
	//	}
	//	fmt.Println(newName)

	// 2: Only for a dir, not a recursive way
	//dir := "sample"
	//files, err := ioutil.ReadDir(dir)
	//if err != nil {
	//	panic(err)
	//}
	//count := 0
	//var toRename []string
	//for _, file := range files {
	//	if file.IsDir() {
	//		fmt.Println("Dir: ", file.Name())
	//	} else {
	//		_, err := match(file.Name(), 0)
	//		if err == nil {
	//			count++
	//			toRename = append(toRename, file.Name())
	//		}
	//	}
	//}
	//for _, origFilename := range toRename {
	//	origPath := filepath.Join(dir, origFilename)
	//	newFilename, err := match(origFilename, count)
	//	if err != nil {
	//		panic(err)
	//	}
	//	newPath := filepath.Join(dir, newFilename)
	//	fmt.Printf("mv %s => %s\n", origPath, newPath)
	//	err = os.Rename(origPath, newPath)
	//	if err != nil {
	//		panic(err)
	//	}
	//}

	//3
	var dry bool
	flag.Bool("dry", true, "whether or not this should be a real or dry run")
	flag.Parse()

	walkDir := "sample"
	var toRename []string
	filepath.Walk(walkDir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if _, err := match(info.Name()); err == nil {
			toRename = append(toRename, path)
		}
		return nil
	})

	for _, oldPath := range toRename {
		dir := filepath.Dir(oldPath)
		filename := filepath.Base(oldPath)
		newFilename, _ := match(filename)
		newPath := filepath.Join(dir, newFilename)
		fmt.Printf("mv %s => %s\n", oldPath, newPath)
		if !dry {
			err := os.Rename(oldPath, newPath)
			if err != nil {
				fmt.Println("Error renaming:", oldPath, newPath, err.Error())
			}
		}
	}

}

//match returns the new name, or an error if the file name
//didn't match our pattern
// 1 This func doesn't use regexp
//func match(fileName string) (string, error) {
//	//"birthday", "001", "txt"
//	pieces := strings.Split(fileName, ".")
//	ext := pieces[len(pieces)-1]
//	tmp := strings.Join(pieces[0:len(pieces)-1], ".")
//	pieces = strings.Split(tmp, "_")
//	name := strings.Join(pieces[0:len(pieces)-1], "_")
//	number, err := strconv.Atoi(pieces[len(pieces)-1])
//	if err != nil {
//		return "", fmt.Errorf("%s didn't match our pattern", fileName)
//	}
//	// Birthday - 1.txt
//	return fmt.Sprintf("%s - %d.%s", strings.Title(name), number, ext), nil
//
//}

func match(filename string) (string, error) {
	if !re.MatchString(filename) {
		return "", fmt.Errorf("%s didn't match our pattern", filename)
	}
	return re.ReplaceAllString(filename, replaceString), nil
}
