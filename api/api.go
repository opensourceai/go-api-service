package api

import (
	"github.com/google/wire"
	v1 "github.com/opensourceai/go-api-service/api/v1"
)

type Api struct {
	BoardApi *v1.BoardApi
	PostApi  *v1.PostApi
	UserAPi  *UserApi
}

var providerApi = wire.NewSet(
	v1.ProviderBoard,
	v1.ProviderPost,
	ProviderAuth,
)
var Provider = wire.NewSet(providerApi, wire.Struct(new(Api), "*"))
