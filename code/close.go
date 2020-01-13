package main
import "fmt"
func main() {
    c1 := make(chan string)
    go func() {
       fmt.Println("go1")
       close(c1) 
       fmt.Println("go2")
    }()

    select {
    case <-c1:
      fmt.Println("in case")
    }
    fmt.Println("c1 is", <-c1)
}
