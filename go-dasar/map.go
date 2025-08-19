package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "Fazri",
		"address": "Bukittinggi",
		"gender":  "Man",
	}

	fmt.Println(person)

	delete(person, "gender")
	fmt.Println(person)
}
