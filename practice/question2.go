//2、 以下代码有什么问题，说明原因
package main

import (
	"fmt"
)

type student struct {
	Name string
	Age  int
}

func main() {
	pase_student()
}

func pase_student() {
	m := make(map[string]*student)
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}

	for _, stu := range stus {
		m[stu.Name] = &stu
	}

	for _, stu := range m {
		fmt.Println(stu)
	}
}

/**
*
输出
&{wang 22}
&{wang 22}
&{wang 22}
原因：
range是复制操作 stu 是独立的变量
修改后：
for idx, stu := range stus {
	m[stu.Name] = &stus[idx]
}
*/
