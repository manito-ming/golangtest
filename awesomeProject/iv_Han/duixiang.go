package iv_Han

import "fmt"

type self struct {
}

func (self *self) myplus(a int, b int) int {
	return a + b + 1
}
func (self self) myplus1(a int, b int) int {
	return a + b + 1
}

var c self ///申明是self类型的

func Test() {
	var a = c.myplus(1, 2)
	var b = c.myplus1(2, 2)
	fmt.Println("a:", a, "b:", b)
}
