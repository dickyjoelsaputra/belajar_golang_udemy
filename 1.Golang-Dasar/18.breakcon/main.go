package main

import "fmt"

func main() {
	// break berhenti perulangan

	// continui skip perulangan sampai kondisi terpenuhi

	// for i := 0; i < 10; i++ {
	// 	if i == 5 {
	// 		break
	// 	}

	// 	fmt.Println("Perulangan ke", i)
	// }

	for i := 0; i < 10; i++ {
		if i % 2 != 0 {
			continue
		}

		fmt.Println("perulangan ke", i)
	}
}