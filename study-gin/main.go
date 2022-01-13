package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	defer r.Run()

	register(r)
}
