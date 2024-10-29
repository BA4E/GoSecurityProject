package main

import "fmt"

type Student struct {
	Name string
	Age  int
}

func (stu Student) Print() string {
	return fmt.Sprintf("Name: %s, Age: %d", stu.Name, stu.Age)
}

func (stu *Student) Change(name string) {
	stu.Name = name
}

func main() {
	student := Student{
		Name: "BA4E",
		Age:  20,
	}

	fmt.Println(student.Print())

	student.Change("ba4e")

	fmt.Println(student.Print())

	s := &Student{
		Name: "test",
		Age:  0,
	}

	fmt.Println(s.Print())

	s.Change("TEST")

	fmt.Println(s.Print())

	//Name: BA4E, Age: 20
	//Name: ba4e, Age: 20
	//Name: test, Age: 0
	//Name: TEST, Age: 0

}
