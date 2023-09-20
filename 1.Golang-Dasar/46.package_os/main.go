package main

import (
	"fmt"
	"os"
)

func main() {
	// go run main.go Dicky Joel
	data := os.Args
	fmt.Println(data)
	fmt.Println(data[1])
	fmt.Println(data[2])

	name , err := os.Hostname()
	if err != nil {
		fmt.Println("Error", err.Error())
	}

	fmt.Println("Hostname : ", name)
}
