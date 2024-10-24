package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup

func Basic1Test() {
	var x int32 = 10
	atomic.AddInt32(&x, 1)
	fmt.Println(x)

	atomic.StoreInt32(&x, 20)
	fmt.Println(x)

	old := atomic.SwapInt32(&x, 30)
	fmt.Println(x, old)

	atomic.CompareAndSwapInt32(&x, 30, 40)
	fmt.Println(x)
	atomic.CompareAndSwapInt32(&x, 3, 50)
	fmt.Println(x)
}

func Add(x *int32) {
	defer wg.Done()
	atomic.AddInt32(x, 1)
}
func Sub(x *int32) {
	defer wg.Done()
	atomic.AddInt32(x, -1)
}

func main() {
	var n int32 = 100
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go Add(&n)
		wg.Add(1)
		go Sub(&n)
	}
	wg.Wait()
	fmt.Println(n)

}
