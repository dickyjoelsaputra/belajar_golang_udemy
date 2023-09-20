package main

import "fmt"

func main() {
	data := NewMap("")
	fmt.Println(data)
}

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}