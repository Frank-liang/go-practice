package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func tripletToBinary(triplet string) string {
	if triplet == "rwx" {
		return "111"
	}
	if triplet == "-wx" {
		return "011"
	}
	if triplet == "--x" {
		return "001"
	}
	if triplet == "---" {
		return "000"
	}
	if triplet == "r-x" {
		return "101"
	}
	if triplet == "r--" {
		return "100"
	}
	if triplet == "--x" {
		return "001"
	}
	if triplet == "rw-" {
		return "110"
	}
	if triplet == "-w-" {
		return "101"
	}
	return "unknown"
}

func convertToBinary(permissions string) string {
	binaryPermissions := permissions[1:]
	p1 := binaryPermissions[0:3]
	p2 := binaryPermissions[3:6]
	p3 := binaryPermissions[6:9]
	return tripletToBinary(p1) + tripletToBinary(p2) + tripletToBinary(p3)
}

func main() {
	arguements := os.Args
	if len(arguements) == 1 {
		fmt.Printf("Usage: %s filename\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	filename := arguements[1]
	info, _ := os.Stat(filename)
	mode := info.Mode()

	fmt.Println(filename, "mode", mode)
	fmt.Println("As strings is", mode.String()[1:10])
	fmt.Println("As binary is", convertToBinary(mode.String()))
}
