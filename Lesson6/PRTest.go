package main

import "fmt"

func main() {
	defer func() {
		if err := recover(); err != nil {
			println(err.(string)) // 将 interface{} 转型为具体类型。
		}
	}()
	G()
	fmt.Println("main")
}

func G() {
	//defer func() {
	//	if err := recover(); err != nil {
	//		println(err.(string)) // 将 interface{} 转型为具体类型。
	//	}
	//}()
	F()
	fmt.Println("G")
}

func F() {
	defer func() {
		if err := recover(); err != nil {
			println(err.(string)) // 将 interface{} 转型为具体类型。
		}
	}()

	panic("panic error!")
	fmt.Println("F")
}
