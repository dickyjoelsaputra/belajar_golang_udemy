package main

import "fmt"

func main() {
	name := "dickysss"

	switch name {
	case "dicky":
		fmt.Println("Hello dicky")
	case "dicksy":
		fmt.Println("Hello dicksy")
	case "joko":
		fmt.Println("Hello joko")
	default:
		fmt.Println("Hello aja dech")
	}


	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Terlalu panjang")
	case false:
		fmt.Println("Nama sudah benar")
	}

	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Terlalu panjang")
	case length < 5:
		fmt.Println("Teralu pendek")
	default:
		fmt.Println("Nama sudah benar")
	}
}
