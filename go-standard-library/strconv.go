package main

import (
	"fmt"
	"strconv"
)

func main() {
	result, err := strconv.ParseBool("true")
	if err != nil {
		fmt.Println("Error", err.Error())
	} else {
		fmt.Println(result)
	}

	resultInt, err := strconv.Atoi("100") //mengubah string ke integer
	if err != nil {
		fmt.Println("Error", err.Error())
	} else {
		fmt.Println(resultInt)
		fmt.Printf("Tipe data: %T", resultInt)
	}

	var stringInt = strconv.Itoa(1000) //mengubah integer ke string
	fmt.Println(stringInt)
	fmt.Printf("Tipe data: %T", stringInt)
}
