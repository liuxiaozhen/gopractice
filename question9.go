//9、下面的迭代会有什么问题？
package main

import (
	"fmt"
	"sync"
)

func (set *threadSafeSet) Iter() <-chan interface{} {
	ch := make(chan interface{})
	go func() {
		set.RLock()

		for elem := range set.s {
			ch <- elem
		}

		close(ch)
		set.RUnlock()

	}()
	return ch
}

/**
。。。
*
*/
