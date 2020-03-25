/*
 *    Copyright 2020 opensourceai
 *
 *    Licensed under the Apache License, Version 2.0 (the "License");
 *    you may not use this file except in compliance with the License.
 *    You may obtain a copy of the License at
 *
 *        http://www.apache.org/licenses/LICENSE-2.0
 *
 *    Unless required by applicable law or agreed to in writing, software
 *    distributed under the License is distributed on an "AS IS" BASIS,
 *    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *    See the License for the specific language governing permissions and
 *    limitations under the License.
 */

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
