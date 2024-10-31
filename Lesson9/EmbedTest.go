package main

import (
	"embed"
	_ "embed"
	"fmt"
)

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
