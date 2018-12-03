package vii_myexception

import "fmt"

func testexception1() { //相当于不写try catch
	funca()
	funcb1()
	funcc()
}
func testexception21() { //相当于b21里面写了try catch,但是写错了
	funca()
	funcb21()
	funcc()
}
func funca() {
	fmt.Println("AAAAA")
}
func funcb1() {
	panic("B error") ///相当于throw
}
func funcc() {
	fmt.Println("CCCCC")
}

func funcb21() {
	panic("B error")
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()

	//分析：
	/*
		panic("B error")
		try{}catch(e){}
	*/
}
