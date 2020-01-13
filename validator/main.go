package main

import (
	"fmt"
	"github.com/smokezl/govalidators"
	"github.com/zhaojianpeerfintech/validator/define"
)

func main() {
	validator := govalidators.New()
	student := &define.Student{
		Uid:          1234567,
		Name:         "张三1111",
		Age:          31,
		Sex:          "male1",
		Email:        "@qq.com",
		PersonalPage: "www.abcd.com",
		Hobby:        []string{"swimming", "singing"},
		CreateTime:   "2018-03-03 05:60:00",
		Class: []define.Class{
			define.Class{
				Cid:       12345678,
				Cname:     "语文",
				BeginTime: "13:00",
			},
			define.Class{
				Cid:       22345678,
				Cname:     "数学",
				BeginTime: "13:00",
			},
			define.Class{
				Cid:       32345678,
				Cname:     "数学",
				BeginTime: "13:60",
			},
		},
	}
	errList := validator.Validate(student)
	if errList != nil {
		for _, err := range errList {
			fmt.Println("err:", err)
		}
	}
}
