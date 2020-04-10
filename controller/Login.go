package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {
	isLogin := c.GetBool("isLogin")

	c.HTML(http.StatusOK, "login.html", gin.H{
		"title": "gin Shop",
		"isLogin":isLogin,
	})
}
