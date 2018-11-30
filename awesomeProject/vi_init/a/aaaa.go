package a

import (
	"awesomeProject/vi_init/b" //会执行b里面的
	"fmt"
)

func init() {
	fmt.Println("-------a------")
}
func Testa() {
	b.Testb()
}
