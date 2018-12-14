package main

import (
	"fmt"
	"time"
	"sync"
)


var m = make(map[int] uint64);
var lock sync.Mutex//读写锁

type task struct {
	n int
}
func calc(t *task){
	var sum uint64
	sum =1
	for i := 1; i < t.n; i++ {
		sum*=uint64(i)

	}
	lock.Lock()
	m[t.n] = sum
	lock.Unlock()
}

func main() {
	for i := 0; i < 1000; i++ {
		t:=&task{n:i}
		go calc(t)
	}
	time.Sleep(10*time.Second)
	lock.Lock()
	for k,v:=range m{//遍历map
		fmt.Printf("%d!=%v\n",k,v)
	}
	lock.Unlock()
}
