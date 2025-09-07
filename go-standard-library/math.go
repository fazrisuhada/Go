package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Ceil(1.50))  //bulatkan ke atas
	fmt.Println(math.Floor(1.50)) //bulatkan ke bawah
	fmt.Println(math.Round(1.50)) //bulatkan bilangan
	fmt.Println(math.Max(10, 11)) //mengambil nilai terbesar
	fmt.Println(math.Min(10, 11)) //mengambil nilai terkecil
}
