package main

import "fmt"

func main() {
	// counter := 1

	// for counter <= 10 {
	// 	fmt.Println("Perulangan ke", counter)
	// 	counter++
	// }

	// for i := 0; i <= len("titid"); i++ {
	// 	fmt.Println("Perulangan ke", i)
	// }

	// slice := []string{"dicky", "joel", "saputra"}

	// for i := 0; i < len(slice); i++ {
	// 	fmt.Println("Perulangan ke ", slice[i] , i)
	// }

	// for _, v := range slice {
	// 	fmt.Println(v)
	// }

	maps := map[any]any{
		"name":  "Dicky Joel Saputra",
		"titid": 20,
		"hobby": []string{"coding", "gaming"},
	}

	for i, v := range maps {
		fmt.Println(i, "=", v)

		if i == "hobby" {
			for _, v2 := range v.([]string) {
				fmt.Println("hobby", v2)
			}
		}
	}
}
