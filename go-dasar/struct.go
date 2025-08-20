package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHello(name string) {
	fmt.Println("Hello", name, "my name is", customer.Name)
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

	armand.sayHello("Fazri")
	yuda.sayHello("Fazri")
	deni.sayHello("Fazri")
}
