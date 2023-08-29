package main

import "fmt"

func main() {
	data1 , _ , data3 := getFullName()
	fmt.Println(data1 , data3)
}

func getFullName()(string , string , string){
	return "Dicky" , "Joel" , "Saputra"
}