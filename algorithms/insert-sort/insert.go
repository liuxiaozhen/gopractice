package main

//参考
//https://github.com/arnauddri/algorithms/blob/master/algorithms/sorting/bubble-sort/bubble.go
import (
	"fmt"
	"math/rand"
	"time"
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
	arrs := getList(30)
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
