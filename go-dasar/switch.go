package main

import (
	"fmt"
)

func main() {
	name := "Fazri"

	switch name {
	case "Fazri":
		fmt.Println("Hai, Fazri")
	case "Budi":
		fmt.Println("Hai, Budi")
	default:
		fmt.Println("Nama Kamu Siapa?")
	}

	// switch short statement
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama Kamu terlalu panjang.")
	case false:
		fmt.Println("Nama Kamu sudah benar.")
	}

	// switch tanpa kondisi
	panjang := len(name)
	switch {
	case panjang > 10:
		fmt.Println("Nama terlalu panjang.")
	case panjang > 5:
		fmt.Println("Nama lumayan panjang.")
	default:
		fmt.Println("Nama sudah benar.")
	}
}
