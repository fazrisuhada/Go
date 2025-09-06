package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Fazri Suhada", "fazi"))
	fmt.Println(strings.Split("Salsabila Mustika Maharani Ababil", " "))
	fmt.Println(strings.ToUpper("Fazri Suhada"))
	fmt.Println(strings.ToLower("Fazri Suhada"))
	fmt.Println(strings.Trim("           Fazri Suhada       ", " "))
	fmt.Println(strings.ReplaceAll("Fazri Suhada", "Fazri", "Salsabila"))
}
