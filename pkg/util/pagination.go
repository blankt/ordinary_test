package util

import (
	"github.com/gin-gonic/gin"
	"ordinary_test/pkg/setting"
	"strconv"
)

func GetPage(c *gin.Context) int {
	result := 0
	page, _ := strconv.Atoi(c.Query("page"))
	if page > 0 {
		result = (page - 1) * setting.DatabaseSetting.PageSize
	}
	return result
}
