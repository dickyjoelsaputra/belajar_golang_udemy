package main

import "fmt"

func main() {

	data1, data2, data3 := getFullName()

	fmt.Println(data1 , data2 , data3)
}

func getFullName() (firstName, middleName, lastName string) {
	firstName = "Dicky"
	middleName = "Joel"
	lastName = "Saputra"
	return
}