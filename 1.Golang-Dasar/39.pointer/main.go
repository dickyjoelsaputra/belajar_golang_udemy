package main

import "fmt"

type Address struct {
	Kota     string
	Provinsi string
	Negara   string
}

func main() {
	// address1 := Address{"Jakarta", "Dki Jakarta", "Indonesia"}
	// address2 := &address1

	// address2.Kota = "Bandung"

	// *address2 = Address{"NTT" , "POJOK" , "KIRI"}

	// log.Printf("%v" , address1)
	// log.Printf("%v" , *address2)

	var numberA int = 4
	var numberB *int = &numberA

	fmt.Println("numberA (value)   :", numberA)  // 4
	fmt.Println("numberA (address) :", &numberA) // 0xc20800a220

	fmt.Println("numberB (value)   :", *numberB) // 4
	fmt.Println("numberB (address) :", numberB)  // 0xc20800a220
}
