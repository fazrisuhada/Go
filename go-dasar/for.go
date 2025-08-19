package main

import "fmt"

func main() {

	// kode program for
	counter := 1
	for counter <= 10 {
		fmt.Println("Angke ke-", counter)
		counter++
	}

	// for dengan statment
	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Angka ke", counter)
	}

	// for range
	fruits := []string{"Apel", "Manga", "Alpukat", "Jeruk"}
	for index, value := range fruits {
		fmt.Println("index", index, "=", value)
	}

	for _, value := range fruits {
		fmt.Println(value)
	}
}
