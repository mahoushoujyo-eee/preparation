package main

import (
	"fmt"
	"sync"
	"w2/weekend"
	"w2/work"
)

func main(){
	weekend.RunHealthCollector()
}


func main2(){
	go work.RunServer()
	work.RunClient()
}

func main1() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++{
		wg.Add(1)
		go func ()  {
			defer wg.Done()
			fmt.Println("go routine")	
			work.OnceDemo()
		}()
	}

	wg.Wait()
	fmt.Println("done")
}
