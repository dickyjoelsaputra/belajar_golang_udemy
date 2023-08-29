package main

import "fmt"

func main() {
	goodbye := getGoodBye
	fmt.Println(goodbye(""))
}

func getGoodBye(name string) string {
	return "Good Bye " + name
}