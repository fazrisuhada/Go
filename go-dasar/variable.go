package main

import "fmt"

func main() {
	// cara 1 menggunakan var
	var buah string
	buah = "Apel"

	fmt.Println(buah)

	// cara 2 menggunakan :=
	jenkel := "Pria"
	fmt.Println(jenkel)

	// cara 3 multiple variable
	var (
		makanan = "Rendang"
		minuman = "Air Mineral"
	)
	fmt.Println(makanan, minuman)
}
