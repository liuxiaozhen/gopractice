//希尔排序
package main

import (
	"demo/utils"
	"fmt"
)

func Sort(arr []int) {
	increment := len(arr) / 2
	for increment > 0 {
		for i := increment; i < len(arr); i++ {
			j := i
			temp := arr[i]

			for j >= increment && arr[j-increment] > temp {
				arr[j] = arr[j-increment]
				j = j - increment
			}
			arr[j] = temp
		}
		increment = int(increment / 2)
	}
}

func main() {
	arrs := utils.GetList(21)
	Sort(arrs)
	fmt.Println(arrs)
}
