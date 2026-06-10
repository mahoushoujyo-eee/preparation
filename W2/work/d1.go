package work

import (
	"fmt"
	"sync/atomic"
)

func PrintWithGoroutine() {
	for i := 0; i < 1000; i++ {
		go func(i int) {
			fmt.Println(i)
		}(i)
	}
}

func PrintWithGoroutineInOrder() {
	var n atomic.Int64
	n.Store(0)
	for i := 0; i < 1000; i++ {
		go func() {
			fmt.Println(n.Add(1))
		}()
	}
}
