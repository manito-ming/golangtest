package i_datatype

import (
	"fmt"
	"reflect"
	"unsafe"
)

func DD() {
	var str string = "hello 中国"
	fmt.Println("len: ", len([]rune(str))) ///rune是数数

	var as [5]int32 ///array不能动态扩容
	as[0] = 1
	as[1] = 2

	var as1 [3]int32
	as1[0] = 1

	var as2 = make([]int32, 2) ///slice  可以动态扩容  加make的都可以动态扩容
	as2[0] = 1
	as2[1] = 2
	fmt.Println("动态扩容需要使用append函数")
	//as2[2]=3
	//as2[3]=4
	as2 = append(as2, 3)        //可以动态扩容
	fmt.Println("容量", cap(as2)) ///直接扩容成为二倍
	as2 = append(as2, 4)
	fmt.Println(as2)

	var as3 = [...]int32{1, 2, 3}

	var mymap = make(map[int]int32, 10)
	mymap[0] = 10
	fmt.Println("相当于hashmap的", mymap[0])

	fmt.Println("数据类型 ", reflect.TypeOf(as), reflect.TypeOf(as1), reflect.TypeOf(as2), reflect.TypeOf(as3))
	fmt.Println("种类 ", reflect.TypeOf(as).Kind(), reflect.TypeOf(as1).Kind(), reflect.TypeOf(as2).Kind(), reflect.TypeOf(as3).Kind())

	var a = 1
	var b = 2
	a, b = b, a //交换
	fmt.Println(a, b)

	//var c,d,e int
	c, d, e := 5, 7, "as"
	_, d = d, c ///7被抛弃了
	fmt.Println(c, d)
	fmt.Println(c, d, e)
}
func TypeTest() {
	m := mkmap()
	println(m["a"])

	s := mkslice()
	println(s[0])

	a := 10
	b := byte(a)
	c := a + int(b) ///必须显式转换
	println("显式转换： ", c)

	x := 100
	p := (*int)(&x) //如果转换的目标是指针、单向通道或者没有返回值类型的函数类型，必须使用括号
	println(p)
}
func mkslice() []int {
	s := make([]int, 0, 10)
	s = append(s, 100)
	return s
}
func mkmap() map[string]int {
	m := make(map[string]int)
	m["a"] = 1
	return m
}
func Mystruct() {
	type person struct {
		id   int
		name string
	}
	var p = new(person)
	p.id = 99
	p.name = "xx"
	var i = uintptr(unsafe.Pointer(p))
	var j = (*int)(unsafe.Pointer(uintptr(i)))
	fmt.Println(*j)
}
