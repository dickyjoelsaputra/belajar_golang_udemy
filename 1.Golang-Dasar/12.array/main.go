package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "Raisa"
	names[1] = "Andriana"
	names[2] = "Saputri"

	println(names[0])

	var values = [3]any{
		"9",
		95,
		80,
	}

	fmt.Println(values)

	fmt.Println(len(names))
	fmt.Println(len(values))

}