//下面代码会触发异常吗？请详细说明
package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(1)
	int_chan := make(chan int, 1)
	string_chan := make(chan string, 1)
	string_chan <- "hello"
	int_chan <- 1
	select {
	case value := <-int_chan:
		fmt.Println(value)
	case value := <-string_chan:
		panic(value)
	}
}

/***
*
会触发异常，
取决于select优先接收到信号 int_chan or string_chan

使用for监听所有可能的channel
for{
	select {
	case value := <-int_chan:
		fmt.Println(value)
	case value := <-string_chan:
		panic(value)
	}
}
*/
