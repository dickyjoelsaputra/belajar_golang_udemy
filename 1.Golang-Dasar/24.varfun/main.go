package main

import "fmt"

func main() {
	total := sumAll(10, 10, 10, 10, 10 , 10)
	fmt.Println(total)

	slice := []int{10, 10, 10, 10, 10}
	totals := sumAll(slice...)
	fmt.Println(totals)
}

func sumAll(numbers ...int) int {
	total := 0
	for _, value := range numbers {
		total += value
	}

	return total
}