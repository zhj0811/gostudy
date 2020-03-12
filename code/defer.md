func main() {
    fmt.Println(test())					//顺序5 	输出2
}

func test() (i int) {

    defer func() { i++ }()				//顺序4
    defer func() { fmt.Println(i) }()	//顺序3		输出1

    //todo
    //...

    fmt.Println(0)		//顺序1		输出0
    return 1			//顺序2
    //defer执行时机
}
输出结果0 1 2

defer执行是在return之后，且按照defer声明的先进后出顺序执行
当defer被声明时，其参数就会被实时解析
defer可以读取有名返回值

func test2() {
    defer func() {
        if err := recover(); err != nil {
            fmt.Println(err)
        }
    }()

    file, err := os.Open("path")
    if err != nil {
        panic(err)
    }

    defer file.Close()

    //todo
    //...

    return
    //defer执行时机
}

释放占用的资源 file.close
捕捉处理异常 defer func {	if recover() != nil {}	}()

func test() {
    i := 0
    defer fmt.Println(i)                       //输出 0		当defer被声明时，其参数就会被实时解析
    defer func(x int) { fmt.Println(x) }(i)    //输出 0
    defer func(x *int) { fmt.Println(*x) }(&i) //输出 1
    defer func() { fmt.Println(i) }()          //输出 1		defer函数体内的变量是在return后执行
    i++

    //todo
    //...

    fmt.Println(i) //输出 1
    return   
}
output: 1	1	1	0	0


type Test struct {
    name string
}

func (this *Test) Point() { // this  为指针
    fmt.Println(this.name)
}

func (this  Test) Value() { //this  为值类型 
    fmt.Println(this.name)
}

func test5() {
    ts := []Test{{"a"}, {"b"}, {"c"}}
    for _, t := range ts {
        defer t.Point() //输出 c c c
        defer t.Value() //输出 c b a
    }
}

output: c 	c	b	c	a	c


package main
 
import "fmt"
 
func main(){
    defer func(){ // 必须要先声明defer，否则不能捕获到panic异常
        fmt.Println("c")
        if err:=recover();err!=nil{
            fmt.Println(err) // 这里的err其实就是panic传入的内容，55
        }
        fmt.Println("d")
    }()
    f()
}
 
func f(){
    fmt.Println("a")
    panic(55)
    fmt.Println("b")
    fmt.Println("f")
}
output: 
a
c
55
d
exit code 0, process exited normally.
