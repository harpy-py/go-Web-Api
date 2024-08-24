package api

import (
	"github.com/gin-gonic/gin"
	"github.com/harpy-py/go-Web-Api/api/routers"
)

func Initserver()  {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())
	
	v1 := r.Group("/api/v1/")
	{
		health := v1.Group("/health")
		routers.Health(health)
	}
	r.Run(":2002")
}