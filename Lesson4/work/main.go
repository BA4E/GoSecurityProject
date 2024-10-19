package main

import (
	"flag"
	"fmt"
	"net"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
)

var (
	lock sync.Mutex
)

func main() {
	hostname := flag.String("host", "127.0.0.1", "target hostname")
	ports := flag.String("ports", "1-65535", "target ports")
	threads := flag.Int("threads", 5, "number of threads")
	flag.Parse()
	split := strings.Split(*ports, "-")
	start, _ := strconv.Atoi(split[0])
	end, _ := strconv.Atoi(split[1])

	var open []int
	ch := make(chan int, *threads)
	fmt.Println("开始扫描")
	for i := start; i <= end; i++ {
		ch <- 1
		go func(hostname string, port int) {
			address := fmt.Sprintf("%s:%d", hostname, port)
			conn, _ := net.DialTimeout("tcp", address, 2*time.Second)
			if conn != nil {
				lock.Lock()
				open = append(open, port)
				lock.Unlock()
			}
			<-ch
		}(*hostname, i)
	}
	close(ch)
	sort.Ints(open)
	fmt.Printf("Port\tStatus\n")
	for _, v := range open {
		if v < 1000 {
			fmt.Printf("%d  \tOPEN\n", v)
		} else {
			fmt.Printf("%d\tOPEN\n", v)
		}
	}
	fmt.Println("扫描结束")
}
