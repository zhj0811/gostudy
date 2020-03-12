package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// var lock = sync.Mutex{}
var lock = sync.RWMutex{}
var buff [10]int

func producer() {
	lock.Lock()
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		num := rand.Intn(100)
		fmt.Println("生产者生产了", num)
		buff[i] = num
		time.Sleep(time.Millisecond * 300)
	}
	lock.Unlock()
}

func consumer() {
	lock.Lock()
	for i := 0; i < 10; i++ {
		num := buff[i]
		fmt.Println("消费者消费到了", num)
	}
	lock.Unlock()
}

// go的互斥锁不支持重入，也不支持重复unlock
// 一个已经锁住的互斥锁不能再次被锁住，不管是同一个还是另一个goroutine
// 一个已经释放的互斥锁也不能再次被释放，不管是同一个还是另一个goroutine

// 读写锁的规则，类似于mysql的读写锁，本质就是一种运行关联操作并行和不可并行的一种逻辑实现。
// 可以随便读。多个goroutin同时读。执行写锁的操作的话，会进行等待直到所有的读锁都释放
// 写的时候，啥都不能干。不能读，也不能写
// sync.RWMutex 支持4个方法： Lock、Unlock （写锁）、RLock、RUnlock（读锁）

// 当读写锁被一个goroutine加了读锁，此读写锁还能被任何goroutine加读锁，但不能加写锁(会等待直到所有读锁释放)
// 当读写锁被一个goroutine加了写锁，任何goroutine将不能再对此读写锁加锁，不管是加读锁还是加写锁
// 当读写锁被一个goroutine加了读锁，在同一个gorounte内，它不能再被加写锁，但可以加读锁；同样
// 当读写锁被一个goroutine加了写锁，在同一个gorounte内，它不能再被加读锁，加写锁也不能。

// func main() {

// 	go producer()
// 	// 我们想要的是, 只有生产者生产了, 我们才能消费
// 	// 注意点: 在多go程中, 如果生产者生产的太慢, 那么消费者就会消费到错误的数据
// 	go consumer()
// 	// 注意点: 看上去通过给生产者以及消费者同时加锁就能解决, 只有生产完了才能消费
// 	//         但是取决于谁先执行加锁操作, 所以不完美
// 	for {
// 	}
// }

func main() {
	var rwm sync.RWMutex
	for i := 0; i < 3; i++ {
		go func(i int) {
			fmt.Println("try to lock read ", i)
			rwm.RLock()
			fmt.Println("get locked ", i)
			time.Sleep(time.Second * 2)
			// rwm.RLock()
			fmt.Println("try to unlock for reading ", i)
			// rwm.RUnlock()
			fmt.Println("unlocked for reading ", i)
			rwm.RUnlock()
		}(i)
	}
	time.Sleep(time.Millisecond * 1000)
	fmt.Println("try to lock for writing")
	rwm.Lock()
	fmt.Println("locked for writing")
}
