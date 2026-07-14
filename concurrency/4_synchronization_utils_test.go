package concurrency

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestOnce(t *testing.T) {
	var once sync.Once
	var group sync.WaitGroup

	for range 100 {
		group.Add(1)
		go func() {
			defer group.Done()
			once.Do(func() {
				fmt.Println("Loading config...")
				time.Sleep(2 * time.Second)
			})
		}()
	}

	group.Wait()
}

func TestPool(t *testing.T) {
	var pool sync.Pool

	pool.Put("Aaron")
	pool.Put("Evan")
	pool.Put("Juli")

	for range 10 {
		go func() {
			data := pool.Get()
			fmt.Println(data) // some prints will be nil, cause no one is putting anything into the pool after use
			time.Sleep(time.Second)
			pool.Put(data)
		}()
	}

	time.Sleep(2 * time.Second)
	fmt.Println("Done")
}

func TestPoolNew(t *testing.T) {
	pool := sync.Pool{
		New: func() any {
			return "New"
		},
	}

	pool.Put("Aaron")
	pool.Put("Evan")
	pool.Put("Juli")

	for range 10 {
		go func() {
			data := pool.Get()
			fmt.Println(data) // no nil prints cause New is defined, so get() will return "New" if there are no data in the pool
			time.Sleep(time.Second)
			pool.Put(data)
		}()
	}

	time.Sleep(2 * time.Second)
	fmt.Println("Done")
}

type JSONParser struct {
	Data []byte
}

var parserPool = sync.Pool{
	New: func() any {
		fmt.Println("Create parser")
		return &JSONParser{}
	},
}

func Parse(data []byte) {
	parser := parserPool.Get().(*JSONParser)

	parser.Data = data
	fmt.Println(string(parser.Data))
	parser.Data = nil
	parserPool.Put(parser)

}

func TestAnotherPool(t *testing.T) {
	var wg sync.WaitGroup

	for i := range 10 {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()

			Parse(fmt.Appendf(nil, "User %d", i))
		}(i)
	}

	wg.Wait()
}

func TestSyncCond(t *testing.T) {
	cond := sync.NewCond(&sync.Mutex{})
	group := &sync.WaitGroup{}

	for i := 1; i <= 10; i++ {
		group.Add(1)
		// wait condition
		go func() {
			cond.L.Lock()
			cond.Wait()
			fmt.Println("Done", i)
			cond.L.Unlock()
			group.Done()
		}()
	}

	go func() {
		for range 10 {
			time.Sleep(10 * time.Millisecond)
			cond.Signal() // without this will blocked forever or deadlock
		}
	}()

	group.Wait()
}

func TestSyncCondBroadcast(t *testing.T) {
	cond := sync.NewCond(&sync.Mutex{})
	group := &sync.WaitGroup{}

	for i := 1; i <= 10; i++ {
		group.Add(1)
		// wait condition
		go func() {
			cond.L.Lock()
			cond.Wait()
			fmt.Println("Done", i)
			cond.L.Unlock()
			group.Done()
		}()
	}

	go func() {
		time.Sleep(10 * time.Millisecond)
		cond.Broadcast()
	}()

	group.Wait()
}

func TestAtomicInt64(t *testing.T) {
	var wg sync.WaitGroup
	var counter int64 = 0
	for range 100 {
		wg.Add(1)

		go func() {
			for j := 1; j <= 33; j++ {
				atomic.AddInt64(&counter, 1)
			}

			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("Counter Load", atomic.LoadInt64(&counter))
	fmt.Println("Counter", counter)
}

func TestAtomicCompareAndSwap(t *testing.T) {
	var running atomic.Int32
	var wg sync.WaitGroup

	for range 10 {
		wg.Add(1)

		go func() {
			defer wg.Done()

			if running.CompareAndSwap(0, 1) {
				fmt.Println("Server started")

				time.Sleep(time.Second)

				// server stopped
				running.Store(0)

				fmt.Println("Server stopped")
			} else {
				fmt.Println("Already running")
			}
		}()
	}

	wg.Wait()

	fmt.Println("Running:", running.Load())
}
