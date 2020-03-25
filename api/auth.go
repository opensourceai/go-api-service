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

package api

import (
	"fmt"
	"github.com/google/wire"
	"github.com/opensourceai/go-api-service/internal/models"
	"github.com/opensourceai/go-api-service/internal/service"
	"github.com/opensourceai/go-api-service/middleware/jwt"
	"github.com/opensourceai/go-api-service/pkg/logging"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/opensourceai/go-api-service/pkg/app"
	"github.com/opensourceai/go-api-service/pkg/e"
	"github.com/opensourceai/go-api-service/pkg/util"
)

type auth struct {
	Username string `valid:"Required; MaxSize(50)"` // 用户名
	Password string `valid:"Required; MaxSize(50)"` // 密码
}

type AuthApi struct {
}

var ProviderAuth = wire.NewSet(NewAuthApi)

func NewAuthApi(service2 service.UserService) (*AuthApi, error) {
	userService = service2
	return &AuthApi{}, nil
}

var userService service.UserService

//
func NewAuthRouter(router *gin.Engine) {
	auth := router.Group("/auth")
	{
		auth.POST("/login", login)
		auth.POST("/register", register)
	}
	group := router.Group("/auth")
	group.Use(jwt.JWT())
	{
		group.GET("/test", authTest)

	}

}

// @Summary 获取认证信息
// @Tags Auth
// @Produce  json
// @Param user body auth true "user"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /auth/login [post]
func login(c *gin.Context) {
	user := auth{}
	appG := app.Gin{C: c}
	httpCode, errCode := app.BindAndValid(c, &user)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	userFound, isExist, err := userService.Login(models.User{Username: user.Username, Password: user.Password})
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_AUTH_CHECK_TOKEN_FAIL, nil)
		return
	}

	if !isExist {
		appG.Response(http.StatusUnauthorized, e.ERROR_AUTH, nil)
		return
	}
	// 生成token
	token, err := util.GenerateToken(userFound.ID, userFound.Username)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_AUTH_TOKEN, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, map[string]string{
		"token": token,
	})
}

// @Summary 添加用户
// @Tags Auth
// @Produce  json
// @Param user body models.User true "user"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /auth/register [post]
func register(c *gin.Context) {
	user := models.User{}
	appG := app.Gin{C: c}
	httpCode, errCode := app.BindAndValid(c, &user)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}
	if err := userService.Register(&user); err != nil {
		logging.Error(err)
		appG.Response(http.StatusBadRequest, e.ERROR, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

// @Summary auth测试
// @Tags Auth
// @Produce  json
// @Param str query string true "string"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Security ApiKeyAuth
// @Router /auth/test [get]
func authTest(c *gin.Context) {
	query := c.Query("str")
	get, _ := c.Get("username")
	fmt.Println("get", get)
	str := fmt.Sprintf("%s", get)
	g := app.Gin{C: c}
	g.Response(http.StatusOK, e.SUCCESS, str+query)

}
