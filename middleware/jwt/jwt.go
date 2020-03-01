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

		code = e.SUCCESS
		token := c.Query("token")
		if token == "" {
			code = e.INVALID_PARAMS
		} else {
			_, err := util.ParseToken(token)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
				default:
					code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
				}
			}
		}

		//c.JSON(http.StatusUnauthorized, gin.H{
		//	"code": code,
		//	"msg":  e.GetMsg(code),
		//	"data": data,
		//}

		if code != e.SUCCESS {
			g := app.Gin{c}
			g.Response(http.StatusUnauthorized, code, data)
			c.Abort()
			return
		}

		c.Next()
	}
}
