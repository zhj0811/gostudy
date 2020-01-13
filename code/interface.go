//package main
//
//import (
//   "fmt"
//)
//func main() { 
//   var t1,t2 interface{} 
//   fmt.Println(t1==nil,t1==t2) 
//  
//   t1,t2=100,100
//   fmt.Println(t1==t2) 
//  
//   t1,t2=map[string]int{},map[string]int{} 
//   fmt.Println(t1==t2) 
//}
package main

import "fmt"

// notifier是一个定义了通知类行为的接口
type notifier interface {
　 notify()
}

// 定义user及user.notify方法
type user struct {
　 name　string
　 email string
}

func (u *user) notify() {
　 fmt.Printf("Sending user email to %s<%s>\n",
　　　 u.name,
　　　 u.email)
}

// 定义admin及admin.notify方法
type admin struct {
　 name　string
　 email string
}

func (a *admin) notify() {
　 fmt.Printf("Sending admin email to %s<%s>\n",
　　　 a.name,
　　　 a.email)
}

func main() {
　 // 创建一个user值并传给sendNotification
　 bill := user{"Bill", "bill@email.com"}
　 sendNotification(&bill)

　 // 创建一个admin值并传给sendNotification
　 lisa := admin{"Lisa", "lisa@email.com"}
　 sendNotification(&lisa)
}

// sendNotification接受一个实现了notifier接口的值
// 并发送通知
func sendNotification(n notifier) {
　 n.notify()
}
