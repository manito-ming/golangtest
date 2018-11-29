package myStruct

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name  string `json:"name"` //json打包时的key为name
	Age   int    `json:"age"`
	Score int    `json:"score"`
}

func MyStruct() {
	var stu Student = Student{
		"stu",
		10,
		20,
	}
	data, err := json.Marshal(stu)
	if err != nil {
		fmt.Println("json encode failed,err ", err)
		return
	}
	fmt.Println(string(data))
}
