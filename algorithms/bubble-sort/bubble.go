package main

//参考
//https://github.com/arnauddri/algorithms/blob/master/algorithms/sorting/bubble-sort/bubble.go
import (
	"fmt"
	"math/rand"
	"time"
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
	arrs := getList(1000)
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
