package test

import (
	"fmt"
	"sync"
	"time"
	"testing"
)

func onces() {
	//var t *testing.T
	fmt.Println("onces")
}
func onced() {
	//var t *testing.T
	fmt.Println("onced")
}

func TestOnce (t *testing.T){
	var once sync.Once
	for i, v := range make([]string, 10) {
		once.Do(onces)
		t.Log("count:", v, "---", i)
	}
	for i := 0; i < 10; i++ {

		go func() {
			//sync.Once.Do(f func())是一个挺有趣的东西,能保证once只执行一次，无论你是否更换once.Do(xx)这里的方法,这个sync.Once块只会执行一次。
			once.Do(onced)	//未执行
			t.Log("213")
		}()
	}
	time.Sleep(4000)
}
//func TestMain (t *testing.M){
//	var once sync.Once
//	for i, v := range make([]string, 10) {
//		once.Do(onces)
//		fmt.Println("count:", v, "---", i)
//	}
//	for i := 0; i < 10; i++ {
//
//		go func() {
//			once.Do(onced)
//			fmt.Println("213")
//		}()
//	}
//	time.Sleep(400000)
//}

//fmt.Println(time.Now().Unix()) //获取当前秒
//fmt.Println(time.Now().UnixNano())//获取当前纳秒
//fmt.Println(time.Now().UnixNano()/1e6)//将纳秒转换为毫秒
//fmt.Println(time.Now().UnixNano()/1e9)//将纳秒转换为秒
//c := time.Unix(time.Now().UnixNano()/1e9,0) //将毫秒转换为 time 类型