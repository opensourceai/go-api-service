package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
	"github.com/opensourceai/go-api-service/internal/models"
	"github.com/opensourceai/go-api-service/internal/service"
	"github.com/opensourceai/go-api-service/middleware/jwt"
	"github.com/opensourceai/go-api-service/pkg/app"
	"github.com/opensourceai/go-api-service/pkg/e"
	"github.com/opensourceai/go-api-service/pkg/page"
	"net/http"
	"strconv"
)

type PostApi struct {
}

var ProviderPost = wire.NewSet(NewPostApi, service.ProviderPost)

var postService service.PostService

func NewPostApi(service2 service.PostService) (*PostApi, error) {
	postService = service2
	return &PostApi{}, nil
}

func NewPostRouter(router *gin.Engine) {
	post := router.Group("/v1/post")
	// 无需认证
	{
		post.GET("/:id", getPost)
	}
	// 需认证
	post.Use(jwt.JWT())
	{
		post.POST("", addPost)
		post.DELETE("", deletePost)
		post.GET("", getPostList)
		post.PUT("", updatePost)
	}

}

// @Summary 获取帖子信息
// @Tags Post
// @Produce  json
// @Param id path string true "id"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Security ApiKeyAuth
// @Router /v1/post/{id} [get]
func getPost(context *gin.Context) {
	appG := app.Gin{C: context}
	// 请求异常处理
	defer app.Recover(&appG)

	id := context.Param("id")
	if id == "" {
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}
	//userInfo := app.GetUserInfo(context)
	post, err := postService.GetPost(id)
	if err != nil {
		appG.Response(http.StatusNotFound, e.ERROR_POST_NOT_EXIST, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, post)

}

// @Summary 修改用户自身帖子
// @Tags Post
// @Produce  json
// @Param post body models.Post true "post"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Security ApiKeyAuth
// @Router /v1/post [put]
func updatePost(context *gin.Context) {
	appG := app.Gin{C: context}
	// 请求异常处理
	//defer app.Recover(&appG)
	post := models.Post{}
	httpCode, errCode := app.BindAndValid(context, &post)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}
	userInfo := app.GetUserInfo(context)
	if err := postService.UpdatePost(userInfo.UserId, &post); err == gorm.ErrRecordNotFound {
		appG.Response(http.StatusNotFound, e.NOT_FOUND, nil)
		return
	} else if err != nil {
		appG.Response(http.StatusBadRequest, e.ERROR, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

// @Summary 获取用户自身帖子
// @Tags Post
// @Produce  json
// @Param page query string true "pageNum"
// @Param size query string true "pageSize"
// @Param orderBy query string false "orderBy"
// @Param sorter query string false "sorter"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Security ApiKeyAuth
// @Router /v1/post [get]
func getPostList(context *gin.Context) {
	appG := app.Gin{C: context}
	// 请求异常处理
	defer app.Recover(&appG)
	var err error
	// 第几页
	pageNum := context.Query("page")
	var pageNumInt int

	if pageNum == "" {
		panic("pageNum 不能为空.")
	} else {
		if pageNumInt, err = strconv.Atoi(pageNum); err != nil {
			panic("pageNum 无法转换.")
		}
	}
	// 每页数据个数
	pageSize := context.Query("size")
	var pageSizeInt int

	if pageSize == "" {
		panic("pageSize 不能为空.")

	} else {
		if pageSizeInt, err = strconv.Atoi(pageSize); err != nil {
			panic("pageSize 无法转换.")
		}
	}

	// 排序字段
	sorter := context.Query("sorter")
	if sorter == "" {
		sorter = "asc"
	}

	// 排序方式
	orderBy := context.Query("orderBy")
	if orderBy == "" {
		orderBy = "id"
	}

	pageObj := page.NewPage(pageNumInt, pageSizeInt, orderBy, sorter)

	// 参数检验
	httpCode, errCode := app.Valid(pageObj)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}
	info := app.GetUserInfo(context)

	if post, err := postService.GetOwnPost(pageObj, info.UserId); err != nil {
		appG.Response(http.StatusBadRequest, e.ERROR, post)
		return
	} else {
		appG.Response(http.StatusOK, e.SUCCESS, post)
		return
	}

}

type postIds struct {
	Ids []int `json:"ids"` // 帖子IDs
}

// @Summary 删除用户帖子
// @Tags Post
// @Produce  json
// @Param postIds body postIds true "postIds"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Security ApiKeyAuth
// @Router /v1/post [delete]
func deletePost(context *gin.Context) {
	postIds := postIds{}
	appG := app.Gin{C: context}
	defer app.Recover(&appG)

	httpCode, errCode := app.BindAndValid(context, &postIds)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}
	//token := context.GetHeader("Authorization")
	//userId, exists := context.Get("userId")
	userInfo := app.GetUserInfo(context)
	if err := postService.DeletePost(userInfo.UserId, postIds.Ids...); err != nil {
		if err == gorm.ErrRecordNotFound {
			appG.Response(http.StatusBadRequest, e.ERROR_POST_NOT_EXIST, nil)
			return
		}
		appG.Response(http.StatusBadRequest, e.ERROR, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, nil)

}

// @Summary 新增帖子
// @Tags Post
// @Produce  json
// @Param post body models.Post true "post"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Security ApiKeyAuth
// @Router /v1/post [post]
func addPost(context *gin.Context) {
	post := models.Post{}
	appG := app.Gin{C: context}
	defer app.Recover(&appG)

	// 参数校验
	httpCode, errCode := app.BindAndValid(context, &post)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}
	if err := postService.AddPost(&post); err != nil {
		appG.Response(http.StatusBadRequest, e.ERROR, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, nil)

}
