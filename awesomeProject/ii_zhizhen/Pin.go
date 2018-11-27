package ii_zhizhen

import (
	"fmt"
	"unsafe"
)

func Pin() {
	var a *int16 = new(int16)
	var b int16 = 0x0102
	a = &b
	fmt.Println("值", *a)
	fmt.Println((unsafe.Pointer(a)))

	var c uintptr = uintptr(unsafe.Pointer(a))

	var d = c + 1
	fmt.Println(c, " ", d)
	var c1 = (*int8)(unsafe.Pointer(c))
	var d1 = (*int8)(unsafe.Pointer(d))
	fmt.Println(*c1, *d1) ///黑科技，可以给地址加一，因为go指针不允许运算

	var arr = [...]int8{1, 2, 3, 4, 5}
	var sli = arr[2:5]
	fmt.Println(sli)
}
func Pin1() {
	x := 10
	var p *int = &x //获取地址，保存到指针变量
	*p += 20        //用指针间接引用，并更新对象,直接将x地址里面存放的数值改变
	fmt.Println(&x, x)
	fmt.Println(p, *p)
}
func Map() {
	m := map[string]int{"a": 1}
	fmt.Println(m["a"])
}
