//1、写出下面代码输出内容。
package main

import (
	"fmt"
)

func main() {
	defer_call()
}

func defer_call() {
	defer func() { fmt.Println("打印前") }()
	defer func() { fmt.Println("打印中") }()
	defer func() { fmt.Println("打印后") }()
	panic("触发异常")
}

/**
打印后
打印中
打印前
panic: 触发异常

goroutine 1 [running]:
main.defer_call()
	E:/mygo/src/demo/question1.go:16 +0xc7
main.main()
	E:/mygo/src/demo/question1.go:8 +0x27
exit status 2
*/
