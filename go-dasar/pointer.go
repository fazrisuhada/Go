package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {

	var address1 = Address{"Bukittinggi", "Sumatra Barat", "Indonesia"}

	var address2 = &address1 //pointer gunakan &

	address2.City = "Payakumbuh"
	fmt.Println(address1)
	fmt.Println(address2)

	// address2 = &Address{"Jakarta","DKI Jakarta","Indonesia"}
	*address2 = Address{"Jakarta", "DKI Jakarta", "Indonesia"}

	fmt.Println(address1)
	fmt.Println(address2)

}
