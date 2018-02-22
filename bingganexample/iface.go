package main

import (
	"fmt"
	"net/http"
)

// 名字: IInstance
// 方法: Instance() float64

type IInstance interface {
	Instance() float64
}

type Path []Point

func (p Path) Instance() float64 {

}

func handle(w http.ResponseWriter, r *httpRequest) {
	fmt.Fprintf(w, "hello\n")
}

func main() {
	var i IInstance
	p := Path{{1, 2}, {3, 4}}
	i = p
	fmt.Println(i)
}
