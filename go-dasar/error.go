package main

import (
	"errors"
	"fmt"
)

func Pembagian(number int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagi 0")
	} else {
		return number / pembagi, nil
	}
}

func main() {
	hasil, err := Pembagian(10, 0)
	if err == nil {
		fmt.Println("Hasil", hasil)
	} else {
		fmt.Println("Error", err.Error())
	}
}
