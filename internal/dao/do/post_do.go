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

/*
 * @Package do
 * @Author Quan Chen
 * @Date 2020/3/19
 * @Description
 *
 */
package do

import (
	"github.com/opensourceai/go-api-service/internal/models"
)

type CommentDO struct {
	models.Comment
	FromUser UserDo `json:"from_user" gorm:"foreignkey:FromUserId"`
	ToUser   UserDo `json:"to_user" gorm:"foreignkey:ToUserId"`
}

// 命名数据表名
func (CommentDO) TableName() string {
	return "comment"
}

// UserDo
type UserDo struct {
	models.Model
	Username    string `json:"username" grom:"column:username;not null"`     // 用户名
	Name        string `json:"name"  grom:"column:name;not null"`            // 昵称
	Description string `json:"description" grom:"column:description"`        // 描述
	Sex         int    `json:"sex" grom:"column:sex;not null"`               // 性别
	AvatarSrc   string `json:"avatar_src" grom:"column:avatar_src;not null"` // 头像地址
	Email       string `json:"email" grom:"column:email"`                    // 电子邮件
	WebSite     string `json:"web_site" grom:"column:web_site"`              // 网站
	Company     string `json:"company" grom:"column:company"`                // 公司
	Position    string `json:"position" grom:"column:position"`              // 职位
}

// 命名数据表名
func (UserDo) TableName() string {
	return "user"
}
