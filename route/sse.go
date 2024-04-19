package route

import (
	"backend/util"

	"github.com/gin-gonic/gin"
)

func connect(c *gin.Context) {
	roomID := c.Param("id")
	util.SSEConnect(c, roomID)
	NewResult(c).Success("连接成功", nil)
}

func addSSERoute(r *gin.Engine) {
	api := r.Group("/sse/:id")
	api.GET("", connect)
}
