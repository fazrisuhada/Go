package helper

import (
	"fmt"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Fazri")

	if result != "Hello Fazri" {
		t.Error("Result must be 'Hello Fazri'")
		// t.Fatal("Result must be 'Hello Fazri'")
	}

	fmt.Println("Test Hello World Done")
}
