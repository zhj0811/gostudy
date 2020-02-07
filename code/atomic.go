package main

import (
	"fmt"
	"sync/atomic"
)

func main() {

	var a int32 = 1
	var b int32 = 2
	var c int32 = 2
	atomic.CompareAndSwapInt32(&a, a, b)
	fmt.Println("a, b:", a, b)
	atomic.CompareAndSwapInt32(&b, b, c)
	fmt.Println("a, b, c:", a, b, c)

	var x int32 = 100
	var y int32
	atomic.StoreInt32(&y, atomic.LoadInt32(&x))
	fmt.Println("x, y:", x, y)

	var j int32 = 1
	var k int32 = 2
	j_old := atomic.SwapInt32(&j, k)
	fmt.Println("old,new:", j_old, j)

	var v atomic.Value
	v.Store(100)
	fmt.Println("v:", v.Load())
}
