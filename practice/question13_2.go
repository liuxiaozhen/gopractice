package main

import (
	"fmt"
	"runtime"
	// "time"
)

var count = 15

func main() {
	runtime.GOMAXPROCS(1)
	go func() {
		for i := 1; i < count; i++ {
			fmt.Println(2 * i)
			runtime.Gosched()
		}
	}()

	go func() {
		for i := 1; i < count; i++ {
			fmt.Println(2*i - 1)
			runtime.Gosched()
		}
	}()

	var input string
	fmt.Scanln(&input)
}
