package vv_sync

import "sync"

type Once struct {
	m    sync.Mutex
	done uint32
}

func MyOnce2() {
	o := &Once{}
	go do2(o)
	go do2(o)
}
func do2(o *Once) {

}
