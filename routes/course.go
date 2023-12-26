package routers

import (
	web "www.0819.fun/app/Http/Controllers/v1"
	"www.0819.fun/middleware"

	"github.com/gin-gonic/gin"
)

func initCourse(r *gin.Engine) {
	//http://localhost/v1/course
	v1 := r.Group("/v1", middleware.TokenCheck, middleware.AuthCheck)
	v1.POST("/course", web.CreateCourse)
	v1.GET("/course", web.GetCourse)
	v1.PUT("/course", web.EditCourse)      //修改课程信息
	v1.DELETE("/course", web.DeleteCourse) //删除课程信息
}
