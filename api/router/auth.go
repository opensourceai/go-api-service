package router

import (
	"github.com/opensourceai/go-api-service/models"
	"github.com/opensourceai/go-api-service/pkg/logging"
	"github.com/opensourceai/go-api-service/service"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/opensourceai/go-api-service/pkg/app"
	"github.com/opensourceai/go-api-service/pkg/e"
	"github.com/opensourceai/go-api-service/pkg/util"
)

type auth struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

var userService service.UserService

func init() {
	userService = new(service.UserServiceImpl)
}

func Auth(router *gin.Engine) {
	auth := router.Group("/auth")
	{
		auth.POST("/login", login)
		auth.POST("/register", register)
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

	isExist, err := userService.Login(models.User{Username: user.Username, Password: user.Password})
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_AUTH_CHECK_TOKEN_FAIL, nil)
		return
	}

	if !isExist {
		appG.Response(http.StatusUnauthorized, e.ERROR_AUTH, nil)
		return
	}

	token, err := util.GenerateToken(user.Username, user.Password)
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
