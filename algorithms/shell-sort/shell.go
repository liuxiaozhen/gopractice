//希尔排序
package main

import (
	"fmt"
	"math/rand"
	"time"
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
	arrs := getList(13)
	Sort(arrs)
	fmt.Println(arrs)
}

func getList(n int) []int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	arr := make([]int, n, n)
	for i := 0; i < n; i++ {
		arr[i] = r.Intn(n)
	}
	return arr
}
