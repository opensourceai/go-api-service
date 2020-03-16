package api

import (
	"net/http"

	v1 "github.com/opensourceai/go-api-service/api/v1"

	"github.com/gin-gonic/gin"
	// swagger
	_ "github.com/opensourceai/go-api-service/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/opensourceai/go-api-service/pkg/export"
	"github.com/opensourceai/go-api-service/pkg/qrcode"
	"github.com/opensourceai/go-api-service/pkg/upload"
)

// InitRouter initialize routing information
func NewApi() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.StaticFS("/export", http.Dir(export.GetExcelFullPath()))
	r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))
	r.StaticFS("/qrcode", http.Dir(qrcode.GetQrCodeFullPath()))
	// swagger
	//url := ginSwagger.URL("http://0.0.0.0:8000/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Oauth认证
	OauthApi(r)
	// 普通认证
	NewAuthRouter(r)
	// 版块
	v1.NewBoardRouter(r)
	// 主题
	v1.NewPostRouter(r)

	return r
}
