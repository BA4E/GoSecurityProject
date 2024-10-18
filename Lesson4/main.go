package main

import (
	"fmt"
	"strconv"
	"sync"
)

var wg sync.WaitGroup

func show(msg string) {
	defer wg.Done() //show运行完毕后协程池数量-1
	fmt.Println(msg)
}

func main() {
	for i := 0; i < 5; i++ {
		wg.Add(1) //协程池数量+1
		go show("This is " + strconv.Itoa(i))
	}
	wg.Wait() //等待协程池数目为0后再运行下面的代码
	fmt.Println("All finished")
	fmt.Println("main exit")

}
