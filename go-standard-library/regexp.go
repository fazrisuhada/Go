package main

import (
	"fmt"
	"regexp"
)

func main() {
	var regex = regexp.MustCompile(`a[a-z]i`)

	fmt.Println(regex.MatchString("adi"))
	fmt.Println(regex.MatchString("abu"))
	fmt.Println(regex.MatchString("aJi"))

	fmt.Println(regex.FindAllString("adi ali abi abu ani aji ARI", 10))
}
