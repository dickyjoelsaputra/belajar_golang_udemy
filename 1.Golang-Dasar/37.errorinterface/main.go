package main

import (
	"errors"
	"fmt"
)

func main() {
	data, err := Pembagian(100, 0)
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	fmt.Println("Hasil", data)

}

func Pembagian(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagi tidak boleh nol")
	} else {
		return nilai / pembagi, nil
	}
}
