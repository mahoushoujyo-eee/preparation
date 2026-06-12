package weekend

import (
	"fmt"
	"sync"
	"time"
)

func ConcurrentCall() {
	semaphore := make(chan struct{}, 5)
	res := make(chan int, 6)
	var wg sync.WaitGroup

	for i := range 10 {
		wg.Add(1)
		go func(i int) {
			semaphore<- struct{}{}
			fmt.Println("call service", i)
			time.Sleep(2*time.Second)
			fmt.Println("finish service", i)
			<-semaphore
			res <- i
			wg.Done() 
		}(i)
	}

	go func ()  {
		wg.Wait()
		close(semaphore)
		close(res)
	}()

	for oneRes := range res {
		fmt.Println("result of service", oneRes)
	}

	fmt.Println("all service calling finish")
}