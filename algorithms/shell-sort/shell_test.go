package main

import (
	"fmt"
	"testing"

	"github.com/liuxiaozhen/gopractice/utils"
)

func TestShellSort(t *testing.T) {
	for num := 1; num < 10000; num++ {
		list := utils.GetList(num)
		Sort(list)
		for i := 0; i < len(list)-2; i++ {
			if list[i] > list[i+1] {
				fmt.Println(list)
				t.Error()
			}
		}
	}
}
