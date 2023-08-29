package main

import "fmt"

func main() {
	name := "s"

	if name == "dicky" {
		fmt.Println("Hello dicky")
	}else if name == "dicksy"{
		fmt.Println("Hello dicksy")
	}else{
		fmt.Println("Hello bukan dicky")
	}

	if length := len(name); length > 5 {
		fmt.Println("Terlalu panjang")
	}
}