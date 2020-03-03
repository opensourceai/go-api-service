package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/opensourceai/go-api-service/middleware/jwt"
	"github.com/opensourceai/go-api-service/models"
	"github.com/opensourceai/go-api-service/pkg/app"
	"github.com/opensourceai/go-api-service/pkg/e"
	"github.com/opensourceai/go-api-service/pkg/page"
	"github.com/opensourceai/go-api-service/service"
	"net/http"
)

var postService service.PostService

func init() {
	postService = new(service.PostServiceImpl)
}
func PostApi(router *gin.Engine) {
	post := router.Group("/v1/post")
	post.Use(jwt.JWT())
	{
		post.POST("", addPost)
		post.DELETE("", deletePost)
		post.GET("", getPost)
	}
}

type postPage struct {
	models.Post
	page.Page
}

func getPost(context *gin.Context) {

}

type postIds struct {
	Ids []uint
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
	httpCode, errCode := app.BindAndValid(context, &postIds)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}
	//token := context.GetHeader("Authorization")

	if err := postService.DeletePost(postIds.Ids...); err != nil {
		appG.Response(http.StatusBadRequest, e.ERROR, nil)
		return
	}

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
