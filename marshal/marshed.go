package main

import (
	"encoding/json"
	"fmt"
)

/*
如果给定一个JSON key 为Field

首先查找json tag名字为Field的字段
然后查找字段名字为Field的字段
最后查找字段名字为Field等大小写不敏感的匹配字段
如果没有找到，则忽略这个key， 不会报错
*/

type S struct {
	DiskIOPS string `json:"diskIOPS"`
}

// json.Unmashal() to S
type D struct {
	DiskIOPS string `db:"disk-IOPS"`
}

func main() {
	var d = D{
		DiskIOPS: "100",
	}
	var s = S{}
	jd, _ := json.Marshal(&d)
	fmt.Println(string(jd))
	err := json.Unmarshal(jd, &s)
	if err != nil {
		panic(err)
	}
	fmt.Println(s)
}
