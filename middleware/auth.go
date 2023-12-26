package middleware

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthCheck(c *gin.Context) {
	userId, _ := c.Get("user_id")
	userName, _ := c.Get("user_name")
	fmt.Printf("userID:%s,userName:%s", userId, userName)
	c.Next()
}

var token = "123456"

func TokenCheck(c *gin.Context) {
	accessToken := c.Request.Header.Get("access_token")
	if accessToken != token {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "token error",
		})
		c.AbortWithError(http.StatusInternalServerError, errors.New("token 检查失败"))
	}
	c.Set("user_name", "gjx")
	c.Set("user_id", "20419000508")
}
