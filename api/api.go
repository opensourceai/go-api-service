/*
 *    Copyright 2020 opensourceai
 *
 *    Licensed under the Apache License, Version 2.0 (the "License");
 *    you may not use this file except in compliance with the License.
 *    You may obtain a copy of the License at
 *
 *        http://www.apache.org/licenses/LICENSE-2.0
 *
 *    Unless required by applicable law or agreed to in writing, software
 *    distributed under the License is distributed on an "AS IS" BASIS,
 *    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *    See the License for the specific language governing permissions and
 *    limitations under the License.
 */

package api

import (
	"github.com/google/wire"
	v1 "github.com/opensourceai/go-api-service/api/v1"
)

type Api struct {
	BoardApi   *v1.BoardApi
	PostApi    *v1.PostApi
	CommentAPi *v1.CommentApi
	AuthApi    *AuthApi
	UserApi    *v1.UserApi
}

//Api构造器集合
var providerApi = wire.NewSet(
	v1.ProviderBoard,
	v1.ProviderPost,
	v1.ProviderComment,
	ProviderAuth,
	v1.ProviderUser,
)

//结构构造器,避免后续调用构造器集合时代码冗余(不使用的情况下,若后续调用构造器集合现需要多此初始化集合)
var Provider = wire.NewSet(providerApi, wire.Struct(new(Api), "*"))
