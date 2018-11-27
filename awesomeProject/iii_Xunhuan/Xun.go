package iii_Xunhuan

import (
	"fmt"
	"strconv"
)

func Xun() {

	var a = 1
	var c = strconv.Itoa(a) + "aaa" ///数字转换成字符串
	fmt.Println(c)

	var b = "aaa"
	v, err := strconv.Atoi(b) ///字符串转换成数字
	fmt.Println(v)
	fmt.Println(err)
	var d = v + 2
	fmt.Println("d", d)

	a = 100
	if a < 20 {
		fmt.Println("a小于20")
	} else {
		fmt.Println("a不小于20")
	}

	var grade string = "B"
	var marks int = 90
	switch marks {
	case 90:
		grade = "A"
	case 80:
		grade = "B"
	case 50, 60, 70:
		grade = "C"
	default:
		grade = "D"
	}
	switch {

	case grade == "A":
		fmt.Println("优秀")
	case grade == "B", grade == "C":
		fmt.Printf("良好\n")
	case grade == "D":
		fmt.Printf("及格\n")
	case grade == "F":
		fmt.Printf("不及格\n")
	default:
		fmt.Printf("差\n")
	}

	fmt.Println("第一种for循环")

	var aa int
	var bb int = 15
	for aa = 0; aa < 10; aa++ {
		fmt.Print("输出:", aa, "\t")
	}
	fmt.Println()

	fmt.Println("第二种for循环")
	for aa < bb {
		aa++
		fmt.Print(aa, "\t")
	}
	fmt.Println()
	fmt.Println("第三种for循环")
	//for {///是一个死循环，会死机的
	//	fmt.Println()
	//}

}
