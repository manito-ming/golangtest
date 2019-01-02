package vv_sync

import (
	"fmt"
	"time"
)

func Mytimer() {
	time1 := time.NewTimer(time.Second)
	<-time1.C
	fmt.Println("Time 1 expired")

	time2 := time.NewTimer(time.Second) //在这停的时候,time2就是被stop
	go func() {
		<-time2.C
		fmt.Println("Timer 2 expired")
	}()
	stop2 := time2.Stop()
	if stop2 {
		fmt.Println("Time 2 stopped")
	}
}
func MyTicker() {
	tick := time.NewTicker(time.Millisecond * 500) //时钟中断
	go func() {
		for v := range tick.C {
			fmt.Println("Tick at", v)
		}
	}()
	time.Sleep(time.Millisecond * 1600)
	tick.Stop()
	fmt.Println("Tick stopped")
}
