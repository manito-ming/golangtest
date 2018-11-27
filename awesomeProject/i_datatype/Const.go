package i_datatype

import "fmt"

const (
	a = iota //0
	b        //1
	c        //2
	d = "ha" //独立值，iota += 1
	e        //"ha"   iota += 1
	f = 100  //iota +=1
	g        //100  iota +=1
	h = iota //7,恢复计数
	i        //8
)

func ConstTest1() {
	fmt.Println("打印常量：", a, b, c, d, e, f, g, h, i) ///可当做枚举使用
}

func ConstTest() {
	const LENGTH int = 10
	const WIDTH int = 5
	var area int
	const a, b, c = 1, false, "str" ///多重赋值
	area = LENGTH * WIDTH
	fmt.Printf("面积为： %d", area)
	fmt.Println()
	fmt.Println(a, b, c)

}
