package myapi

import (
	"encoding/base64"
	"fmt"
)

func MyBase64() {
	data := "abc123!@#$%^&*()~"
	//base64编码
	sEnc := base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println("经过Base编码的： ", sEnc)
	sDec, _ := base64.StdEncoding.DecodeString(sEnc)
	fmt.Println("对上面的编码结果进行解码 ： ", sDec)

	//用在url之中的
	uEnc := base64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println("在url中的编码  ", uEnc)
	uDec, _ := base64.URLEncoding.DecodeString(uEnc)
	fmt.Println("解码", string(uDec))
}
