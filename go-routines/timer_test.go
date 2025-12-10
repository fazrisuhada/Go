package belajar_golang_goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestAfterFunction(t *testing.T) {
	group := sync.WaitGroup{}
	group.Add(1)

	time.AfterFunc(5*time.Second, func() {
		fmt.Println(time.Now())
		group.Done()
	})

	fmt.Println(time.Now())
	group.Wait()
}
