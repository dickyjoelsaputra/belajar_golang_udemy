package main

import (
	"container/list"
	"fmt"
)

func main() {
	data := list.New()

	data.PushBack("A")
	data.PushBack("B")
	data.PushBack("C")

	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}