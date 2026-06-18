package work

import (
	"fmt"
	"log"
	"sync"
	"sync/atomic"
	"time"
)

func MapDemo() {
	var syncMap sync.Map
	syncMap.Store("key1", 11)
	syncMap.Store("key2", 12)

	raw, ok := syncMap.Load("key1")
	if !ok {
		log.Printf("key1 not found")
		return
	}

	v, ok := raw.(int)
	if !ok {
		log.Printf("key1 is not int, got %T", raw)
		return
	}

	log.Println(v)
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

func (c *Counter) AddAndGet(val1 int) int {
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

func (a *AtomicCounter) Get() int64 {
	return a.val.Load()
}

func (a *AtomicCounter) Inc() {
	a.val.Add(1)
}

func MutexDemo() {
	// TryLock：未持锁时尝试获取，成功就用，失败就跳过；不会阻塞。
	var oneMutex sync.Mutex
	if oneMutex.TryLock() {
		defer oneMutex.Unlock()
		fmt.Println("trylock acquired")
	} else {
		fmt.Println("trylock skipped")
	}

	var oneRwMutex sync.RWMutex

	oneRwMutex.RLock()
	// 只读临界区
	oneRwMutex.RUnlock()

	oneRwMutex.Lock()
	// 写临界区
	oneRwMutex.Unlock()

	rLock := oneRwMutex.RLocker()
	rLock.Lock()   // 内部是 RLock
	rLock.Unlock() // 内部是 RUnlock
}

func WaitGroupDemo() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			fmt.Println("go with ", i)
			defer wg.Done()
		}(i)
	}

	wg.Wait()
	fmt.Println("all goroutine finish")
}

var once sync.Once

func OnceDemo() {
	once.Do(func() {
		fmt.Println("init once")
	})
}

// ---- Cond Demo ----

// 之前的例子不仅有bug产生死锁，而且也没有真正用到Cond，只是一种强硬演示
type BoundedQueue struct {
	mu       sync.Mutex
	notFull  *sync.Cond // 队列不满生产者可以Push
	notEmpty *sync.Cond // 队列非空消费者可以Pop
	items    []int
	cap      int
}

func NewBoundedQueue(cap int) *BoundedQueue {
	q := &BoundedQueue{cap: cap}
	q.notFull = sync.NewCond(&q.mu)
	q.notEmpty = sync.NewCond(&q.mu)
	return q
}

func (q *BoundedQueue) Push(v int) {
	q.mu.Lock()
	defer q.mu.Unlock()

	// 存在假唤醒，比如BroadCast但是没有抢到锁，就要继续wait，所以要for而非if
	for len(q.items) == q.cap {
		q.notFull.Wait()
	}

	q.items = append(q.items, v)

	// Push一次通知一个就行了，所以是Signal
	q.notEmpty.Signal()
}

func (q *BoundedQueue) Pop() int {
	q.mu.Lock()
	defer q.mu.Unlock()

	for len(q.items) == 0 {
		q.notEmpty.Wait()
	}

	v := q.items[0]
	q.items = q.items[1:]
	q.notFull.Signal()
	return v
}

func (q *BoundedQueue) Len() int {
	q.mu.Lock()
	defer q.mu.Unlock()
	return len(q.items)
}

func CondDemo() {
	q := NewBoundedQueue(3)
	var wg sync.WaitGroup

	for id := range 2 {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := range 5 {
				v := id*100 + j
				q.Push(v)
				log.Printf("[P%d] pushed %3d   (len=%d)\n", id, v, q.Len())
			}
		}(id)
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			time.Sleep(80 * time.Millisecond)
			v := q.Pop()
			log.Printf("[C ] popped  %3d   (len=%d)\n", v, q.Len())
		}
	}()

	wg.Wait()
	log.Printf("done")
}

// ---- Pool Demo ----

type Buffer struct {
	Data []byte
	Pos  int
}

var bufferPool *sync.Pool = &sync.Pool{
	New: func() any {
		return &Buffer{
			Data: make([]byte, 1024),
			Pos:  0,
		}
	},
}

func PoolDemo() {
	buf := bufferPool.Get().(*Buffer)
	defer bufferPool.Put(buf)

	buf.Pos = 0

	fmt.Println("buffer pos", buf.Pos)
}

func AtomicDemo() {
	var val atomic.Int64
	val.Store(2)
	val.Add(1)
	res := val.Load()
	fmt.Println("after store(2)+add(1), val =", res) // 3

	ok := val.CompareAndSwap(3, 5)
	fmt.Println("cas 3->5:", ok, "now =", val.Load()) // true, 5

	ok = val.CompareAndSwap(3, 7)
	fmt.Println("cas 3->7:", ok, "now =", val.Load()) // false, 5
}
