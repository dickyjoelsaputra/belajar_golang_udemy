package main

import "fmt"

func main() {
	result := random()
	resultString := result.(string)
	fmt.Println(resultString)

}

func random() interface{} {
	return "Ups"
}
