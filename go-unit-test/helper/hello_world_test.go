package helper

import (
	"fmt"
	"runtime"
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

func TestSkip(t *testing.T) {
	if runtime.GOOS == "windows" {
		fmt.Println("Can't run on window")
	}

	result := HelloWorld("Suhada")
	assert.Equal(t, "Hello Suhada", result, "result must be 'Hello Suhada'")
	fmt.Println("Test Skip Done.")
}

func TestSubTest(t *testing.T) {
	t.Run("Fazri", func(t *testing.T) {
		result := HelloWorld("Fazri")
		require.Equal(t, "Hello Fazri", result, "result must be 'Hello Fazri'")
		fmt.Println("Test Skip Done.")
	})

	t.Run("Suhada", func(t *testing.T) {
		result := HelloWorld("Suhada")
		require.Equal(t, "Hello Suhada", result, "result must be 'Hello Suhada'")
		fmt.Println("Test Skip Done.")
	})
}

func TestMain(m *testing.M) {
	fmt.Println("Sebelum unit test")

	m.Run()

	fmt.Println("Sesudah unit test")
}
