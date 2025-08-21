package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	fazri := Man{"Fazri"}
	fazri.Married()
	fmt.Println(fazri.Name)
}
