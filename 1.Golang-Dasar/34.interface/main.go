package main

import (
	"fmt"
	"math"
)

func main() {
	// dicky := Person{Name: "Dicky"}
	// SayHello(dicky)

	// dog := Animal{Name: "Doggy"}
	// fmt.Println(dog.GetLegs())
	// SayHello(dog)

	var bangundatar Hitung

	bangundatar = Lingkaran{diameter: 10.0}
	fmt.Println("Luas Lingkaran\t\t:", bangundatar.Luas())
	fmt.Println("Keliling Lingkaran\t:", bangundatar.Keliling())
	fmt.Println("Jari-jari Lingkaran\t:", bangundatar.(Lingkaran).JariJari())
	fmt.Println("Asal Lingkaran\t\t:", bangundatar.(Lingkaran).Asal())
}

type Hitung interface {
	Luas() float64
	Keliling() float64
}

type Lingkaran struct{
	diameter float64
}

func (l Lingkaran) Asal() float64 {
	return l.diameter / 2
}

func (l Lingkaran) JariJari() float64 {
	return l.diameter * 2
}

func (l Lingkaran) Luas() float64 {
    return math.Pi * math.Pow(l.JariJari(), 2)
}

func (l Lingkaran) Keliling() float64 {
	return math.Pi * l.diameter
}


type Animal struct {
	Name string
}

func (a Animal) GetName() string {
	return a.Name
}

func (a Animal) GetLegs() int {
	return 4
}

type HasName interface {
	GetName() string
}

func SayHello(hasName HasName) {
	fmt.Println("Hello", hasName.GetName())
}

// type Person struct {
// 	Name string
// }

// func (person Person) Getname() string {
// 	return person.Name
// }
