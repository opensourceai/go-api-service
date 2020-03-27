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

// 评论
// 表名:hive_comment
type Comment struct {
	Model
	CommentContent string `json:"comment_content"` // 评论内容
	PostID         int    `json:"post_id"`         // 主题ID
	FromUserID     int    `json:"from_user_id"`    // 评论者ID
	ToUserID       int    `json:"to_user_id"`      // 被评论者ID
}
