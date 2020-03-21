package api

import (
	"github.com/google/wire"
	v1 "github.com/opensourceai/go-api-service/api/v1"
)

type Api struct {
	BoardApi   *v1.BoardApi
	PostApi    *v1.PostApi
	UserAPi    *UserApi
	CommentAPi *v1.CommentApi
}

var providerApi = wire.NewSet(
	v1.ProviderBoard,
	v1.ProviderPost,
	v1.ProviderComment,
	ProviderAuth,
)
var Provider = wire.NewSet(providerApi, wire.Struct(new(Api), "*"))
