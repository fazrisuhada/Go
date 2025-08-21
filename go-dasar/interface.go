package main

import "fmt"

type HasName interface {
	GetName() string
}

func sayHai(hasName HasName) {
	fmt.Println("Hello", hasName.GetName())
}

// Kode interface 1
type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

// kode interface 2
type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

func main() {
	person := Person{Name: "Fazri"}
	sayHai(person)

	animal := Animal{Name: "Kucing"}
	sayHai(animal)
}
