package vv_sync

import (
	"fmt"
	"runtime"
	"runtime/debug"
	"sync"
	"sync/atomic"
)

func MyPooltest() {
	defer debug.SetGCPercent(debug.SetGCPercent(-1))
	var i int32 = 0
	newFunc := func() interface{} {
		return atomic.AddInt32(&i, 1)
	}
	pool := sync.Pool{New: newFunc}

	v1 := pool.Get()
	fmt.Println(v1)

	pool.Put(10)
	pool.Put(20)
	pool.Put(30)
	v2 := pool.Get()
	fmt.Println(v2)

	debug.SetGCPercent(100)
	runtime.GC()

	v3 := pool.Get() //因为上面调用了gc,所以没用的会被回收掉，返回的是默认的
	fmt.Println(v3)

	pool.New = nil //这样就pool变为空,所以想回收就置为空
	v4 := pool.Get()
	fmt.Println(v4)
}
