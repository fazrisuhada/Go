package main

import (
	"errors"
	"fmt"
)

var (
	ValidationError = errors.New("validation error.")
	IdNotFoundError = errors.New("id not found error.")
)

func GetById(uuid string) error {

	if uuid == "" {
		return ValidationError
	}

	if uuid != "abc12" {
		return IdNotFoundError
	}

	return nil
}

func main() {

	err := GetById("abc123")

	if err != nil {
		if errors.Is(err, ValidationError) {
			fmt.Println("validation error")
		} else if errors.Is(err, IdNotFoundError) {
			fmt.Println("id not found error")
		} else {
			fmt.Println("unknow error")
		}
	}
}
