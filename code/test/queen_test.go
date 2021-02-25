package test

import (
	"fmt"
	"testing"
)

// 8皇后算法问题 回溯算法
// 回溯算法求解八皇后问题的原则是：有冲突解决冲突，没有冲突往前走，无路可走往回退，走到最后是答案。
func TestQueen(t *testing.T) {
	var balance = [8]int{0, 0, 0, 0, 0, 0, 0, 0}
	queen(balance, 0)
}
func queen(a [8]int, cur int) {
	if cur == len(a) {
		fmt.Print(a)
		fmt.Println()
		return
	}
	for i := 0; i < len(a); i++ {
		a[cur] = i
		flag := true
		for j := 0; j < cur; j++ {
			ab := i - a[j]
			temp := 0
			if ab > 0 {
				temp = ab
			} else {
				temp = -ab
			}
			//if a[j] == a[cur] || temp == cur-j { 同义
			if a[j] == i || temp == cur-j {
				flag = false
				break
			}
		}
		if flag {
			queen(a, cur+1)
		}
	}
}
