//+build wireinject

package api

import (
	"github.com/google/wire"
	"github.com/opensourceai/go-api-service/internal/dao/mysql"
)

func InitApi() (*Api, func(), error) {
	panic(wire.Build(Provider, mysql.NewDao))
}
