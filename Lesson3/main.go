package main

import "fmt"

type Info struct {
	Name string
	age  int
}

func main() {
	fmt.Printf("%v", 1234)

	fmt.Printf("%+v", struct {
		Name string
		age  int
	}{"A", 1})

	fmt.Printf("%#v", struct {
		Name string
		age  int
	}{"A", 1})

	fmt.Printf("%T", 1234)

	fmt.Printf("%%")

	fmt.Printf("%t", true)

	fmt.Printf("%b", 6)

	fmt.Printf("%c", 97)

	fmt.Printf("%q", 97)

	fmt.Printf("%q", "abcd")

	fmt.Printf("%d", 6)

	fmt.Printf("%o", 9)

	fmt.Printf("%O", 9)

	fmt.Printf("%x", 11)

	fmt.Printf("%X", 11)

	fmt.Printf("%U", 'a')

	s := "a"
	s2 := "1"
	fmt.Printf("%s %s", s, s2)

	i := 2
	p := &i
	fmt.Printf("%p", p)

	/*
		var info Info
		_, err := fmt.Scanf("%s,%d", &info.Name, &info.age)
		if err != nil {
			fmt.Println("err:", err)
			return
		}
		fmt.Println(info)
	*/
	i2 := 2
	switch i2 {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	case 3:
		fmt.Println("3")
	default:
		fmt.Println("default")
	}

	fmt.Println("主函数")

	m := make(map[string]int)
	m["a"] = 1
	fmt.Println(m)
	m["a"] = 2
	fmt.Println(m)
	delete(m, "a")
	fmt.Println(m)

	mm := map[string]int{"a": 1, "b": 2}
	fmt.Println(mm)

	s3 := "a"
	m[s3] = 3
	fmt.Println(m)
}

func TestFor() {
	ints := [...]int{1, 2, 3, 4, 5}
	for i := 0; i < len(ints); i++ {
		fmt.Printf("%d ", ints[i])
	}
	fmt.Printf("\n")
	for i, i2 := range ints {
		fmt.Printf("%d %d ", i, i2)
	}
	fmt.Printf("\n")

	s := "Hello World!"
	for _, v := range s {
		fmt.Printf("%c ", v)
	}
	fmt.Printf("\n")

	m := map[string]int{"a": 1, "b": 2}

	for k, v := range m {
		fmt.Printf("%s %d ", k, v)
	}
	fmt.Printf("\n")
	ch := make(chan int, 5)
	ch <- 1
	ch <- 2
	close(ch)
	for c := range ch {
		fmt.Printf("%d ", c)
	}
}

func TestSM() {
	type Company struct {
		Name string
	}
	type Person struct {
		Name    string
		age     int
		company Company
	}
	person := Person{Name: "a", age: 10, company: Company{Name: "a"}}
	person2 := Person{"b", 10, Company{Name: "b"}}
	fmt.Println(person.Name)
	fmt.Println(person2.age)
	person2.age = 11
	fmt.Println(person2.age)
	ptr := &person2
	ptr.age = 12

}

func (I Info) String() {
	fmt.Println(I)
}

func (I *Info) AgeAdd() {
	I.age++
}

func TestAnonymous() {
	s := struct {
		Name string
		age  int
	}{Name: "a", age: 10}

	fmt.Println(s)
}

func TestAdd(s1, s2 int) (int, int) {
	return s1 + s2, s1 - s2
}

func init() {
	fmt.Println("初始化")
}

func TestChange(i ...int) {
	for _, i3 := range i {
		fmt.Println(i3)
	}
}

func Calc(s1, s2 int) (x, y int) {
	x = s1 + s2
	y = s1 - s2
	return
}

func Pack() func(int) int {
	sum := 0
	return func(i int) int {
		sum += i
		return sum
	}
}
