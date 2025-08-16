package main

import "fmt"

func main() {
	// konfersi tipe data angka
	var (
		number16 int16 = 32767
		number32 int32 = int32(number16)
		number64 int64 = int64(number16)
	)

	fmt.Println(number16)
	fmt.Println(number32)
	fmt.Println(number64)

	// konfersi string
	fullName := "Fazri Suhada"
	e := fullName[0]
	eString := string(e)

	fmt.Println(fullName)
	fmt.Println(e)
	fmt.Println(eString)
}
