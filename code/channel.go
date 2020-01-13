package main

import(
        "fmt"
        "time"
)
func produce(p chan<- int) {
    for i := 0; i < 10; i++ {
        p <- i
        fmt.Println("send:", i)
    }
}
func consumer(c <-chan int) {
    for i := 0; i < 10; i++ {
        v := <-c
        fmt.Println("receive:", v)
    }
}
func main() {
    ch := make(chan int, 10)
    go produce(ch)
    go consumer(ch)
    time.Sleep(1 * time.Second)
}
/*
创建无缓冲channel
chreadandwrite :=make(chan int)	//
chonlyread := make(<-chan int) 	//创建只读channel 
chonlywrite := make(chan<- int) //创建只写channel
创建缓冲channel
chreadandwrite :=make(chan int，1)	//第一个值会存储在缓冲区
chonlyread := make(<-chan int) 	//创建只读channel 
chonlywrite := make(chan<- int) //创建只写channel
channel的机制是先进先出，如果给channel赋值了，那么必须要读取它的值，不然就会造成阻塞，当然这个只对无缓冲的channel有效。
对于有缓冲的channel，发送方会一直阻塞直到数据被拷贝到缓冲区；如果缓冲区已满，则发送方只能在接收方取走数据后才能从阻塞状态恢复。

关闭一个未初始化(nil) 的 channel 会产生 panic
重复关闭同一个 channel 会产生 panic
向一个已关闭的 channel 中发送消息会产生 panic
从已关闭的 channel 读取消息不会产生 panic，且能读出 channel 中还未被读取的消息，若消息均已读出，则会读到类型的零值。从一个已关闭的 channel 中读取消息永远不会阻塞，并且会返回一个为 false 的 ok-idiom，可以用它来判断 channel 是否关闭
关闭 channel 会产生一个广播机制，所有向 channel 读取消息的 goroutine 都会收到消息
ch := make(chan int, 10)
ch <- 11
ch <- 12

close(ch)

for x := range ch {
    fmt.Println(x)
}

x, ok := <- ch
fmt.Println(x, ok)

-----
output:

11
12
0 false
-------
ch := make(chan int, 10)
ch <- 11
ch <- 12

close(ch)
x, ok := <- ch
fmt.Println(x, ok)

-----
output:
11 true
go程序会优先执行主线程，主线程执行完成后，主线程没有提供足够的缓冲时间输出goroutine的结果，程序会很快退出
*/


对于非缓存chennel
forever := make(chan int)
forever <- 1 //阻塞后续运行
go func(){	//因forever <- 1  main(routine)阻塞挂起，go func()未执行
	<-forever
}()

forever := make(chan int)
go func(){
	forever <- 1 
}					
<-forever	// 非缓冲channel写（读）操作在go func 中，非缓冲channel读（写）在go func 才不会阻塞


	forever := make(chan bool)
	go func() {
		fmt.Println("test")
		<-forever
		fmt.Println("test")	//主线程已结束，程序退出未执行
		<-forever
		fmt.Println("test")	//未执行
		// forever <- true
	}()
	forever <- true

//output：test