package web

import "github.com/gin-gonic/gin"

func CreateCourse(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "create course",
	})
}

func EditCourse(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "edit course",
	})
}

func GetCourse(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "get course",
	})
}

func DeleteCourse(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "delete course",
	})
}
