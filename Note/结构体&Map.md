结构体
```go
//定义
type Company struct {  
    Name string  
}
type Person struct {  
    Name string  
    age  int
    company Company //嵌套  
}  
//实例化
person := Person{Name: "a", age: 10}  
person2 := Person{"b", 10}
//访问
fmt.Println(person.Name)
//直接修改
person2.age = 11
//指针修改
ptr := &person2  
ptr.age = 12
//方法
type Info struct {  
    Name string  
    age  int  
}
func (I Info) String() {  
    fmt.Println(I)  
}
info := Info{"1", 2}  
info.String()
//指针方法
func (I *Info) AgeAdd() {  
    I.age++  
}
info := Info{"1", 2}  
info.AgeAdd()  
fmt.Println(info)
//适合单次使用匿名结构体
s := struct {  
    Name string  
    age  int  
}{Name: "a", age: 10}
```

# Map
```go
//定义1
m := make(map[string]int)
//访问&修改
m["a"] = 1  
fmt.Println(m)  
m["a"] = 2
fmt.Println(m)
//删除
delete(m, "a")  
fmt.Println(m)  
//定义2  
mm := map[string]int{"a": 1, "b": 2}  
fmt.Println(mm)
```