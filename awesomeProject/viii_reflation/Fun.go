package main

import (
	"fmt"
	"reflect"
)

type X struct {

}

func (X)Test(x,y int)(int,error)  {
	return x+y,fmt.Errorf("err: %d",x+y)
}
//反射调用方法
func Cal() {
	var a X
	v:=reflect.ValueOf(&a)
	m:=v.MethodByName("Test")

	in:=[]reflect.Value{
		reflect.ValueOf(2),
		reflect.ValueOf(2),

	}
	out:=m.Call(in)
	for _,v:=range out{
		fmt.Println(v)
	}
}

func (X) Format(s string,a...interface{})string {
	return fmt.Sprintf(s,a...)
}
//反射调用变参的方法
func Cal1() {
	var a X
	v:=reflect.ValueOf(&a)
	m:=v.MethodByName("Format")
	in:=[]reflect.Value{//所有的参数都必须处理
		reflect.ValueOf("%s=%d"),
		reflect.ValueOf("x"),
		reflect.ValueOf(100),
	}
	out:=m.Call(in)//使用普通的Call方法处理
	fmt.Println(out)
fmt.Println("------------")
	in=[]reflect.Value{//所有的参数都必须处理
		reflect.ValueOf("%s=%d"),
		reflect.ValueOf([]interface{}{"x",100}),//使用interface就可以
		//reflect.ValueOf([]interface{}{"ss",200}),
	}
	out=m.CallSlice(in)//使用CallSlice处理
	fmt.Println(out)
}

func main() {
	Cal()
	fmt.Println("==================")
	Cal1()
}
