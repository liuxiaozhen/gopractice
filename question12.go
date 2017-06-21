//下面的输出结果，及解释
package main

import (
	"fmt"
	"reflect"
)

func main() {
	var val interface{} = int64(58)
	fmt.Println(reflect.TypeOf(val))
	val = 50
	fmt.Println(reflect.TypeOf(val))
}

/**
Output:
int64
int
接口类型的变量底层是作为两个成员来实现，一个是type，一个是data。type用于存储变量的动态类型，data用于存储变量的具体数据。
一、显示的将类型为int64的数据58赋值给了interface类型的变量val，所以val的底层结构应该是：(int64, 58)。
二、字面量的整数在golang中默认的类型是int，所以这个时候val的底层结构就变成了：(int, 50)。
*/
