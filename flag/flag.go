package flag

import (
	"github.com/gin-gonic/gin"
)

// FlagIsNilUseQuery 用于webserver接口数据获取,可以在启动时通过flag获取参数，也可以通过解析uri参数query获取
func FlagIsNilUseQuery(flagParam, GinParam string, c *gin.Context) string {
	if flagParam == "" {
		return c.Query(GinParam)
	}
	return flagParam
}
