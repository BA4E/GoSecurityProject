package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

// 股票价格模拟器：用atomic存储和读取股票价格
type Stock struct {
	price atomic.Value // 用于线程安全存储价格
}

// 初始化股票价格
func NewStock(initialPrice float64) *Stock {
	stock := &Stock{}
	stock.price.Store(initialPrice) // 设置初始价格
	return stock
}

// 获取当前股票价格
func (s *Stock) GetPrice() float64 {
	return s.price.Load().(float64)
}

// 更新股票价格
func (s *Stock) UpdatePrice(newPrice float64) {
	s.price.Store(newPrice)
}

// 模拟API调用，随机生成价格变动
func fetchStockPrice() float64 {
	return 100.0 + rand.Float64()*10.0 // 生成100-110之间的随机价格
}

func main() {
	rand.Seed(time.Now().UnixNano()) // 设置随机数种子
	stock := NewStock(100.0)         // 初始化股票

	// 创建两个通道：一个用于接收价格更新，另一个用于停止定时器
	priceChan := make(chan float64)
	stopChan := make(chan bool)

	// 定时获取价格的协程
	go func() {
		ticker := time.NewTicker(2 * time.Second) // 每2秒获取一次价格
		defer ticker.Stop()                       // 确保定时器停止

		for {
			select {
			case <-ticker.C: // 定时获取新的价格
				newPrice := fetchStockPrice()
				priceChan <- newPrice // 发送到通道
			case <-stopChan: // 收到停止信号时退出
				fmt.Println("Ticker stopped.")
				return
			}
		}
	}()

	// 处理价格更新的协程
	go func() {
		for price := range priceChan {
			stock.UpdatePrice(price) // 更新股票价格
			fmt.Printf("Updated Stock Price: %.2f\n", stock.GetPrice())
		}
	}()

	// 模拟程序运行10秒后停止
	time.Sleep(10 * time.Second)
	stopChan <- true // 发送停止信号

	// 确保主程序等待所有协程结束
	time.Sleep(1 * time.Second)
	fmt.Println("Final Stock Price:", stock.GetPrice())
}
