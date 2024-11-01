package main

import "fmt"

type MathTest interface {
	Add() int
}

type Num struct {
	n int
}

func (num Num) Add() int {
	return num.n + num.n
}

type MyInt int

func (i MyInt) Add() int {
	return int(i + 3)
}

func main() {
	var m MathTest = Num{1}
	fmt.Println(m.Add())

	var mm MathTest = MyInt(2)
	fmt.Println(mm.Add())
}
