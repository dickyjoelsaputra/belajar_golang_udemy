package main

import "fmt"

type Costumer struct {
	Name, Address string
	Age           int
}

func main() {
	var dicky Costumer
	dicky.Name = "Dicky"
	dicky.Address = "Indonesia"
	dicky.Age = 20

	fmt.Println(dicky)
	fmt.Println(dicky.Name)

	joko := Costumer{
		Name:   "Joko",
		Age:    20,
		Address: "Indonesia",
	}

	fmt.Println(joko)

	budi := Costumer{"Budi", "Indonesia", 20}
	fmt.Println(budi)
}
