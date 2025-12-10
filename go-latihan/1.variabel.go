package main

import "fmt"

func main() {
	/*
		Buatlah program yang mendeklarasikan 3 variabel dengan tipe data berbeda (string, int, bool) menggunakan 3 cara yang berbeda (var, :=, dan var dengan tipe eksplisit). Tampilkan semua variabel tersebut.
	*/

	var name string
	name = "Fazri"

	isMale := true

	var age int = 31

	fmt.Println(name, age, isMale)
}
