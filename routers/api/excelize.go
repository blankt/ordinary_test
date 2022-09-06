package api

import (
	"github.com/gin-gonic/gin"
	"ordinary_test/pkg/app"
)

func GetArticleList(c *gin.Context) {
	appG := app.Gin{C: c}

}
