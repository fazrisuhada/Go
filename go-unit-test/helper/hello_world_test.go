package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Fazri")

	if result != "Hello Fazri" {
		t.Error("Result must be 'Hello Fazri'")
		// t.Fatal("Result must be 'Hello Fazri'")
	}

	fmt.Println("Test Hello World Done")
}

func TestHelloWorldAssertion(t *testing.T) {
	result := HelloWorld("Suhada")

	assert.Equal(t, "Hello Suhada", result, "result must be 'Hello Suhada'")
	fmt.Println("Test Hello World Assertion Done.")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Suhada")

	require.Equal(t, "Hello Suhada", result, "result must be 'Hello Suhada'")
	fmt.Println("Test Hello World Require Done.")
}
