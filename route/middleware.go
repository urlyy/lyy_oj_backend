package route

import (
	"backend/util"
	"fmt"
	"regexp"
	"strconv"

	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	exclude = []string{
		"/user/login",
		"/user/register",
		"/user/forget-pass",
		"/user/forget-pass/captcha",
	}
)

func match(pattern string, str string) bool {
	// 编译正则表达式
	regex, err := regexp.Compile(pattern)
	if err != nil {
		fmt.Println("Error compiling regex:", err)
		return false
	}
	// 使用正则表达式进行匹配
	return regex.MatchString(str)
}

func jwtAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		for _, e := range exclude {
			if match(e, c.Request.RequestURI) {
				c.Next()
				return
			}
		}
		//获取到请求头中的token
		inputToken := c.Request.Header.Get(util.GetProjectConfig().JWT.Key)
		if inputToken == "" {
			NewResult(c).Fail("请登录!")
			return
		}
		claims := util.ParseToken(inputToken)
		if claims == nil {
			NewResult(c).Fail("请重新登录!")
			return
		}
		rdsLoginTime, err := util.RedisGet(util.RedisTokenKey(claims.UserID))
		//单设备登录
		if err != nil || rdsLoginTime != strconv.FormatInt(claims.LoginTime, 10) {
			NewResult(c).Fail("账号已在其他地方登录，当前无法进行操作")
			return
		}
		// 后续的处理函数可以用过c.Get("username")来获取当前请求的用户信息
		c.Set("userID", claims.UserID)
		c.Next()
	}
}

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

func addMiddleware(r *gin.Engine) {
	r.Use(configCors(), jwtAuthMiddleware())
}
