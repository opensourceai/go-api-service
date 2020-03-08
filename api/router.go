package api

import (
	v1 "github.com/opensourceai/go-api-service/api/router/v1"
	"net/http"

	"github.com/gin-gonic/gin"
	// swagger
	_ "github.com/opensourceai/go-api-service/docs"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"

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
	// swagger
	//url := ginSwagger.URL("http://0.0.0.0:8000/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 认证
	router.Auth(r)
	router.Oauth2(r)

	//r.Use(jwt.JWT())
	//if setting.ServerSetting.RunMode == "prod" {
	//	// 添加全局token认证中间件
	//	r.Use(jwt.JWT())
	//}

	// 用户
	v1.UserApi(r)
	// 帖子
	v1.PostApi(r)

	return r
}
