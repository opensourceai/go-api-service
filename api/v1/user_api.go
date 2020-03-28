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
	"github.com/opensourceai/go-api-service/api/v1/dto"
	"github.com/opensourceai/go-api-service/internal/service"
	"github.com/opensourceai/go-api-service/middleware/jwt"
	"github.com/opensourceai/go-api-service/pkg/app"
	"github.com/opensourceai/go-api-service/pkg/e"
)

// 用户API依赖注入结构体
type UserApi struct {
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

// 路由
func NewUserRouter(router *gin.Engine) {
	user := router.Group("/v1/user")
	user.Use(jwt.JWT())
	{
		user.PUT("", updateUser)
	}
}

// @Summary 用户密码修改或用户信息修改
// @Tags User
// @Produce  json
// @Param user body dto.UserDTO true "user"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Security ApiKeyAuth
// @Router /v1/user [put]
func updateUser(context *gin.Context) {
	appG := app.Gin{C: context}
	// 参数绑定
	userDTO := &dto.UserDTO{}
	httpCode, errCode := app.BindAndValid(context, userDTO)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}
	onlineUser := app.GetUserInfo(context)

	if err := userService.ServiceUpdateUser(onlineUser, userDTO); err != nil {
		appG.Fail(nil)
		return
	}

	appG.Success(nil)
}

