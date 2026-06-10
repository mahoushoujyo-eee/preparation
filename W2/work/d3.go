package work

import (
	"fmt"
	"sync"
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