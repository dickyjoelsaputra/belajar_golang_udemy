package main

import "fmt"

func main() {
	bulan := [...]string{"januari", "februari", "maret", "april", "mei", "juni", "juli", "agustus", "september", "oktober", "november", "desember"}

	slice1 := bulan[4:7]
	slice2 := bulan[6:9]

	fmt.Println("slice1", slice1)
	fmt.Println("len slice1", len(slice1))
	fmt.Println("cap slice1", cap(slice1))

	fmt.Println("slice2", slice2)
	fmt.Println("len slice2", len(slice2))
	fmt.Println("cap slice2", cap(slice2))

	newSlice := make([]string , 2 , 5)
	newSlice[0] = "Muhammad"
	newSlice[1] = "Zakaria"

	fmt.Println(newSlice)
}