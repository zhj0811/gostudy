package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Printf("%#v", os.Args)
		panic("args must equal 3")
	}
	err := os.Rename(os.Args[1], os.Args[2])
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("success")
}
