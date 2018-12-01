package main

import (
	"fmt"
	"reflect"
)

type person struct {
	int
	name string
}

func (p person) Dis() { //*student不属于student 仅属于*student对象
	fmt.Println(p)
}

type student struct {
	person
	age int
	Sex int
}

func (s student) Sayhello(toname string) (string, int) { //如果方法不是大写第一个字母的公开方法，是无法被反射获取的
	return s.name + " say hello to " + toname, 1
}
func (s *student) Dis() {
	fmt.Println(s)
}

func Myref() {
	s := student{person: person{int: 1, name: "xxx"}, age: 23}
	t := reflect.TypeOf(s)        //获取字段属性以及字段相关的东西
	v := reflect.TypeOf(s).Name() //反射出名字
	x := reflect.TypeOf(s).Kind() ///反射出struct
	fmt.Println(t, " ", v, " ", x)

	v1 := reflect.ValueOf(s) //获取字段的值以及值相关的东西
	v11 := reflect.ValueOf(s).Kind()
	//x1:=reflect.ValueOf(s).Pointer()
	fmt.Println(v11, " ", v11, " ")

	if k := t.Kind(); k == reflect.Struct {
		for i := 0; i < t.NumField(); i++ { //可以去遍历出struct结构里面的东西
			key := t.Field(i)
			val := v1.Field(i)
			fmt.Println(key.Name, key.Type, "|", val)
		}
	}
	fmt.Println("===========type(s)==================")
	for i := 0; i < t.NumMethod(); i++ { //仅能获取非指针对象的引用方法
		m := t.Method(i)
		fmt.Println(m.Name, " ", m.Type)
	}
	fmt.Println("==========")

	v2 := reflect.ValueOf(&s).Kind()
	t3 := reflect.TypeOf(&s) //通过指针的
	v3 := reflect.ValueOf(&s)
	fmt.Println("valueof: ", v3, "kind:", v2, "typeof: ", t3)
	fmt.Println("============type(&)ptr==================")
	if k := t3.Kind(); k == reflect.Ptr { //指针通过elem获取内容
		tt := t3.Elem()
		vv := v3.Elem()
		for i := 0; i < tt.NumField(); i++ {
			key := tt.Field(i)
			val := vv.Field(i)
			fmt.Println(key.Name, key.Type, "|", val)
		}
	}

	fmt.Println("================type(&s)=================")
	for i := 0; i < t3.NumMethod(); i++ { //仅能获取非指针对象的引用方法
		m := t3.Method(i)
		fmt.Println(m.Name, m.Type)
	}

	fmt.Println("===========================================")
	fmt.Println(t.FieldByName("person"))
	fmt.Println(t.FieldByIndex([]int{0}))
	fmt.Println(t.FieldByIndex([]int{0, 0}), t.FieldByIndex([]int{0, 1}))
	m2, _ := t.MethodByName("Sayhello")
	fmt.Println(m2)
}
