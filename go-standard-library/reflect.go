package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name    string `required:"true" max:"10"`
	Address string `required:"true" max:"10"`
}

func readField(value any) {
	var valueType reflect.Type = reflect.TypeOf(value)
	fmt.Println("Type name : ", valueType.Name())
	for i := 0; i < valueType.NumField(); i++ {
		var structField = valueType.Field(i)
		fmt.Println(structField.Name, "with type", structField.Type)
		fmt.Println(structField.Tag.Get("required"))
		fmt.Println(structField.Tag.Get("max"))
	}
}

func isValid(value any) (result bool) {
	result = true
	t := reflect.TypeOf(value)
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if f.Tag.Get("required") == "true" {
			data := reflect.ValueOf(value).Field(i).Interface()
			result = data != ""
			if result == false {
				return result
			}
		}
	}
	return result
}

func main() {
	readField(Person{"Fazri", "Bandung"})
	Person := Person{
		Name:    "Fazri",
		Address: "",
	}
	fmt.Println(isValid(Person))
}
