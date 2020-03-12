package main

import (
    "flag"
    "fmt"
)

func main() {
    backup_dir := flag.String("b", "/home/default_dir", "backup path")
    debug_mode := flag.Bool("d", false, "debug mode")
    flag.Parse()
    fmt.Println("backup_dir: ", *backup_dir)
    fmt.Println("debug_mode: ", *debug_mode)
}

/**																|##go build main.go exec.out
启动命令：go run main.go -b /home/backup						|##	  $./exec.out --help
输出结果：														|##		-b string
	backup_dir:  /home/backup  // 覆盖了默认路径				|##			 backup path (default "/default")
	debug_mode:  false         // 启动命令无-d flag，启用默认值	|##		-d	 debug mode
*/
	
##package flag	//从命令行获取参数
##@param name	//flag name		./exec.out --help中flag
##@param value	//default value
##@param usage	//usage string 	./exec.out --help中flag描述信息	
##func String(name string, value string, usage string) *string {}


