package json

import (
	"encoding/json"
	"fmt"
)

type Message struct {
	Name string `json:"msg_name"`       // 对应JSON的msg_name
	Body string `json:"body,omitempty"` // 如果为空置则忽略字段
	Time int64  `json:"-"`              // 直接忽略字段  unmarshal tag 的-也不会被解析，但是会初始化其 零值：
}

func main() {
	m := Message{
		Name: "Alice",
		Body: "",
		Time: 1294706395881547000,
	}
	data, err := json.Marshal(m)
	if err != nil {
		fmt.Printf(err.Error())
		return
	}
	fmt.Println(data)
	fmt.Println(string(data)) //{"msg_name":"Alice"}
}

/*
Unmarshal 时结构体中未明确的interface ，可通过初始化结构体，同时若[]byte中无此字段，该字段信息仍然存在
*/
