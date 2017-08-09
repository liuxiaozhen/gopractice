package main

//参考
//https://github.com/arnauddri/algorithms/blob/master/algorithms/sorting/bubble-sort/bubble.go
import (
	"fmt"

	"github.com/liuxiaozhen/gopractice/utils"
)

func Sort(arr []int) {
	quickSort(arr, 0, len(arr)-1)
}

func quickSort(arr []int, left int, right int) {
	key := arr[left]
	p := left
	i, j := left, right

	for i <= j {
		for j >= p && arr[j] >= key {
			j--
		}
		if j >= p {
			arr[p] = arr[j]
			p = j
		}

		if arr[i] <= key && i <= p {
			i++
		}
		if i <= p {
			arr[p] = arr[i]
			p = i
		}
	}
	arr[p] = key
	if p-left > 1 {
		quickSort(arr, left, p-1)
	}
	if right-p > 1 {
		quickSort(arr, p+1, right)
	}
}

func main() {
	arrs := utils.GetList(32)
	Sort(arrs)
	fmt.Println(arrs)
}
