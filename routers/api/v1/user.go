package v1

import (
	"github.com/gin-gonic/gin"
)

func UserApi(r *gin.Engine) {

	r.Group("/v1/user")
	{
		r.POST("/register", register)
	}
}

// @Summary 添加用户
// @Produce  json
// @Param user body models.User true "user"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/user [post]
func register(c *gin.Context) {

}
