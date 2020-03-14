package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/opensourceai/go-api-service/internal/service"
	"github.com/opensourceai/go-api-service/pkg/app"
	"github.com/opensourceai/go-api-service/pkg/e"
	"github.com/opensourceai/go-api-service/pkg/page"
	"net/http"
	"strconv"
)

var boardService service.BoardService

func init() {
	boardService = new(service.BoardServiceImpl)
}
func BroadApi(router *gin.Engine) {
	broad := router.Group("/v1/broad")
	{
		broad.GET("", GetBroadList)
		broad.GET("/:id", GetBroad)
		broad.GET("/:id/posts", GetPostListInBroad)
	}
}

// @Summary 获取全部版块列表信息
// @Tags Broad
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /v1/broad [get]
func GetBroadList(context *gin.Context) {
	appG := app.Gin{C: context}
	if boards, err := boardService.GetBoardList(); err == nil {
		appG.Success(boards)
	} else {
		appG.Fail(nil)
	}
}

// @Summary 获取某版块信息
// @Tags Broad
// @Produce  json
// @Param id path string true "id"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /v1/broad/{id} [get]
func GetBroad(context *gin.Context) {
	appG := app.Gin{C: context}
	id := context.Param("id")
	if id == "" {
		appG.Fail(nil)
		return
	}
	if idInt, err := strconv.Atoi(id); err != nil {
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
	} else {
		// 获取数据
		if board, err := boardService.GetBoard(idInt); err == nil {
			appG.Success(board)
		} else {
			appG.Fail(nil)
		}
	}

}

// @Summary 获取某版块的帖子
// @Tags Broad
// @Produce  json
// @Param id path string true "id"
// @Param page query page.Page true "page"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /v1/broad/{id}/posts [get]
func GetPostListInBroad(context *gin.Context) {
	appG := app.Gin{C: context}

	// 板块ID
	id := context.Param("id")
	if id == "" {
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
	}
	// 转整型
	if idInt, err := strconv.Atoi(id); err != nil {
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
	} else {
		// 绑定分页参数
		bindPage := page.BindPage(context)
		// 获取数据
		if list, err := boardService.GetPostList(idInt, bindPage); err != nil {
			appG.Response(http.StatusBadRequest, e.ERROR_POST_NOT_EXIST, nil)
		} else {
			appG.Success(list)
		}
	}

}
