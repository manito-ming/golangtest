package main

import (
	"fmt"

)

const (
	a = 1
	b = 2
	c = iota
	d
	e
)

type namestring struct {
	LL   int16   //2
	BB   int32   //4    没有这个的话，是直接在LL后面补6个空，但是加上这个之后是在LL后面补两个空，然后再加BB
	Lens int     //8
	Pos  uintptr //指针  8位
}

func main() {
	fmt.Println(a,b,c,d,e)
	//x1.AAA()
	var a=0x1111
	var b=0x2222
	var c=a+b   ///会在编译时直接计算
	fmt.Println(c)
	fmt.Println("----------------------DD-----------------------------")
	//i_datatype.DD()
	//i_datatype.TypeTest()
	//i_datatype.Mystruct()
	fmt.Println("----------------------ConstTest1-----------------------------")
	//i_datatype.ConstTest1()
	//fmt.Println("-------day2----------")
	//example3.TimeTest()
	//example3.StringTest()
	//fmt.Println("----------------------ConstTest-----------------------------")
	//i_datatype.ConstTest()
	//fmt.Println("----------------------Pin-----------------------------")
	//ii_zhizhen.Pin()
	//ii_zhizhen.Pin1()
	//ii_zhizhen.Map()
	//st.ClassTest()
	//fmt.Println("----------------------Xun-----------------------------")
	//
	//iii_Xunhuan.Xun()
	//fmt.Println("---------------------------------------------------")
	//type name=int
	//type xx name
	//var str namestring
	//fmt.Println(unsafe.Sizeof(str))
	//fmt.Println(unsafe.Sizeof(str.LL),unsafe.Sizeof(str.BB))//查看占几位
	//fmt.Println(unsafe.Alignof(str.LL),unsafe.Alignof(str.BB),unsafe.Alignof(str))  ///按几个字节对齐
	//fmt.Println(unsafe.Sizeof(str.LL),unsafe.Sizeof(str.BB),unsafe.Sizeof(str.Lens),unsafe.Sizeof(str.Pos))//查看占几位
	//
	//]]]]
	fmt.Println("-----------------hanshu---------------")
	//fmt.Println("---------------------Atm------------------------------------")
	//atm.Atm()
	//myStruct.MyStruct()
	//fmt.Println("----------------------init------------------------------")
	//vi_init.TestInit()
	//iv_Han.Funequals()

	//fmt.Println("----------------------init------------------------------")
	//iv_Han.TestSlice()
	//iv_Han.Map()
	//fmt.Println("返回值",iv_Han.Mydefer())

	//iv_Han.Mystruct()
	fmt.Println("----------------------------接口------------------")
	//fmt.Println("----------------------Myintf------------------------------")
	//myintf.Myintf()
	//fmt.Println("----------------------TestCall------------------------------")
	//myintf.TestCall()
	//myintf.TestAABBCC()
	//iii_myinterface.Myintf2()
	//fmt.Println("----------------------面向对象------------------------------")
	//iv_Han.Test()
	//fmt.Println("----------------------反射--------------------------------")
	//viii_reflation.Myref()
	fmt.Println("----------------------并发-------------------------------")
	//vv_goroution.Mygroutine()
	//vv_goroution.Myselect()
	//fmt.Println("----------------------sync-------------------------------")
	//vv_sync.Myonce()
	//vv_sync.MyPooltest()
	//vv_sync.Mytimer()
	//vv_sync.MyTicker()
	fmt.Println("----------------api---------------")
    //myapi.MyBase64()
    //myapi.Myjson()
}
