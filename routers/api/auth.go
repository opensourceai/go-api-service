package api

import (
	"github.com/opensourceai/go-api-service/service"
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"

	"github.com/opensourceai/go-api-service/pkg/app"
	"github.com/opensourceai/go-api-service/pkg/e"
	"github.com/opensourceai/go-api-service/pkg/util"
)

type auth struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

func GetAuth(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}

	username := c.Query("username")
	password := c.Query("password")

	a := auth{Username: username, Password: password}
	ok, _ := valid.Valid(&a)

	if !ok {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	authService := service.Auth{Username: username, Password: password}
	isExist, err := authService.Check()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_AUTH_CHECK_TOKEN_FAIL, nil)
		return
	}

	if !isExist {
		appG.Response(http.StatusUnauthorized, e.ERROR_AUTH, nil)
		return
	}

	token, err := util.GenerateToken(username, password)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_AUTH_TOKEN, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, map[string]string{
		"token": token,
	})
}

// @Summary 获取认证信息
// @Produce  json
// @Param username query string true "userName"
// @Param password query string true "password"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /auth [get]
func Auth(router *gin.Engine) {
	router.GET("/auth", func(c *gin.Context) {
		appG := app.Gin{C: c}
		valid := validation.Validation{}

		username := c.Query("username")
		password := c.Query("password")

		a := auth{Username: username, Password: password}
		ok, _ := valid.Valid(&a)

		if !ok {
			app.MarkErrors(valid.Errors)
			appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
			return
		}

		authService := service.Auth{Username: username, Password: password}
		isExist, err := authService.Check()
		if err != nil {
			appG.Response(http.StatusInternalServerError, e.ERROR_AUTH_CHECK_TOKEN_FAIL, nil)
			return
		}

		if !isExist {
			appG.Response(http.StatusUnauthorized, e.ERROR_AUTH, nil)
			return
		}

		token, err := util.GenerateToken(username, password)
		if err != nil {
			appG.Response(http.StatusInternalServerError, e.ERROR_AUTH_TOKEN, nil)
			return
		}

		appG.Response(http.StatusOK, e.SUCCESS, map[string]string{
			"token": token,
		})
	})
}
