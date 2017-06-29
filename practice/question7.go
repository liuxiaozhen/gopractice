//7、请写出以下输入内容
package main

import (
	"fmt"
)

func main() {
	s := make([]int, 5)
	s = append(s, 1, 2, 3)
	fmt.Println(s)
	copy(s, []int{1, 2, 3})
	fmt.Println(s)
}

/**
output:
[0 0 0 0 0 1 2 3]
[1 2 3 0 0 1 2 3]
*/
