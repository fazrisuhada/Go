package main

import "fmt"

func getCompleteName() (name, address, phone string) {
	name = "Fazri Suhada"
	address = "Bandung"
	phone = "08129039490"

	return name, address, phone
}

func main() {
	a, b, c := getCompleteName()
	fmt.Println(a, b, c)
}
