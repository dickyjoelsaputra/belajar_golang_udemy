package main

import "fmt"

func main() {
	var kosong interface{} = Ups(1)
	fmt.Println(kosong.(int))
}

func Ups(param int) interface{} {
	if param == 1 {
		return "Ups"
	} else if param == 2 {
		return 2
	} else {
		return true
	}
}