package main

import (
	"fmt"
	"slices"
)

func main() {
	orang := []string{"Adi", "Budi", "Carlie", "Dodi", "Eka"}
	nilai := []int{70, 90, 40, 100}

	fmt.Println(slices.Min(orang))
	fmt.Println(slices.Min(nilai))
	fmt.Println(slices.Max(orang))
	fmt.Println(slices.Max(nilai))
	fmt.Println(slices.Contains(orang, "Budi"))
	fmt.Println(slices.Index(orang, "Dodi"))
}
