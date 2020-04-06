package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//初始化项目
func main(){
	r:= gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK,gin.H{"msg":"ok"})
	})
	r.Run()
}
