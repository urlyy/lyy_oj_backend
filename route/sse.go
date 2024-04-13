package route

import (
	"backend/util"
	"encoding/json"
	"time"

	"github.com/gin-gonic/gin"
)

func connect(c *gin.Context) {
	roomID := c.Param("id")
	util.SSEConnect(c, roomID)
	NewResult(c).Success("连接成功", nil)
}

func sendNotification(c *gin.Context) {
	roomID := c.Param("id")
	data := c.PostForm("msg")
	// println("Data received from client :", data)
	msg := map[string]interface{}{
		"text":     data,
		"sendTime": time.Now(),
	}
	str, err := json.Marshal(msg)
	if err != nil {
		NewResult(c).Fail("服务端发送通知异常")
		return
	}
	util.SSEBroadcast(roomID, string(str))
	NewResult(c).Success("发送成功", nil)
}

func addSSERoute(r *gin.Engine) {
	api := r.Group("/sse/:id")
	api.POST("/notify", sendNotification)
	api.GET("", connect)
}
