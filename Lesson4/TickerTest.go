package main

import (
	"fmt"
	"time"
)

func BasicTest() {
	timer1 := time.NewTimer(2 * time.Second)
	t1 := time.Now()
	fmt.Printf("t1:%v\n", t1)
	t2 := <-timer1.C
	fmt.Printf("t2:%v\n", t2)
}

func StopTest() {
	// 1.获取ticker对象
	ticker := time.NewTicker(1 * time.Second)
	// 子协程
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			//<-ticker.C
			if i < 5 {
				fmt.Println(<-ticker.C, i)
			} else {
				fmt.Println(time.Now(), i)
			}
			if i == 5 {
				//停止
				ticker.Stop()
			}
		}
	}()
	wg.Wait()

}

func ResetTest() {
	timer5 := time.NewTimer(3 * time.Second)
	timer5.Reset(1 * time.Second)
	fmt.Println(time.Now())
	fmt.Println(<-timer5.C)

}
func main() {
	ResetTest()
}
