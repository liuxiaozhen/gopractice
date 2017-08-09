//冒泡排序实现
package main

import (
	"fmt"

	"github.com/liuxiaozhen/gopractice/utils"
)

func Sort(arr []int) {
	for j := len(arr) - 1; j > 1; j-- {
		swap := false
		for i := 1; i <= j; i++ {
			if arr[i-1] > arr[i] {
				arr[i-1], arr[i] = arr[i], arr[i-1]
				swap = true
			}
		}
		if swap == false {
			break
		}
	}
}

func main() {
	arrs := utils.GetList(21)
	Sort(arrs)
	fmt.Println(arrs)
}
