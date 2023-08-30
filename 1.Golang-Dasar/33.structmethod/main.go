package main

import "fmt"

type Costumer struct {
	Name, Address string
	Age           int
}

func main() {
	dicky := Costumer{Name: "Dicky", Address: "Indonesia", Age: 20}
	dicky.sayHello()
}

func (c Costumer) sayHello() {
	fmt.Println("Hello", c.Name)
}