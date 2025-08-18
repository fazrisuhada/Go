package main

import "fmt"

func main() {
	// contoh pertama
	var rak [3]string

	rak[0] = "laptop"
	rak[1] = "handphone"
	rak[2] = "Mouse"

	fmt.Println(rak)

	// contoh ke dua
	var values = [...]int{
		70, 80, 90, 100,
	}
	fmt.Println(len(values))
	fmt.Println(values[2])
}
