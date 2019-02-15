package main

import (
	_"fmt"
	"github.com/gin-gonic/gin"
	route "./route"
)

func main() {
	const serverHost string = "0.0.0.0:8080"
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	
	r.GET("/", route.Index)
	r.GET("/randQuestion", route.RandQuestion)
	r.Run(serverHost)
}