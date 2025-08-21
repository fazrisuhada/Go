package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func changeCountryToIndonesia(address *Address) {
	address.Country = "Indonesia"
}

func main() {
	var address = Address{}
	changeCountryToIndonesia(&address)

	fmt.Println(address)
}
