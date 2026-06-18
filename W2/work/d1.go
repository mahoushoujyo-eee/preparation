package work

import (
	"fmt"
	"sync"
)

func PrintWithGoroutine() {
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Print(i, "  ")
		}(i)
	}
	wg.Wait()
}

// func PrintWithGoroutineInOrder() {
// 	var n atomic.Int64
// 	n.Store(0)
// 	for i := 0; i < 1000; i++ {
// 		go func() {
// 			fmt.Println(n.Add(1))
// 		}()
// 	}
// }
