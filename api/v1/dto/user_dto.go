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
 * @Package dto
 * @Author Quan Chen
 * @Date 2020/3/19
 * @Description 用户相关数据传输结构体
 *
 */
package dto

type UserDTO struct {
	ID          int    `json:"id" valid:"Required;Min(1)"` // 用户ID
	Name        string `json:"name"`                       // 名称
	Password    string `json:"password"`                   // 密码
	Description string `json:"description"`                // 描述
	Sex         string `json:"sex" valid:"MaxSize(1)"`     // 性别
	AvatarSrc   string `json:"avatar_src"`                 // 头像地址
	Email       string `json:"email" valid:"Email"`        // 电子邮件
	WebSite     string `json:"web_site"`                   // 网站
	Company     string `json:"company"`                    // 公司
	Position    string `json:"position"`                   // 职位
}
