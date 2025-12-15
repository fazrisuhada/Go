package belajar_golang_context

import (
	"context"
	"fmt"
	"testing"
)

func TestContex(t *testing.T) {
	background := context.Background()
	fmt.Println(background)

	todo := context.TODO()
	fmt.Println(todo)
}
