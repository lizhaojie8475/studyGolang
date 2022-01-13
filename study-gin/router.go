package main

import "github.com/gin-gonic/gin"

func register(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	}).Use()
}
