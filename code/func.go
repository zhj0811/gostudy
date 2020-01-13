package main

import "fmt"

//闭包函数测试
// func adder(y int) func() int {
// 	sum := 0
// 	return func() int {
// 		sum += y
// 		return sum
// 	}
// }

// func main() {
// 	pos, neg := adder(1), adder(2)
// 	for i := 0; i < 10; i++ {
// 		fmt.Println(
// 			pos(),
// 			neg(),
// 		)
// 	}
// }

func test(i int) func(j int) {

	return func(j int) {
		i += 1
		j += 1
		fmt.Printf("i: %v,j: %v\n", i, j)
	}
}

func main() {
	var i = 10
	print := test(i)
	var k = 11
	print(k)
	print(11)
	fmt.Println(k)
	fmt.Println(i)
}
