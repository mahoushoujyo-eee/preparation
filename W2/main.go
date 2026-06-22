package main

import (
	"fmt"
	"log"
	"path"
	"reflect"
	"runtime"
	"sync"
	"w2/weekend"
	"w2/work"
)

func main() {
	runWeekendWork()
}

func runWeekendWork() {
	weekend.RunHealthCollector()
}

func runWork45() {
	go work.RunServer()
	work.RunClient()
}

func runOnceDemo() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println("go routine")
			work.OnceDemo()
		}()
	}

	wg.Wait()
	fmt.Println("done")
}

func runWork123() {
	runWithLog(work.PrintWithGoroutine)

	runWithLog(work.ChannelDemo)
	runWithLog(work.SelectDemo)
	runWithLog(work.RunDemo)

	runWithLog(work.MapDemo)
	runWithLog(work.MutexDemo)
	runWithLog(work.WaitGroupDemo)
	runWithLog(work.OnceDemo)
	runWithLog(work.CondDemo)
	runWithLog(work.PoolDemo)
	runWithLog(work.AtomicDemo)
}

func runWithLog(workFunc func()) {
	// 找了个可以运行时获取函数名的方法
	name := path.Base(
		runtime.FuncForPC(reflect.ValueOf(workFunc).Pointer()).Name(),
	)
	log.Printf("---- %s start ----", name)
	workFunc()
	log.Printf("---- %s end ----", name)
}
