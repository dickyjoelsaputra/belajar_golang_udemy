package main

import (
	database "empatlima/database"
	"fmt"
)

func main() {
	data := database.GetDatabase()
	fmt.Println(data)
}