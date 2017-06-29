package main

import (
	"fmt"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	list := getList(10000)
	Sort(list)
	for i := 0; i < len(list)-2; i++ {
		if list[i] > list[i+1] {
			fmt.Println(list)
			t.Error()
		}
	}
}

func BenchmarkBubbleSort(b *testing.B) {
	list := getList(10000)
	for i := 0; i < b.N; i++ {
		Sort(list)
	}
}
