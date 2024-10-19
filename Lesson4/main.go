package main

import (
	"fmt"
	"sync"
)

//var wg sync.WaitGroup

//func show(msg string) {
//	defer wg.Done() //show运行完毕后协程池数量-1
//	fmt.Println(msg)
//}

//func main() {
//for i := 0; i < 5; i++ {
//	wg.Add(1) //协程池数量+1
//	go show("This is " + strconv.Itoa(i))
//}
//wg.Wait() //等待协程池数目为0后再运行下面的代码
//fmt.Println("All finished")
//fmt.Println("main exit")

//}

//	func show(msg string) {
//		for i := 0; i < 10; i++ {
//			fmt.Println(msg)
//			if i > 5 {
//				runtime.Goexit()
//			}
//		}
//	}
//
//	func main() {
//		go show("Go")
//
//		runtime.Gosched()
//		fmt.Println("main end")
var (
	x    int = 100
	wg   sync.WaitGroup
	lock sync.Mutex
) //}

func Add() {
	defer wg.Done()
	lock.Lock()
	x = x + 1
	fmt.Printf("x++: %d\n", x)
	lock.Unlock()
}

func Sub() {
	defer wg.Done()
	lock.Lock()
	x = x - 1
	fmt.Printf("x--: %d\n", x)
	lock.Unlock()
}

func main() {
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go Add()
		wg.Add(1)
		go Sub()
	}
	wg.Wait()
	fmt.Println("main exit")
}
