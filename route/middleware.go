package route

import (
	"backend/util"

	"github.com/gin-gonic/gin"
)

var (
	exclude = [4]string{
		"/user/login",
		"/user/register",
		"/user/forget-pass",
		"/user/forget-pass/captcha",
	}
)

func jwtAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		for _, e := range exclude {
			if c.Request.RequestURI == e {
				c.Next()
				return
			}
		}
		//获取到请求头中的token
		inputToken := c.Request.Header.Get("Authorization")
		if inputToken == "" {
			NewResult(c).Fail("请登录!")
		} else {
			claims := util.ParseToken(inputToken)
			if claims == nil {
				NewResult(c).Fail("请重新登录!")
			} else {
				c.Set("userID", claims.UserID)
				// 后续的处理函数可以用过c.Get("username")来获取当前请求的用户信息
				c.Next()
			}
		}
	}
}

func addMiddleware(r *gin.Engine) {
	r.Use(configCors(), jwtAuthMiddleware())
}
