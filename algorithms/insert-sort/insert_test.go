package main

import (
	"demo/utils"
	"fmt"
	"testing"
)

func TestInsertSort(t *testing.T) {
	list := utils.GetList(10000)
	Sort(list)
	for i := 0; i < len(list)-2; i++ {
		if list[i] > list[i+1] {
			fmt.Println(list)
			t.Error()
		}
	}
}

func BenchmarkInsertSort(b *testing.B) {
	list := utils.GetList(10000)
	for i := 0; i < b.N; i++ {
		Sort(list)
	}
}
