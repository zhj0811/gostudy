package main

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
)

func main() {
	filename, line, funcname := "???", 0, "???"
	pc, filename, line, ok := runtime.Caller(0)
	if ok {
		funcname = runtime.FuncForPC(pc).Name()      // main.(*MyStruct).foo
		funcname = filepath.Ext(funcname)            // .foo
		funcname = strings.TrimPrefix(funcname, ".") // foo
	}

	fmt.Println(filename, funcname, line)

}
