package main

import (
	"fmt"
	"strings"
)

func main() {
	var jul string = "Dicky Joel Saputra"
	// strings.Count(jul)
	fmt.Println("clone : ", strings.Clone(jul))
	fmt.Println("compare : ", strings.Compare("Dicky", jul))
	fmt.Println("contains : ", strings.Contains(jul, "Dicky"))
	fmt.Println("containsAny" , strings.ContainsAny(jul, "D"))
	fmt.Println("containsRune : ", strings.ContainsRune(jul, 'D'))
	fmt.Println("count : ", strings.Count(jul, "a"))
	fmt.Println("equalFold : ", strings.EqualFold(jul, "Dicky Joel Saputra"))
	fmt.Println("fields : ", strings.Fields(jul))
	// strings.
}
