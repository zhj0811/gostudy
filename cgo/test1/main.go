package main

//使用#cgo定义库路径

/*
#cgo CFLAGS: -I .
#cgo LDFLAGS: -L . -lhello
#include "hello.h"
*/
import "C"

func main() {
	C.hello()
}
