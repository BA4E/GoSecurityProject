package main

import (
	"fmt"
	"github.com/golang-module/dongle"
)

func main() {
	for q := false; q != true; {
		fmt.Println("输入en编码，de解码，exit退出")
		var op string
		_, err := fmt.Scanf("%s", &op)
		if err != nil {
			return
		}

		switch op {
		case "en":
			for q := false; q != true; {
				var op2 string
				fmt.Println("输入1编码，exit退出")
				_, err := fmt.Scanf("%s", &op2)
				if err != nil {
					return
				}
				switch op2 {
				case "1":
					fmt.Println("输入要编码的字符串")
					var s string
					_, err := fmt.Scanf("%s", &s)
					if err != nil {
						return
					}
					to := dongle.Encode.FromString(s).ByBase64().ToString()
					fmt.Printf("编码结果: %s\n", to)
				case "exit":
					q = true
				default:
					fmt.Println("错误输入")
				}
			}
		case "de":
			for q := false; q != true; {
				var op2 string
				fmt.Println("输入1解码，exit退出")
				_, err := fmt.Scanf("%s", &op2)
				if err != nil {
					return
				}
				switch op2 {
				case "1":
					fmt.Println("输入要解码的字符串")
					var s string
					_, err := fmt.Scanf("%s", &s)
					if err != nil {
						return
					}
					from := dongle.Decode.FromString(s).ByBase64().ToString()
					fmt.Printf("编码结果: %s\n", from)
				case "exit":
					q = true
				default:
					fmt.Println("错误输入")
				}
			}
		case "exit":
			q = true
		default:
			fmt.Println("错误输入")
			return
		}

	}

}
