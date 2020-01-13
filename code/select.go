package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("test")
	ch := make(chan bool)
	go process(ch)

	for {
		select {
		case ch <- true:
			fmt.Println("selct")
		}
	}

	fmt.Println("test3")
}

func process(c <-chan bool) {
	sum := 5050
	for {
		fmt.Println(sum)
		time.Sleep(1000000000)
		<-c
	}

}
