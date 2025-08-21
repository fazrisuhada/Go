package main

import "fmt"

type ValidationError struct {
	Message string
}

func (v *ValidationError) Error() string {
	return v.Message
}

type NotFoundError struct {
	Message string
}

func (n *NotFoundError) Error() string {
	return n.Message
}

func SaveData(uuid string, data any) error {
	if uuid == "" {
		return &ValidationError{Message: "Validation Error"}
	}

	if uuid != "1-2-3" {
		return &NotFoundError{Message: "Not Found Error"}
	}

	return nil
}

func main() {
	err := SaveData("", nil)

	if err != nil {
		switch finalError := err.(type) {
		case *ValidationError:
			fmt.Println("error: ", finalError.Error())
		case *NotFoundError:
			fmt.Println("error:", finalError.Error())
		default:
			fmt.Println("error", finalError.Error())
		}
	} else {
		fmt.Println("Sukses")
	}
}
