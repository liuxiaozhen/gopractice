//11、以下代码打印出来什么内容，说出为什么。。。
package main

import (
	"fmt"
	"reflect"
)

type People interface {
	Show()
}

type Student struct {
	Name int
}

func (stu *Student) Show() {
	fmt.Println("abc")
}

func live() People {
	var stu *Student
	return stu
}

func main() {
	fmt.Println(reflect.TypeOf(live()))
	fmt.Println(reflect.ValueOf(live()))
	if live() == nil {
		fmt.Println("AAAAAAA")
	} else {
		fmt.Println("BBBBBBB")
	}
}

/**
*
*main.Student
<nil>
BBBBBBB

因为interface虽然看起来像指针，但并不是指针。interface变量仅在类型和值为“nil”时才为“nil”。
interface的类型和值会根据用于创建对应interface变量的类型和值的变化而变化。当你检查一个interface变量是否等于“nil”时，这就会导致未预期的行为。
*/
