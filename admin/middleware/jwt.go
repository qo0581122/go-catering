package middleware

import (
	"fmt"
	"strconv"
	"time"

	"catering/pkg"

	"catering/global"
	"catering/model/common/response"

	"github.com/gin-gonic/gin"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 我们这里jwt鉴权取头部信息 x-token 登录时回返回token信息
		token := c.Request.Header.Get("x-token")
		if token == "" {
			response.FailWithDetailed(gin.H{"reload": true}, "未登录或非法访问", c)
			c.Abort()
			return
		}

		// 获取session
		session, _ := global.SESSION.Get(c.Request, "SessionId")
		t := session.Values["token"]

		//判断session和对session中的token进行比对
		fmt.Println(t, token)
		// if tk, ok := t.(string); !ok {
		// 	response.FailWithDetailed(gin.H{"reload": true}, "授权已过期", c)
		// 	c.Abort()
		// 	return
		// } else if tk != token {
		// 	response.FailWithDetailed(gin.H{"reload": true}, "授权已过期", c)
		// 	c.Abort()
		// 	return
		// }

		j := pkg.NewJWT()
		// parseToken 解析token包含的信息
		// 判断该token是否已经到达有效期
		claims, err := j.ParseToken(token)
		if err != nil {
			if err == pkg.ErrTokenExpired {
				response.FailWithDetailed(gin.H{"reload": true}, "授权已过期", c)
				c.Abort()
				return
			}
			response.FailWithDetailed(gin.H{"reload": true}, err.Error(), c)
			c.Abort()
			return
		}
		if claims.ExpiresAt-time.Now().Unix() < claims.BufferTime {
			claims.ExpiresAt = time.Now().Unix() + global.Config.JWT.ExpiresTime
			newToken, _ := j.CreateTokenByOldToken(token, *claims)
			newClaims, _ := j.ParseToken(newToken)
			c.Header("new-token", newToken)
			c.Header("new-expires-at", strconv.FormatInt(newClaims.ExpiresAt, 10))
		}
		c.Set("claims", claims)
		c.Next()
	}
}
