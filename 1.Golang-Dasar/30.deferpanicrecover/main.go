package main

import "fmt"

func main() {
	// runAplication(10)
	// fmt.Println("Menjalankan fungsi lain")

	runApp(true)
	fmt.Println("Menjalankan fungsi lain")
}
func runApp(error bool) {
	defer endApp()
	if error {
		panic("APLIKASI ERROR")
	}

	fmt.Println("Aplikasi berjalan")
}

func endApp() {
	fmt.Println("Aplikasi Selesai")
	message := recover()
	if message != nil {
		fmt.Println("Error dengan message:", message)
	}
}

// func runAplication(value int){
// 	defer logging()

// 	fmt.Println("Running aplication")

// 	result := 10 / value
// 	fmt.Println("Result", result)

// }

// func logging() {
// 	fmt.Println("Selesai memanggil function")
// }
