package main

import "fmt"

type Address struct {
	City, Province, Country string
}

type Person struct {
	Name string
	Age  int
}

func main() {

	var address1 = Address{"Bukittinggi", "Sumatra Barat", "Indonesia"}

	var address2 = &address1 //pointer gunakan & atau new

	address2.City = "Payakumbuh"
	fmt.Println(address1)
	fmt.Println(address2)

	// address2 = &Address{"Jakarta","DKI Jakarta","Indonesia"}
	*address2 = Address{"Jakarta", "DKI Jakarta", "Indonesia"}

	fmt.Println(address1)
	fmt.Println(address2)

	person1 := Person{"Fazri", 31}
	person2 := new(Person)
	person2.Name = "Suhada"
	person2.Age = 10
	fmt.Println(person1)
	fmt.Println(person2)

}
