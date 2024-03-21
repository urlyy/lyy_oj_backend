package route

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func configCors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Header("Access-Control-Allow-Headers", "*")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")
		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	}
}

var r *gin.Engine

const (
	PAGE_SIZE = 25
)

func init() {
	r = gin.Default()
	r.Use(configCors(), jwtAuthMiddleware())
	addUserRoute(r)
	addDomainRoute(r)
	addProblemRoute(r)
	addHomeworkRoute(r)
	addContestRoute(r)
	addDiscussionRoute(r)
}

func GetRouter() *gin.Engine {
	return r
}
