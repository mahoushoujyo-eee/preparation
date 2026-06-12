package work

import (
	"fmt"
	"log"
	"sync"
	"sync/atomic"
)

func mapDemo() {
	var syncMap sync.Map
	syncMap.Store("key1", 11)
	syncMap.Store("key2", 12)
	val, ok := syncMap.Load("key1")
	if ok {
		val, ok = val.(int)
		fmt.Println(val)
	}
}

// consider read write lock:sync.RWMutex
type Counter struct {
	val int
	mux sync.Mutex
}

func (c *Counter) Get() int {
	c.mux.Lock()
	res := c.val
	c.mux.Unlock()
	return res
}

func (c *Counter) AddAndGet(val1 int) int{
	c.mux.Lock()
	c.val = c.val + val1
	res := c.val
	c.mux.Unlock()
	return res
}

func (c *Counter) Store(val1 int) {
	c.mux.Lock()
	c.val = val1
	c.mux.Unlock()
}

type AtomicCounter struct {
	val atomic.Int64
}

func (a *AtomicCounter) Get() int64{
	return a.val.Load()
}

func (a *AtomicCounter) Inc() {
	a.val.Add(1)
}

func mutexDemo(){
	var oneMutex sync.Mutex
	oneMutex.Lock()
	if !oneMutex.TryLock() {
		// do something
	}
	oneMutex.Unlock()

	var oneRwMutex sync.RWMutex

	oneRwMutex.RLock()
	oneRwMutex.Lock()
	rLock := oneRwMutex.RLocker()
	rLock.Unlock()
	oneRwMutex.Unlock()
}

func WaitGroupDemo(){
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func (i int)  {
			fmt.Println("go with ", i)
			defer wg.Done()	
		}(i)
	}

	wg.Wait()
	fmt.Println("all goroutine finish")
}

var once  sync.Once
func OnceDemo(){
	once.Do(func() {
		fmt.Println("init once")
	})
}

func CondDemo(){
	var mut sync.Mutex
	var oneCond *sync.Cond = sync.NewCond(&mut)
	var wg sync.WaitGroup
	wg.Add(1)
	go func ()  {
		mut.Lock()
		oneCond.Wait()
		mut.Unlock()
		fmt.Println("do something...")
		wg.Done()
	}()

	fmt.Println("signal")
	oneCond.Signal()

	wg.Wait()
	fmt.Println("wg passthrough")
}

type Buffer struct {
    Data []byte
    Pos  int
}

var bufferPool *sync.Pool = &sync.Pool{
	New: func() any {
		return &Buffer{
			Data: make([]byte, 1024),
			Pos: 0,
		}
	},
}

func PoolDemo(){
	tmpBuffer := bufferPool.Get()
	buffer, ok := tmpBuffer.(*Buffer)
	if !ok {
		log.Fatalln("this value is not a expected type")
	}
	fmt.Println("buffer pos", buffer.Pos)
}

func AtomicDemo(){
	var val atomic.Int64
	val.Store(2)
	val.Add(1)
	res := val.Load()
	fmt.Println(res)
	val.CompareAndSwap(3, 5)
}