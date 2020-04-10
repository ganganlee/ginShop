package common

import "github.com/gin-gonic/gin"

//封装统一返回数据
func Response(c *gin.Context,status int,code int,data interface{},msg string){
	c.JSON(status,gin.H{
		"code":code,
		"data":data,
		"msg":msg,
	})
}

//封装统一返回成功数据
func ResponseSuccess(c *gin.Context,data interface{},msg string){
	c.JSON(200,gin.H{
		"code":200,
		"data":data,
		"msg":msg,
	})
}

//封装统一返回错误信息
func ResponseError(c *gin.Context,data interface{},msg string){
	c.JSON(200,gin.H{
		"code":400,
		"data":data,
		"msg":msg,
	})
}
