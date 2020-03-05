package app

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"net/http"

	"github.com/opensourceai/go-api-service/pkg/e"
)

// BindAndValid binds and validates data
func BindAndValid(c *gin.Context, form interface{}) (int, int) {
	err := c.Bind(form)
	if err != nil {
		return http.StatusBadRequest, e.INVALID_PARAMS
	}

	valid := validation.Validation{}
	check, err := valid.Valid(form)
	if err != nil {
		return http.StatusInternalServerError, e.ERROR
	}
	if !check {
		MarkErrors(valid.Errors)
		return http.StatusBadRequest, e.INVALID_PARAMS
	}

	return http.StatusOK, e.SUCCESS
}
func Valid(form interface{}) (int, int) {

	valid := validation.Validation{}
	check, err := valid.Valid(form)
	if err != nil {
		return http.StatusInternalServerError, e.ERROR
	}
	if !check {
		MarkErrors(valid.Errors)
		return http.StatusBadRequest, e.INVALID_PARAMS
	}
	return http.StatusOK, e.SUCCESS

}

type Auth struct {
	UserId   string `json:"user_id"`
	Username string `json:"username"`
}

func GetUserInfo(content *gin.Context) *Auth {
	var userId interface{}
	var username interface{}
	var exists bool
	if userId, exists = content.Get("userId"); !exists {
		panic("认证失败")
	}

	if username, exists = content.Get("username"); !exists {
		panic("认证失败")
	}
	return &Auth{
		UserId:   com.ToStr(userId),
		Username: com.ToStr(username),
	}

}
