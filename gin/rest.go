package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	// "os"
	// "log"
	// "io"
	// "net/http"
	"github.com/unrolled/secure"
)

func main() {

	// 运行模式
	gin.SetMode(gin.DebugMode) //ReleaseMode

	router := gin.Default()
	router.Use(TlsHandler())

	router.GET("/ping", func(c *gin.Context) {
		fmt.Printf("url: %v", c.Request.URL)
		c.JSON(200, gin.H{
			"message": "pong", //访问localhost:8080/ping页面显示{"message": "pong"}
		})
	})
	router.RunTLS(":8080", "./server.crt", "./server.key")
	// router.Run(":8080") // listen and server on 0.0.0.0:8080(default)
}

func TlsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		secureMiddleware := secure.New(secure.Options{
			SSLRedirect: true,
			SSLHost:     "localhost:8080",
		})
		err := secureMiddleware.Process(c.Writer, c.Request)

		// If there was an error, do not continue.
		if err != nil {
			return
		}

		c.Next()
	}
}
