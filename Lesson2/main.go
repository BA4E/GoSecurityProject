package main

import (
	"fmt"
	"strings"
)

func TestConst() {
	const (
		Pi       = 3.14
		Language = "GO"
		IsTrue   = true
	)

	const (
		Zero = iota
		One
		Two
	)

	const (
		mutexLocked = 1 << iota
		mutexWoken
		mutexStarving
		mutexWaiterShift = iota
	)

	fmt.Println("Pi:", Pi)
	fmt.Println("Language:", Language)
	fmt.Println("IsTrue:", IsTrue)
	fmt.Println("Zero:", Zero)
	fmt.Println("One:", One)
	fmt.Println("Two:", Two)

	fmt.Println("mutexLocked:", mutexLocked)
	fmt.Println("mutexWoken:", mutexWoken)
	fmt.Println("mutexStarving:", mutexStarving)
	fmt.Println("mutexWaiterShift:", mutexWaiterShift)
}

func TestArr() {
	var arr = [5]int{1, 2, 3, 4, 5}
	arr2 := [5]int{1, 2, 3, 4, 5}
	arr3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(arr[1], arr2[2], arr3[3])
	arr[2] = 6
	fmt.Println(arr[2])
	fmt.Println(len(arr))
	b := arr
	b[0] = 10
	fmt.Println(b)
	for i := 0; i < len(arr); i++ {
		fmt.Print(arr[i])
	}

	for k, v := range arr {
		fmt.Printf("key: %d,val: %d", k, v)
	}
}

func TestSlice() {
	var num = []int{1, 2, 3, 4, 5}
	num1 := make([]int, 5)
	num2 := make([]int, 5, 10)
	fmt.Println(len(num1))
	fmt.Println(len(num2))

	fmt.Println(num1[0])
	num1[0] = 10
	fmt.Println(num1[0])

	fmt.Println(cap(num1), cap(num2))

	sub := num[1:4]
	fmt.Println(sub)

	num = append(num, 6, 7, 8)

	fmt.Println(num)
	num3 := num
	num3[0] = 10
	fmt.Println(num3)
}

func TestPtr() {
	var ptr *int
	var a int = 10
	ptr = &a
	fmt.Println(*ptr)
	Modifyvalue(ptr)
	fmt.Println(*ptr)

	type Person struct {
		name string
		age  int
	}

	person := Person{"a", 1}
	p := &person
	p.age = 2
	fmt.Println(p.name, p.age, p)

	var pp *int
	fmt.Println(pp)
	if pp != nil {
		fmt.Println(*pp)
	}

	slice := [5]int{1, 23, 4}
	ModifySlice(slice)
	fmt.Println(slice)
}

func Modifyvalue(val *int) {
	*val = 20
}

func ModifySlice(s [5]int) {
	s[0] = 100
	fmt.Println("Mo: ", s)
}

func TestTypeChange() {
	var i int = 43
	var flo64 = float64(i)
	var ui = uint(i)
	fmt.Println(i, flo64, ui)

	flo64 = 3.14
	i = int(flo64)
	fmt.Println(flo64, i)

	var s string = "Hello world!"
	var b []byte = []byte(s)
	var s2 string = string(b)
	fmt.Println(s, b, s2)

	var e interface{} = "Hello"
	d := e.(string)
	fmt.Println(e, d)

}

func Teststring() {
	//多行字符串
	s := `a
a
a
a

a
`
	fmt.Println(s)
	contains := strings.Contains("Hellow world", "world")
	fmt.Println(contains)
}

func main() {
	Teststring()
}
