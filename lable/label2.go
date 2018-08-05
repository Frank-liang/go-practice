package main

import "fmt"

func main() {
	out := "byebye..."
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			for z := 0; z < 10; z++ {
				if z == 1 {
					goto end
				}
			}
		}
	}
end:
	fmt.Println(out)
}
