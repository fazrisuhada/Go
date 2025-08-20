package main

import "fmt"

func logging() {
	fmt.Println("Selesai memanggil function")
}

func main() {
	defer logging()
	fmt.Println("Run App")
}
