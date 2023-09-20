package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	dicky := Man{"Dicky"}
	dicky.Married()

	fmt.Println(dicky)
}
