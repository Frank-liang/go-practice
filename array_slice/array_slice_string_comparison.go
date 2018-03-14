package main

import "fmt"

//slice
func main() {
	fmt.Println(keys(map[string]struct{}{
		"dog":       struct{}{},
		"cat":       struct{}{},
		"mouse":     struct{}{},
		"wolf":      struct{}{},
		"alligator": struct{}{},
	}))
}

func keys(m map[string]struct{}) []string {
	var ret []string
	fmt.Println(cap(ret))
	for key := range m {
		ret = append(ret, key)
		fmt.Println(cap(ret))
	}
	return ret
}

//array
/*func main() {
	fmt.Println(keys(map[string]struct{}{
		"dog":       struct{}{},
		"cat":       struct{}{},
		"mouse":     struct{}{},
		"wolf":      struct{}{},
		"alligator": struct{}{},
	}))
}

func keys(m map[string]struct{}) []string {
	ret := make([]string, 0, len(m))
	fmt.Println(cap(ret))
	for key := range m {
		ret = append(ret, key)
		fmt.Println(cap(ret))
	}
	return ret
}*/
//slice array array 预分配内存，性能稍微好点。 slice动态的分配内存。如果知道量的话，最好是预分配内存
