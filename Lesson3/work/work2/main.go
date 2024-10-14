package main

import (
	"fmt"
)

type Books struct {
	Name   string
	Author string
	Pages  int
}

type Library map[string]Books

func main() {
	m := make(Library)
	book := CreateBook("Go编程", "daem0nu", 200)
	book2 := CreateBook("C编程", "aaa", 150)
	fmt.Printf("%+v\n", book)
	m.AddBook(book)
	m.AddBook(book2)
	fmt.Printf("%+v\n", m)
	m.ListBook()
	//m.EditBook("C编程")
	//m.ListBook()
	m.RemoveBook("C编程")
	m.ListBook()

}
func CreateBook(Name string, Author string, Pages int) Books {
	return Books{Name: Name, Author: Author, Pages: Pages}
}

func (L Library) AddBook(b Books) {
	L[b.Name] = b
}

func (L Library) ListBook() {
	for _, books := range L {
		fmt.Printf("书名: %s,作者: %s,页数: %d\n", books.Name, books.Author, books.Pages)
	}
}

func (L Library) RemoveBook(Name string) {
	delete(L, Name)
}

func (L Library) EditBook(Name string) {
	books := L[Name]
	if L[Name].Name == "" {
		fmt.Println("Not Found")
		return
	}
	fmt.Println("输入要修改的部分:Name,Author,Pages")
	var s string
	fmt.Scanf("%s", &s)
	switch s {
	case "Name":
		var name string
		fmt.Println("输入修改后的书名")
		fmt.Scanf("%s", &name)
		books.Name = name
	case "Author":
		var author string
		fmt.Println("输入修改后的作者")
		fmt.Scanf("%s", &author)
	case "Pages":
		var pages int
		fmt.Scanf("%d", &pages)
		books.Pages = pages
	default:
		fmt.Println("错误输入")
	}
	L[Name] = books
}
