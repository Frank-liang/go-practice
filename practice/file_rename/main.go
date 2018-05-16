package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

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
	type file struct {
		name string
		path string
	}

	dir := "sample"
	var toRename []file
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if _, err := match(info.Name()); err == nil {
			toRename = append(toRename, file{
				name: info.Name(),
				path: path,
			})
		}
		return nil
	})
	for _, f := range toRename {
		fmt.Printf("%q\n", f)
	}
	for _, orig := range toRename {
		var n file
		var err error
		n.name, err = match(orig.name)
		if err != nil {
			fmt.Println("Error matching:", orig.path, err.Error())
		}
		n.path = filepath.Join(dir, n.name)
		fmt.Printf("mv %s => %s\n", orig.path, n.path)
		err = os.Rename(orig.path, n.path)
		if err != nil {
			fmt.Println("Error renaming:", orig.path, err.Error())
		}
	}

}

//match returns the new name, or an error if the file name
//didn't match our pattern
func match(fileName string) (string, error) {
	//"birthday", "001", "txt"
	pieces := strings.Split(fileName, ".")
	ext := pieces[len(pieces)-1]
	tmp := strings.Join(pieces[0:len(pieces)-1], ".")
	pieces = strings.Split(tmp, "_")
	name := strings.Join(pieces[0:len(pieces)-1], "_")
	number, err := strconv.Atoi(pieces[len(pieces)-1])
	if err != nil {
		return "", fmt.Errorf("%s didn't match our pattern", fileName)
	}
	// Birthday - 1.txt
	return fmt.Sprintf("%s - %d.%s", strings.Title(name), number, ext), nil

}
