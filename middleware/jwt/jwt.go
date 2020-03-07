package jwt

import (
	"github.com/opensourceai/go-api-service/pkg/app"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	"github.com/opensourceai/go-api-service/pkg/e"
	"github.com/opensourceai/go-api-service/pkg/util"
)

// JWT is jwt middleware
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}
		var claims *util.Claims
		var err error
		code = e.SUCCESS
		// 获取头信息中的token
		token := c.GetHeader("Authorization")
		if token == "" {
			code = e.ERROR_AUTH_NOT_FOUND_TOKEN
		} else {
			claims, err = util.ParseToken(token)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
				default:
					code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
				}
			}

		}

		if code != e.SUCCESS {
			g := app.Gin{C: c}
			g.Response(http.StatusUnauthorized, code, data)
			c.Abort()
			return
		}
		if claims != nil {
			c.Set("username", claims.Username)
			c.Set("userId", claims.Id)
		}

		c.Next()
	}
}
