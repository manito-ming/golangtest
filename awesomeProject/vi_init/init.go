package vi_init

import (
	"awesomeProject/vi_init/a" ////会现执行a里面的init
	"fmt"
)

func init() {
	fmt.Println("-------init-------")
}
func TestInit() {
	a.Testa()
}
