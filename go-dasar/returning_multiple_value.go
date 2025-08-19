package main

import "fmt"

func getFullName() (string, string) {
	return "Fazri", "Suhada"
}

func getDrinkAndFood() (string, string) {
	return "Randang", "Es Teh"
}
func main() {
	firstName, lastName := getFullName()
	fmt.Println(firstName, lastName)

	// menghiraukan return value
	_, drink := getDrinkAndFood()
	fmt.Println(drink)
}
