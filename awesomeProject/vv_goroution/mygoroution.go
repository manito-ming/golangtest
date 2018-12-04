package vv_goroution

import (
	"fmt"
	"strconv"
	"time"
)

func myrun1(c *chan bool) {
	fmt.Println("go myrun1")
	*c <- true
}
func myrun2(c *chan int) {
	list := make([]int, 0)
	for i := 0; i < 2; i++ { //往外面取出两个
		v, _ := <-*c
		fmt.Println("myrun2:" + strconv.Itoa(v))
		list = append(list, v)
	}
	fmt.Println(list)
	*c <- 0
}

/*限制，只能向channel里面写数据*/
func send(c chan<- int) {
	for i := 0; i < 5; i++ {
		c <- i
		fmt.Println("send" + strconv.Itoa(i))
	}
}

//读取
func recv(c <-chan int) {
	for i := range c {
		fmt.Println("recv ", i)
	}
}
func Mygroutine() {
	c1 := make(chan bool)
	go myrun1(&c1)
	v, err := <-c1
	fmt.Println(v, err)
	fmt.Println("===================================================")

	c2 := make(chan int, 2)
	go myrun2(&c2)
	c2 <- 1
	c2 <- 2
	c2 <- 3                     //此处多写了一个,但是没有发现,这是一个bug
	time.Sleep(2 * time.Second) //如果不谢sleep,下面的就会被消费了
	val, _ := <-c2

	fmt.Println("res ", val)

	crw := make(chan int)
	go send(crw)
	go recv(crw)

}
