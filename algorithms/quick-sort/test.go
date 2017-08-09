// test
package main

import (
	"fmt"
)

type es struct{}

func main() {
	a := struct{}{} // not the zero value, a real new struct{} instance
	b := struct{}{}
	fmt.Println(a == b) // true

	c := es{} // not the zero value, a real new struct{} instance
	d := es{}
	fmt.Println(c == d) // true
}
