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

package dao

import "github.com/opensourceai/go-api-service/internal/models"

type UserDao interface {
	// 添加用户
	DaoAdd(user *models.User) error
	// 根据id软删除用户
	DaoDeleteById(ids ...int) error
	// 编辑用户信息
	DaoEdit(user *models.User) error
	// 获取该用户名的用户信息
	DaoGetUserByUsername(username string) (err error, user models.User)
	// 查询ids中所有用户
	DaoFindByIds(ids ...int) (users []models.User, err error)
}
