package main

import (
	"fmt"
	"testing"

	"github.com/liuxiaozhen/gopractice/utils"
)

func TestBubbleSort(t *testing.T) {
	list := utils.GetList(10000)
	Sort(list)
	for i := 0; i < len(list)-2; i++ {
		if list[i] > list[i+1] {
			fmt.Println(list)
			t.Error()
		}
	}
}

func BenchmarkBubbleSort(b *testing.B) {
	list := utils.GetList(10000)
	for i := 0; i < b.N; i++ {
		Sort(list)
	}
}
