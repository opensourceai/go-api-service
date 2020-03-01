package api

import (
	v1 "github.com/opensourceai/go-api-service/api/router/v1"
	"github.com/opensourceai/go-api-service/middleware/jwt"
	"github.com/opensourceai/go-api-service/pkg/setting"
	"net/http"

	"github.com/gin-gonic/gin"
	// swagger
	_ "github.com/opensourceai/go-api-service/docs"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	"github.com/opensourceai/go-api-service/api/router"
	"github.com/opensourceai/go-api-service/pkg/export"
	"github.com/opensourceai/go-api-service/pkg/qrcode"
	"github.com/opensourceai/go-api-service/pkg/upload"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.StaticFS("/export", http.Dir(export.GetExcelFullPath()))
	r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))
	r.StaticFS("/qrcode", http.Dir(qrcode.GetQrCodeFullPath()))

	//r.GET("/auth", router.GetAuth)
	// swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.POST("/upload", router.UploadImage)

	// 认证
	router.Auth(r)

	//r.Use(jwt.JWT())
	if setting.ServerSetting.RunMode == "prod" {
		// 添加全局token认证中间件
		r.Use(jwt.JWT())
	}

	// 用户
	v1.UserApi(r)

	//apiv1 := r.Group("/router/v1")
	//apiv1.Use(jwt.JWT())
	//{
	//
	//}

	return r
}
