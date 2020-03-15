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
配置环境变量

- GO111MODULE=on
- GOPROXY=https://goproxy.io

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

## Dev
1. fork repository
2. clone repository
   ```shell script
   # [git@github.com:chenquan/go-api-service.git]为自己账户下仓库地址
   git clone git@github.com:chenquan/go-api-service.git
   cd go-api-service 
   ```
3. 开发新功能前必须拉去主库代码到本地master
    
    1. 新建远程库连接(`只需第一次clone之后设置`)
    ```shell script
     git remote add opensourceai git@github.com:opensourceai/go-api-service.git
    ```
    2. 拉取主库最新master代码到本地master
    ```shell script
    git pull opensourceai master:master --rebase -f
    ```
4. 开发新功能/修改
    1. 从本地master分支新建出`feature-*`（`fix-*`）分支（`*`表示对应新功能名称）
    2. 开发完毕之后，push到自己账号下的仓库
    3. 通过PR使用squash方式合并到主库
5. 后续开发循环`3`,`4`
   
   
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

## Tools

- GoLand
