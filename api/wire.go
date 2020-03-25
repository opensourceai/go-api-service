//+build wireinject

package api

import (
	"github.com/google/wire"
	"github.com/opensourceai/go-api-service/internal/dao/mysql"
)
// 初始化构造器Provider并创建好对象
func InitApi() (*Api, func(), error) {
	panic(wire.Build(Provider, mysql.NewDao))
}
