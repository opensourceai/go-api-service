package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/opensourceai/go-api-service/dao/mysql"
	"github.com/opensourceai/go-api-service/pkg/gredis"
	"github.com/opensourceai/go-api-service/pkg/logging"
	"github.com/opensourceai/go-api-service/pkg/setting"
	"github.com/opensourceai/go-api-service/pkg/util"
	"github.com/opensourceai/go-api-service/routers"
	"log"
	"net/http"
)

func init() {
	setting.Setup()
	mysql.Setup()
	logging.Setup()
	if err := gredis.Setup(); err != nil {
		panic(err)
	}

	util.Setup()
}

// @title Tuboshu Service Api
// @version 2.0
// @description Tuboshu Service Api
// @termsOfService https://github.com/opensourceai
func main() {
	gin.SetMode(setting.ServerSetting.RunMode)

	routersInit := routers.InitRouter()
	readTimeout := setting.ServerSetting.ReadTimeout
	writeTimeout := setting.ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	log.Printf("[info] start http server listening %s", endPoint)

	_ = server.ListenAndServe()

	// If you want Graceful Restart, you need a Unix system and download github.com/fvbock/endless
	//endless.DefaultReadTimeOut = readTimeout
	//endless.DefaultWriteTimeOut = writeTimeout
	//endless.DefaultMaxHeaderBytes = maxHeaderBytes
	//server := endless.NewServer(endPoint, routersInit)
	//server.BeforeBegin = func(add string) {
	//	log.Printf("Actual pid is %d", syscall.Getpid())
	//}
	//
	//err := server.ListenAndServe()
	//if err != nil {
	//	log.Printf("Server err: %v", err)
	//}
}
