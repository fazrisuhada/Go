package main

import "fmt"

func endApp() {
	fmt.Println("End App")
}

func runApp(error bool) {
	defer endApp()

	if error {
		panic("Ups")
	}
}
func main() {
	runApp(true)
}
