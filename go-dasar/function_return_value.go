package main

import "fmt"

func hello(name string) string {
	return "Hello, " + name
}

func main() {
	result := hello("Fazri")
	fmt.Println(result)
}
