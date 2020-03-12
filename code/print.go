package main

import (
	"fmt"
	"os"
)

func main(){
	fmt.Printf("Hello %d\n", 23)
	fmt.Fprint(os.Stdout, "Hello ", 23, "\n")
	fmt.Println("Hello", 23)
	fmt.Println(fmt.Sprint("Hello ", 23))
	var x uint64 = 1<<64 - 1
	fmt.Printf("%d %x; %d %x\n", x, x, int64(x), int64(x))
}

package main

import (
        "os"
        "github.com/op/go-logging"
)

func init(){
        format := logging.MustStringFormatter("%{shortfile} %{time:15:04:05.000} [%{module}] %{level:.4s} : %{message}")
        backend := logging.NewLogBackend(os.Stderr, "", 0)
        backendFormatter := logging.NewBackendFormatter(backend, format)
		//输出示例	main.go:19 14:49:17.235 [factorChaincode] DEBU : Init Chaincode...
		//不设置SetBackend默认输出示例	2018/11/26 15:17:07 Init Chaincode...
        logging.SetBackend(backendFormatter).SetLevel(logging.DEBUG, "factorChaincode")
}
var logger = logging.MustGetLogger("factorChaincode")
func main(){
        logger.Debug("Init Chaincode...")       
}



// package main

// import (
// 	"fmt"
// 	"github.com/op/go-logging"
// 	"os"
// )

// var logger = logging.MustGetLogger("logger-test")

// func init() {
// 	f, err := os.OpenFile("../tmp/log_error.txt", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	backend1 := logging.NewLogBackend(os.Stderr, "", 0)
// 	backend2 := logging.NewLogBackend(f, "", 0)
// 	var format = logging.MustStringFormatter(
// 		`%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}`,
// 	)

// 	backend1Formatter := logging.NewBackendFormatter(backend1, format)
// 	backend1Leveled := logging.AddModuleLevel(backend1Formatter)
// 	backend1Leveled.SetLevel(logging.DEBUG, "")

// 	//error信息定位至文件
// 	backend2Formatter := logging.NewBackendFormatter(backend2, format)
// 	backend2Leveled := logging.AddModuleLevel(backend2Formatter)
// 	backend2Leveled.SetLevel(logging.ERROR, "")

// 	// Set the backends to be used.
// 	logging.SetBackend(backend1Leveled, backend2Leveled)
// }
// func main() {
// 	logger.Debugf("debug %s", "secret")
// 	logger.Info("info")
// 	logger.Notice("notice")
// 	logger.Warning("warning")
// 	logger.Error("err")
// 	logger.Critical("crit")
// }
