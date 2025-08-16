package main

import "fmt"

func main() {
	// cara 1
	const bahasa = "Go-lang"
	fmt.Println("Nama Bahasa Pemrograman", bahasa)

	// cara 2 multiple
	const (
		author  = "Google"
		version = "1.25.0"
	)
	fmt.Println(author, version)
}
