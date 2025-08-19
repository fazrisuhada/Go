package main

import "fmt"

func main() {
	name := "Randi"

	if name == "Fazri" {
		fmt.Println("Hai, Fazri")
	} else if name == "Budi" {
		fmt.Println("Hai, Budi")
	} else {
		fmt.Println("Nama Kamu Siapa ?")
	}

	if length := len(name); length > 5 {
		fmt.Println("Nama Kamu terlalu panjang.")
	} else {
		fmt.Println("Nama Kamu tidak terlalu panjang.")
	}
}
