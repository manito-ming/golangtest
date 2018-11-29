package iv_Han

import (
	"fmt"
)

func TestSlice() {
	a := 2
	b := 2
	c := []int{1, 2}
	myfun31(c, a, b)
	cc := []int{1, 2}
	myfun31(c, cc...) //a,b拷贝到新的栈后被组装成了切片，并非真的传入了切片
	fmt.Println(a, b, c)

	/**闭包**/
	var cl = func(x int) func(int) int {
		var p int
		return func(y int) int {
			p++
			x++
			return x + y
		}
	}
	mycl1 := cl(1)
	m1 := mycl1(1)
	m2 := mycl1(1)
	mycl2 := cl(1)
	m3 := mycl2(1)
	fmt.Println(m1, m2, m3)

	/*指针转换**/
	//var add func(int,int)int=func(a,b int)(int){
	//	return a+b
	//}
	//var funptr=new (func(int,int)(int)) //定义函数指针

	//
	//type funcc func(int int)(int)
	//var funptr *func(int,int)(int)=new(func(int,int)int)
	//
	//var funptrxx *funcc = new(funcc)
	//
	//funptr = &add   ///值不能直接赋值，必须强转程指针类型才能赋值
	//
	//
	////funptrxx:=new(funcc)
	//funptrxx=(*funcc)(&add)
	//var ff funcc = add
	//
	//var opt2=func(a,b int,myfunc *funcc)(int){
	//	return (*myfunc)(a,b)
	//}
	//var opt1=func(a,b int,myfuncc funcc)(int){
	//	return myfunc(a,b)
	//}
	//opt2(1,2,(*funcc)(&add))
	//opt2(1,2,&ff)
	//opt1(1,2,add)

}

type op_func func(int, int) int //使用type自定义一种func(int,int)int函数的类型
func add(a int, b int) int {
	return a + b
}

func operator(op op_func, a, b int) int {
	return op(a, b)
}

func Funequals() {
	c := add
	fmt.Println("变量地址", c) //打印的是一个变量地址
	sum := c(1, 2)
	fmt.Println(sum)

	sum1 := operator(c, 1, 2)
	fmt.Println("使用自定义函数类型:", sum1)
}

func myfun31(c []int, a ...int) {
	c[0] = 123
	a[0] = 123
}

/*map函数*/
func Map() {
	var halffun = map[string]func(){
		"aaa:": func() {
			fmt.Println("aaa")
		},
		"bbb": func() {
			fmt.Println("bbb")
		},
	}

	halffun["aaa"]()
}

/*延迟函数*/
func Mydefer() int { ///defer就是延迟,压入栈里面，出栈的时候是先进后出
	var a = 1
	defer fmt.Println("a", a) //直接a就可以传进来
	defer fmt.Println("b", a)
	defer func() { //刚开始无法进入这个函数,a无法传进来
		// 只有等到返回值的时候，会把返回的值传进来
		a = 2
		fmt.Println(a)
	}()
	defer func() { //刚开始无法进入这个函数,a无法传进来
		a = 3
		fmt.Println(a)
	}()
	return a
}

/*struct*/
func Mystruct() {
	type person struct {
		id   int
		name string
	}
	var p1 = person{1, "Aaa"}
	var p2 = person{}
	p2.id = 2
	p2.name = "bbb"
	fmt.Println("struct  ", p1, p2)

	var p3 = new(person)
	//p3.id=3
	//p3.name="ccc"
	(*p3).id = 4
	(*p3).name = "ddd"
	fmt.Println("地址：", p3, "  数值：", *p3)

	/*嵌套匿名结构体*/
	var car = struct {
		id    int
		name  string
		style struct {
			color int
			model string
		}
	}{1, "aaa", struct {
		color int
		model string
	}{color: 1, model: string("xxx")}}
	fmt.Println("car: ", car)

	/*继承*/
	type Person struct {
		id   int
		name string
	}
	type stu struct {
		Person
		age int
	}
	type tea struct {
		Person
		cource string
	}
	var s = stu{
		Person: Person{id: 1, name: "aaa"},
		age:    22,
	}
	var t = tea{}
	t.Person.id = 2
	t.id = 3 //可以直接指向，不需要经过person
	t.Person.name = "bbb"
	t.name = "ccc"
	t.cource = "ddd"

	fmt.Println("stu: ", s, "tea: ", t)

}
