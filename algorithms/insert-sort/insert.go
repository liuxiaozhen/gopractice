package main

//参考
//https://github.com/arnauddri/algorithms/blob/master/algorithms/sorting/bubble-sort/bubble.go
import (
	"fmt"

	"github.com/liuxiaozhen/gopractice/utils"
)

func Sort(arr []int) {
	n := len(arr)
	if n < 2 {
		return
	}
	for i := 1; i < n; i++ {
		for j := i; j > 0 && arr[j] < arr[j-1]; j-- {
			arr[j-1], arr[j] = arr[j], arr[j-1]
		}
	}
}

func main() {
	arrs := utils.GetList(21)
	Sort(arrs)
	fmt.Println(arrs)
}
