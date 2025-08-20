package main

import "fmt"

func sayGoodBye(name string) string {
	return "Good bye " + name
}

func main() {
	goodbye := sayGoodBye
	fmt.Println(goodbye("fazri"))
}
