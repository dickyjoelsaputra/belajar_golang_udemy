package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "Dicky Joel Saputra",
		"address": "Indonesia",
	}
	person["title"] = "Programmer"

	// fmt.Println(person)
	// fmt.Println(person["name"])
	// fmt.Println(person["address"])
	// fmt.Println(person["title"])
	delete(person, "address")
	fmt.Println(len(person))
}