//给定2个现在给定下面一个条件，让你使用多种方式将奇偶合并，并打印自然数，这道题目很简单，
/*
go func() {
     for i := 1; i < 10; i++ {
         println(2*i - 1)
     }
}()
go func() {
     for i := 1; i < 10; i++ {
         println(2 * i)
     }
}()
*/

package main

import (
	"fmt"
	"time"
)

var count = 15

func main() {
	ch := make(chan int)
	go func() {
		for i := 1; i < count; i++ {
			ch <- 2*i - 1
		}
	}()

	go func() {
		for i := 1; i < count; i++ {
			ch <- 2 * i
		}
	}()

	go func() {
		for {
			msg := <-ch
			fmt.Println(msg)
			time.Sleep(time.Microsecond * 1000)
		}
	}()
	var input string
	fmt.Scanln(&input)
}

//参考goroutine调度
//http://studygolang.com/articles/1855
//http://www.tuicool.com/articles/fiqeUz
