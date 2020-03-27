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

package models

// 用户
// 表名:hive_user
type User struct {
	Model
	Username    string `json:"username" grom:"column:username;not null" valid:"Required"`      // 用户名
	Name        string `json:"name"  grom:"column:name;not null" valid:"Required;MaxSize(50)"` // 昵称
	Password    string `json:"password" grom:"column:password;not null" valid:"Required"`      // 密码
	Description string `json:"description" grom:"column:description" valid:"MaxSize(200)"`     // 描述
	Sex         string `json:"sex" grom:"column:sex;not null" valid:"MaxSize(1)"`              // 性别
	AvatarSrc   string `json:"avatar_src" grom:"column:avatar_src;not null"`                   // 头像地址
	Email       string `json:"email" grom:"column:email" valid:"Required;Email;MaxSize(100)"`  // 电子邮件
	WebSite     string `json:"web_site" grom:"column:web_site" valid:"MaxSize(50)"`            // 网站
	Company     string `json:"company" grom:"column:company" valid:"MaxSize(50)"`              // 公司
	Position    string `json:"position" grom:"column:position" valid:"MaxSize(50)"`            // 职位
}
