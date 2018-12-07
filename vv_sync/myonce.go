package vv_sync

import (
	"fmt"
	"sync"
	"time"
)

func Myonce() {

	o := &sync.Once{}
	go do(o)
	go do(o)
	time.Sleep(time.Second)
}
func do(o *sync.Once) {
	fmt.Println("start do")
	o.Do(func() { //只会调用一次
		fmt.Println("Doing something...")
	})
	func() {
		fmt.Println("not once")
	}()
	fmt.Println("DO send")
}
