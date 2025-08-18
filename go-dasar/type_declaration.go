package main

import "fmt"

func main() {
	type noHp string

	var noHpCS noHp = "081287637193"

	var contohNoHp noHp = noHp("082264738299")

	fmt.Println(noHpCS)
	fmt.Println(contohNoHp)
}
