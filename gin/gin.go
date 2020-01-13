package main

import "github.com/gin-gonic/gin"

func main() {

	// 运行模式
	gin.SetMode(gin.DebugMode)//ReleaseMode

    router := gin.Default()
    router.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "pong",	//访问localhost:8080/ping页面显示{"message": "pong"}
        })
    })
	router.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(200, "Hello %s", name)
    })	//访问localhost:8080/user/zh页面显示Hello zh
	router.GET("/welcome", func(c *gin.Context) {
        firstname := c.DefaultQuery("firstname", "Guest")
        lastname := c.Query("lastname")

        c.String(200, "Hello %s %s", firstname, lastname)
    })	//curl http://127.0.0.1:8080/welcome\?firstname\=zhao\&lastname\=jian
	router.POST("/form_post", func(c *gin.Context) {
        message := c.PostForm("message")
        usr := c.DefaultPostForm("usr", "anonymous")
        c.JSON(200, gin.H{
            "status":  gin.H{
                "status_code": 200,
                "status":      "ok",
            },
            "message": message,
            "usr":    usr,
        })
    })	//curl -X POST http://127.0.0.1:8080/form_post -H "Content-Type:application/x-www-form-urlencoded" -d "message=hello&usr=zh"
    //上传单个文件
	router.POST("/upload", func(c *gin.Context) {
        name := c.PostForm("name")
        fmt.Println(name)
        file, header, err := c.Request.FormFile("upload")
        if err != nil {
            c.String(http.StatusBadRequest, "Bad request")
            return
        }
        filename := header.Filename

        fmt.Println(file, err, filename)

        out, err := os.Create(filename)
        if err != nil {
            log.Fatal(err)
        }
        defer out.Close()
        _, err = io.Copy(out, file)
        if err != nil {
            log.Fatal(err)
        }
        c.String(http.StatusCreated, "upload successful")
    })	//curl -X POST http://127.0.0.1:8080/upload -F "upload=@/Users/ghost/Desktop/pic.jpg" -H "Content-Type: multipart/form-data"
	//上传多个文件
	router.POST("/multi/upload", func(c *gin.Context) {
        err := c.Request.ParseMultipartForm(200000)
        if err != nil {
            log.Fatal(err)
        }

        formdata := c.Request.MultipartForm 

        files := formdata.File["upload"] 
        for i, _ := range files { /
            file, err := files[i].Open()
            defer file.Close()
            if err != nil {
                log.Fatal(err)
            }

            out, err := os.Create(files[i].Filename)

            defer out.Close()

            if err != nil {
                log.Fatal(err)
            }

            _, err = io.Copy(out, file)

            if err != nil {
                log.Fatal(err)
            }

            c.String(http.StatusCreated, "upload successful")
        }
    })	//curl -X POST http://127.0.0.1:8080/multi/upload -F "upload=@/Users/ghost/Desktop/pic.jpg" -F "upload=@/Users/ghost/Desktop/journey.png" -H "Content-Type: multipart/form-data"

	v1 := router.Group("/v1")
    v1.GET("/login", func(c *gin.Context) {
        c.String(http.StatusOK, "v1 login")
    })
    v2 := router.Group("/v2")
    v2.GET("/login", func(c *gin.Context) {
        c.String(http.StatusOK, "v2 login")
    })
	router.Run(":8080") // listen and server on 0.0.0.0:8080(default)
}

