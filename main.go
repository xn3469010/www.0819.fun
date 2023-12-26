package main

import (
	"github.com/gin-gonic/gin"
	routers "www.0819.fun/routes"
)

func main() {
	r := gin.Default()
	routers.InitRouters(r)
	r.Run(":8080")
}
