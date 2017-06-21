//下面的代码有什么问题?
package main

import (
	"fmt"
	"sync"
)

func main() {
	u := new(UserAges)
	u.ages = make(map[string]int)
	u.Add("Li", 22)
	fmt.Println(u.Get("Li"))

}

type UserAges struct {
	ages map[string]int //ages 需要初始化
	sync.Mutex
}

func (ua *UserAges) Add(name string, age int) {
	ua.Lock()
	defer ua.Unlock()
	ua.ages[name] = age
}

func (ua *UserAges) Get(name string) int {
	ua.RLock()
	defer ua.RUnlock()
	if age, ok := ua.ages[name]; ok {
		return age
	}
	return -1
}

/***
go中map并发使用是不安全的，当你使用goroutine同时对一个map进行读写操作时，不确定会发生什么（由于读写执行顺序不确定造成的）．针对这种情况，我们要添加读写锁对sync.RWMutex其进行同步．
type UserAges struct {
	ages map[string]int //ages 需要初始化
	sync.RWMutex
}

func (ua *UserAges) Get(name string) int {
	ua.RLock()
	defer ua.RUnlock()
*/
