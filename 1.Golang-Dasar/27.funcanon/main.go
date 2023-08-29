package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You are blocked", name)
	}else{
		fmt.Println("Welcome", name)
	}
}

// func blackListName(name string) bool {
// 	return name == "admin"
// }
// func blackListRoot(name string) bool {
// 	return name == "root"
// }

func main() {
	blacklist := func(name string)bool{
		return name == "anjing"
	}

	registerUser("admin", blacklist)

	registerUser("root", func(name string) bool {
		return name == "root"
	})

	registerUser("anjing", blacklist)
}