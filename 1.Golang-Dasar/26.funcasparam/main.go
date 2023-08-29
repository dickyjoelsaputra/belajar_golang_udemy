package main

import (
	"fmt"
)

type Filter func(int) int

func main() {
	name := "Anjing"

	sayHelloWithFilter(name, spamFilter)

	angka := 0
	filterAngka(angka, angkaGanjilOnly)
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}

func sayHelloWithFilter(name string, filter func(string) string) {
	nameFiltered := filter(name)
	fmt.Println("Hello", nameFiltered)
}

func filterAngka(angka int, filter Filter) {
	angkaFiltered := filter(angka)
	fmt.Println(angkaFiltered)
}

func angkaGanjilOnly(angka int)int , error{
	if angka % 2 != 1 {
		return angka , nil
	} else {
		return 0 , errors
	}
}