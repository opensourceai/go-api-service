# Go API Service
**基于Gin构建的土拨鼠社区基础服务**

[![Go](https://github.com/opensourceai/go-api-service/workflows/Go/badge.svg)](https://github.com/opensourceai/go-api-service/actions)

## How to run

### Required

- Mysql:线上已部署
- Redis 
> 本地部署.使用`test/docker-compose.yaml`本地部署
### Conf

配置文件 `conf/app.ini`

```
[database]
Type = mysql
User = root
Password = root 
Host = 127.0.0.1:3306
Name = blog
TablePrefix = blog_

[redis]
Host = 127.0.0.1:6379
Password =
MaxIdle = 30
MaxActive = 30
IdleTimeout = 200
...
```

### Run
```
$ cd c/go-api-service

$ go run main.go 
```

项目信息和现有API

```
[info] replacing callback `gorm:update_time_stamp` from D:/opensourceai/go-api-service/dao/mysql/dao.go:32
[info] replacing callback `gorm:update_time_stamp` from D:/opensourceai/go-api-service/dao/mysql/dao.go:33
[info] replacing callback `gorm:delete` from D:/opensourceai/go-api-service/dao/mysql/dao.go:34
[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:	export GIN_MODE=release
 - using code:	gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /export/*filepath         --> github.com/gin-gonic/gin.(*RouterGroup).createStaticHandler.func1 (3 handlers)
[GIN-debug] HEAD   /export/*filepath         --> github.com/gin-gonic/gin.(*RouterGroup).createStaticHandler.func1 (3 handlers)
[GIN-debug] GET    /upload/images/*filepath  --> github.com/gin-gonic/gin.(*RouterGroup).createStaticHandler.func1 (3 handlers)
[GIN-debug] HEAD   /upload/images/*filepath  --> github.com/gin-gonic/gin.(*RouterGroup).createStaticHandler.func1 (3 handlers)
[GIN-debug] GET    /qrcode/*filepath         --> github.com/gin-gonic/gin.(*RouterGroup).createStaticHandler.func1 (3 handlers)
[GIN-debug] HEAD   /qrcode/*filepath         --> github.com/gin-gonic/gin.(*RouterGroup).createStaticHandler.func1 (3 handlers)
[GIN-debug] GET    /auth                     --> github.com/opensourceai/go-api-service/routers/api.GetAuth (3 handlers)
[GIN-debug] GET    /swagger/*any             --> github.com/swaggo/gin-swagger.WrapHandler.func1 (3 handlers)
[GIN-debug] POST   /upload                   --> github.com/opensourceai/go-api-service/routers/api.UploadImage (3 handlers)
[info] start http server listening :8000

```
## Swagger
```shell script
swag init # 生成文档
```
Swagger doc: http://localhost:8000/swagger/index.html

## Test Api
```
[GET] /auth Get Auth
```
- userName:hive
- password:hive
## Features

- RESTful API
- Gorm
- Swagger
- logging
- Jwt-go
- Gin
- Graceful restart or stop (fvbock/endless)
- App configurable
- Cron
- Redis