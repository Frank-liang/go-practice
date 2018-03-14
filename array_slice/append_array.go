package main

import "fmt"

func main() {
	vals := make([]int, 5)
	for i := 0; i < 5; i++ {
		vals = append(vals, i)
	}
	fmt.Println(vals)
}

//结果是[0 0 0 0 0 0 1 2 3 4] 此数组，如果不指定cap（最终容量），就默认和len（初始容量）一样. append的作用再是数组的末尾添加一些
//值,而vals 的初始默认值为[0 0 0 0 0], 所以结果如上
//想实现预想的结果,可以如下方法

/*vals := make([]int, 5)
for i := 0; i < 5; i++ {
	  vals[i] = i
  }
  fmt.Println(vals)*/
