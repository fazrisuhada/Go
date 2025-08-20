package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func main() {
	var armand Customer
	armand.Name = "Armand ary"
	armand.Address = "Bukittinggi"
	armand.Age = 30

	fmt.Println(armand)

	// kode struct literal
	yuda := Customer{
		Name:    "Yuda",
		Address: "Ternate",
		Age:     24,
	}
	fmt.Println(yuda)

	deni := Customer{"Deni", "Banuhampu", 30}
	fmt.Println(deni)
}
