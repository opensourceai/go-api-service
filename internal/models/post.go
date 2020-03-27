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

// 帖子
// 表名:hive_post
type Post struct {
	Model
	Title   string `json:"title" gorm:"column:title" valid:"Required"`       // 标题
	Content string `json:"content" gorm:"column:content" valid:"Required"`   // 内容
	Summary string `json:"summary" gorm:"column:summary"`                    // 摘要
	UserID  int    `json:"user_id" gorm:"column:user_id" valid:"Required"`   // 用户ID
	BoardID int    `json:"board_id" gorm:"column:board_id" valid:"Required"` // 版块ID
}
