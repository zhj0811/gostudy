package test

import (
	"testing"
)

func TestSlice(t *testing.T) {
	s := []int{5}
	s = append(s, 7)
	t.Log("cap(s) =", cap(s), "len(s)=", len(s), "ptr(s) =", &s[0])
	s = append(s, 9)
	t.Log("cap(s) =", cap(s), "len(s)=", len(s), "ptr(s) =", &s[0])
	x := append(s, 11)
	t.Log("cap(s) =", cap(s), "len(s)=", len(s), "ptr(s) =", &s[0], "ptr(x) =", &x[0])
	y := append(s, 12)
	t.Log("cap(s) =", cap(s), "len(s)=", len(s), "ptr(s) =", &s[0], "ptr(y) =", &y[0])
	z := make([]int, 2, 4)
	t.Log("cap(z) =", cap(z), "len(z)=", len(z), "ptr(z) =", &z[0])
}


//声明了一个函数类型
type testInt func(int) bool

func isOdd(integer int) bool {
	if integer%2 == 0 {
		return false
	}
	return true
}

func isEven(integer int) bool {
	if integer%2 == 0 {
		return true
	}
	return false
}

//声明的函数在这个地方当做了一个参数
func filter(slice []int, f testInt) []int {
	var result []int
	for _, value := range slice {
		if f(value) {
			result = append(result, value)
		}
	}
	return result
}

//map slice 为引用类型
func TestSliceSec(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5, 7}
	t.Log("slice = ", slice)
	//将函数当做值来传递
	odd := filter(slice, isOdd)
	t.Log("奇数是:", odd)
	//将函数当做值来传递
	even := filter(slice, isEven)
	t.Log("偶数是:", even)
}