package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/opensourceai/go-api-service/models"
	"github.com/opensourceai/go-api-service/pkg/app"
	"github.com/opensourceai/go-api-service/pkg/e"
	"github.com/opensourceai/go-api-service/pkg/logging"
	"github.com/opensourceai/go-api-service/service"
	"net/http"
)

func UserApi(r *gin.Engine) {

	v1 := r.Group("/v1/user")
	{
		v1.POST("/register", register)
	}
}

var userService service.UserService

func init() {
	userService = new(service.UserServiceImpl)
}

// @Summary 添加用户
// @Produce  json
// @Param user body models.User true "user"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /v1/user/register [post]
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
