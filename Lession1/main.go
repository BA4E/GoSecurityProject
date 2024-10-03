package main

import "fmt"

func test(s1 ...string) {

}

func main() {
	//匿名变量
	s1, _ := "s1", "s2"
	fmt.Println(s1)

	//ctrl + /单行注释

	//println输出换行

	arr := []int{1, 2, 3, 4, 5}
	println(arr[0])
	fmt.Println(arr)

	//自增没有++i

	//通道

	//shift + (或者"     快速括号引号

	//切片大小可变，数组不行
	arr = append(arr, 777)
	fmt.Println(arr)

	//省略号...表多个参数
	test()

}
