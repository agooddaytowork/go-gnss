package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/umeat/go-gnss/cmd/database/apis"
)

func main() {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	v1 := r.Group("/api/v1")
	v1.GET("/observation/:id", apis.GetObservation)

	r.Run(fmt.Sprintf(":%v", 8000))
}
