package routers

import (
	"github.com/gin-gonic/gin"
	"ordinary_test/routers/api"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	apiv1 := r.Group("api/v1")
	{
		apiv1.GET("/tags", api.GetTags)
	}
	return r
}
