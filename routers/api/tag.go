package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ordinary_test/pkg/app"
	"ordinary_test/pkg/e"
)

// GetTags @Summary Get multiple article tags
// @Produce  json
// @Param name query string false "Name"
// @Param state query int false "State"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/tags [get]
func GetTags(c *gin.Context) {
	appG := app.Gin{C: c}

	appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{
		"lists": "成功获取tag",
		"total": 1,
	})
}
