package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	// "os"
	// "log"
	// "io"
	"net/http"
	// "github.com/unrolled/secure"
)

func main() {

	// 运行模式
	gin.SetMode(gin.DebugMode) //ReleaseMode

	router := gin.Default()

	router.POST("/post", func(c *gin.Context) {
		var b req
		// if err := c.Bind(&b); err != nil {
		if err := c.ShouldBindJSON(b); err != nil {
			// if err := c.ShouldBindWith(&b, binding.JSON); err != nil {
			// if err := c.ShouldBind(&b); err != nil {
			fmt.Println(err.Error())
			c.String(http.StatusOK, `the body should be req`)
			return
		}
		fmt.Printf("%v\n", b.key)
		fmt.Printf("%v\n", b)
		c.JSON(200, b)
	})

	router.Run(":8080")
}

type req struct {
	key   string `json:"key" binding:"required"`
	value string `json:"value" binding:"required"`
}
