package util

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// jwt包自带的jwt.RegisteredClaims只包含了官方字段
// 更多信息都可以添加到这个结构体中
type CustomClaims struct {
	// 可根据需要自行添加字段
	UserID               int       `json:"user_id"`
	LoginTime            time.Time `json:"login_time"`
	jwt.RegisteredClaims           // 内嵌标准的声明
}

// GenToken 生成JWT
func GenToken(userId int, loginTime time.Time) (string, error) {
	// 创建一个我们自己声明的数据
	claims := CustomClaims{
		userId,
		loginTime,
		jwt.RegisteredClaims{
			// 定义过期时间
			ExpiresAt: jwt.NewNumericDate(loginTime.Add(time.Hour * time.Duration(GetProjectConfig().JWT.Expire))),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 生成签名字符串
	return token.SignedString([]byte(GetProjectConfig().JWT.Secret))
}

func ParseToken(tokenString string) *CustomClaims {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(GetProjectConfig().JWT.Secret), nil
	})
	if err != nil {
		// log.Fatal(err)
		return nil
	} else if claims, ok := token.Claims.(*CustomClaims); ok {
		return claims
	} else {
		return nil
	}
}
