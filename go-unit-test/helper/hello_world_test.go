package helper

import "testing"

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Fazri")

	if result != "Hello Fazri" {
		panic("Result is not 'Hello Fazri'")
	}
}
