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

package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/opensourceai/go-api-service/internal/models"
	"github.com/opensourceai/go-api-service/internal/service"
	"github.com/opensourceai/go-api-service/pkg/app"
	"github.com/opensourceai/go-api-service/pkg/e"
	"github.com/opensourceai/go-api-service/pkg/logging"
	"net/http"
)

type UserApi struct {
}

type Msg struct {
}

//1.等待注入的业务
var userService service.UserService

//2.为userService注入依赖
func NewUserApi(service2 service.UserService) (*UserApi, error) {
	userService = service2
	return &UserApi{}, nil
}

//3.使用wire为NewUserApi注入依赖
var ProviderUser = wire.NewSet(NewUserApi, service.ProviderUser)

func NewUserRouter(router *gin.Engine) {
	user := router.Group("/v1/user")
	{
		user.PUT("/update-pwd", updatePwd)
		user.PUT("/update-message", updateMsg)
	}
}

// @Summary 用户密码修改
// @Tags User
// @Produce  json
// @Param user body models.User true "user"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /v1/user/update-pwd [put]
func updatePwd(c *gin.Context) {
	user := models.User{}
	//var newPwd string
	//user.Password = newPwd
	appG := app.Gin{C: c}
	httpCode, errCode := app.BindAndValid(c, &user)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}
	if err := userService.UpdatePwd(user.Username, user.Password); err != nil {
		logging.Error(err)
		appG.Response(http.StatusBadRequest, e.ERROR, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, nil)

}

// @Summary 用户信息修改
// @Tags User
// @Produce  json
// @Param user body models.User true "user"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /v1/user/update-message [put]
func updateMsg(c *gin.Context) {
	user := models.User{}

	appG := app.Gin{C: c}
	//页面内容绑定到user
	httpCode, errCode := app.BindAndValid(c, &user)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}
	if err := userService.UpdateMsg(user.Username, &user); err != nil {
		logging.Error(err)
		appG.Response(http.StatusBadRequest, e.ERROR, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, nil)

}
