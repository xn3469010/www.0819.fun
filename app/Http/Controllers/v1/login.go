package web

import "github.com/gin-gonic/gin"

type loginReq struct {
	UserName string `form:"user_name"`
	Pwd      string `form:"pwd"`
}

func Login(c *gin.Context) {
	req := &loginReq{}
	c.Bind(&req)
	c.JSON(200, req)
}
