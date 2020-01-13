package main

import (
	"fmt"
	"os"
	"os/signal"
)

func main() {
	c := make(chan os.Signal)
	signal.Notify(c)
	fmt.Println("test1")
	// <-c
	s := <-c
	fmt.Println("get signal: ", s) //docker-compose down会执行
}
