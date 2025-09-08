package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now.Local())

	utc := time.Date(2009, time.September, 10, 23, 0, 0, 0, time.UTC)
	fmt.Println(utc.Local())

	formatter := "2006-01-02 15:04:05"

	value := "2025-09-09 05:50:00"
	valueTime, err := time.Parse(formatter, value)

	if err != nil {
		fmt.Println("Error", err.Error())
	} else {
		fmt.Println(valueTime)
	}
}
