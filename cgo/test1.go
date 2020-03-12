package main

/*

#include<stdio.h>

// #cgo CFLAGS: -I/user/local/include //库对应头文件所在的目录加入头文件检索路径
// #cgo LDFLAGS: -L/user/local/lib -l库 //库所在目录加为链接库检索路径


// //定义/调用头文件

 //定义C函数
 int test_printf(char a)
 {
    printf("%c\r\n", a);
    return 0;
 }
*/
import "C"

import "fmt"

func main() {
	var c C.char = 'c'
	retc := C.test_printf(c)
	//这里retc 是C.int 格式，参考转换图则为go 对应着 int32
	retgo := int32(retc) //返回结果强化

	fmt.Printf("%T", retgo)

}
