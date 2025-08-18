package main

import "fmt"

func main() {

	var (
		a = 2
		b = 4
		c = 3
		d = 6
		e = 5
		f = 10
	)

	var hasil = a*b/c + d - e + f
	fmt.Println(hasil)

	// augmented assignment
	a += 1 // 2 + 1
	b -= 1 // 4 - 1
	c *= 1 // 3 * 1
	d /= 1 // 6 / 1
	fmt.Println(a, b, c, d)

	// unary operator
	e++ // 5 + 1
	fmt.Println(e)
	f-- //10-1
	fmt.Println(f)
}
