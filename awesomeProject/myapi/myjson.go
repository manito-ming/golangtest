package myapi

import (
	"encoding/json"
	"fmt"
)

type Response1 struct {
	Page   int
	Fruits []string
}
type Response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func Myjson() {
	//转化为json
	bolB, _ := json.Marshal(true) //转换为json格式
	fmt.Println(string(bolB))
	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))
	//将数组转化为json
	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))
	//将map转换为json
	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))
	//fmt.Println("使用gojson通过k获取v： ", gojson.Json(string(mapB)).Get("apple"))
	//
	//json := `{"from":"en","to":"zh","trans_result":{"src":"today","dst":"\u4eca\u5929"},"result":["src","today","dst","\u4eca\u5929"]}`
	//c2 := gojson.Json(json).Get("trans_result").Get("dst") //通过gojson来解析go
	//fmt.Println(c2)                                        //&{今天}

}
