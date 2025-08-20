package main

import "fmt"

func sumAll(numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}

	return total
}

func main() {
	fmt.Println(sumAll(5, 5, 5))
	fmt.Println(sumAll(5, 5, 5, 5))
	fmt.Println(sumAll(5, 5, 5, 5, 5))

	// kode function slice parameter
	total2 := sumAll(10, 10, 10, 10)
	fmt.Println(total2)

	numbers2 := []int{10, 10, 10, 10}
	total2 = sumAll(numbers2...)
	fmt.Println()
}
