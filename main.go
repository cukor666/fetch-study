package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// CORSMiddleware 跨域中间件
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}

func main() {
	engine := gin.Default()

	// 使用跨域中间件
	engine.Use(CORSMiddleware())

	engine.GET("/", func(c *gin.Context) {
		log.Println("请求被处理")
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "success",
			"data": "Hello World",
		})
	})

	engine.GET("/p", func(c *gin.Context) {
		log.Println("请求被处理")
		// 接收请求参数
		var data struct {
			Name  string   `form:"name" binding:"required"`
			Age   int      `form:"age" binding:"required"`
			Hobby []string `form:"hobby" binding:"required"`
		}
		if err := c.ShouldBindQuery(&data); err != nil {
			c.JSON(400, gin.H{
				"code": 400,
				"msg":  "参数错误",
			})
			return
		}
		log.Println(data)
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "success",
			"data": data,
		})
	})

	engine.POST("/", func(c *gin.Context) {
		log.Println("请求被处理")
		// 接收请求参数
		var data struct {
			Name  string   `json:"name" binding:"required"`
			Age   int      `json:"age" binding:"required"`
			Hobby []string `json:"hobby" binding:"required"`
		}
		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(400, gin.H{
				"code": 400,
				"msg":  "参数错误",
				"data": err.Error(),
			})
			log.Println("请求失败")
			return
		}
		log.Println(data)
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "success",
			"data": data,
		})
	})

	engine.PUT("/", func(c *gin.Context) {
		log.Println("请求被处理")
		// 接收请求参数
		var data struct {
			Name  string   `json:"name" binding:"required"`
			Age   int      `json:"age" binding:"required"`
			Hobby []string `json:"hobby" binding:"required"`
		}
		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(400, gin.H{
				"code": 400,
				"msg":  "参数错误",
				"data": err.Error(),
			})
			log.Println("请求失败")
			return
		}
		log.Println(data)
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "success",
			"data": data,
		})
	})

	engine.DELETE("/", func(c *gin.Context) {
		log.Println("请求被处理")
		// 接收请求参数
		var data struct {
			ID int `form:"id" binding:"required"`
		}
		if err := c.ShouldBindQuery(&data); err != nil {
			c.JSON(400, gin.H{
				"code": 400,
				"msg":  "参数错误",
			})
			return
		}
		log.Println(data)
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "success",
			"data": data,
		})
	})

	if err := engine.Run(":8080"); err != nil {
		panic(err)
	}
}
