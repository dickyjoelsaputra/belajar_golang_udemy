package main

import "fmt"

func main() {
	data := getHello("")

	getHello("dck")

	fmt.Println(data)
}

func getHello(name string) string {
	if name == "" {
		return "Hello Bro"
	} else {
		return "Hello " + name
	}
}