package main

/*

#include<stdio.h>
#include<stdlib.h>
void c_print(char* str) {
	printf("%s\n", str);
}

*/
import "C" //必须单起一行，并且紧跟注释行之后
import "unsafe"

func main() {
	s := "hello Cgo"
	cs := C.CString(s)
	C.c_print(cs)
	defer C.free(unsafe.Pointer(cs))
}
