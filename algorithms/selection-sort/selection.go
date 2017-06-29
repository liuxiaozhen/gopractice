package main

//参考
//https://github.com/arnauddri/algorithms/blob/master/algorithms/sorting/bubble-sort/bubble.go
import (
	"demo/utils"
	"fmt"
)

func Sort(arr []int) {
	num := len(arr)
	for i := 0; i < num-1; i++ {
		min := arr[i]
		idx := i
		for j := i + 1; j < num; j++ {
			if arr[j] < min {
				min = arr[j]
				idx = j
			}
		}
		arr[i], arr[idx] = arr[idx], arr[i]
	}
}

func main() {
	arrs := utils.GetList(21)
	Sort(arrs)
	fmt.Println(arrs)
}
