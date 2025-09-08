package main

import (
	"container/list"
	"fmt"
)

func main() {
	var datas = list.New()

	datas.PushBack("Kopi")
	datas.PushBack("Teh")
	datas.PushBack("Mineral")

	for e := datas.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
