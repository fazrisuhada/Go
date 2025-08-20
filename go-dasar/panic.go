package main

import "fmt"

func endApp() {
	fmt.Println("End App")

	mesagge := recover()
	fmt.Println("Terjadi error", mesagge)
}

func runApp(error bool) {
	defer endApp()

	if error {
		panic("Ups")
	}
}
func main() {
	runApp(true)
	fmt.Println("Fazri Suhada")
}
