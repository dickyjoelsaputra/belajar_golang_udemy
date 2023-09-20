package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now)

	date := time.Date(2018, time.November, 1, 0, 0, 0, 0, time.Local)
	fmt.Println(date)

	parse , _ := time.Parse("2006-01-02 15:04:05", "2018-11-01 22:50:11")
	fmt.Println(parse)
}