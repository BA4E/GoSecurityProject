embed支持把配置文件内嵌到运行程序中
```go
package main  
  
import (  
    "embed"  
    _ "embed"  
    "fmt")  
  
//go:embed test.txt  
var s string  
var s2 string  
  
//go:embed byteTest  
var b []byte  
  
//go:embed FStest stringTest  
var f embed.FS  
  
//  
  
func main() {  
    fmt.Println(s, s2)  
    fmt.Println(b)  
    file, _ := f.ReadFile("FStest")  
    readFile, _ := f.ReadFile("stringTest")  
    fmt.Println(file)  
    fmt.Println(readFile)  
}
//This is BA4E. 
//[49 50 51 52 53 54]
//[97 97 97 97 97 97 97 13 10 98 98 98 98 98 98 98 98 13 10 99 99 99 99 99 99 99]
//[103 104 106 100 97 13 10 97 115 100 13 10 97 13 10 102 97 13 10 115 102 13 10]
```