package main

import (
	"fmt"
)

func getMaxConsecutiveOnes(N uint) uint {
	var maxCount uint = 0
	var count uint = 0
	var i uint
	for i = 0; i < 64; i++ {
		if (N>>i)&0x1 == 1 {
			count++
		} else {
			if count > maxCount {
				maxCount = count
			}
			count = 0
		}
	}

	return maxCount
}

func main() {

	var N uint
	fmt.Scanln(&N)
	fmt.Printf("%d\n", getMaxConsecutiveOnes(N))

}
