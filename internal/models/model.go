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

type Model struct {
	ID         int `gorm:"primary_key" json:"id"` // ID
	CreatedOn  int `json:"created_on"`            // 新建时间
	ModifiedOn int `json:"modified_on"`           // 修改时间
	DeletedOn  int `json:"deleted_on,omitempty"`  // 删除时间 当未被删除时,即为0,忽略此字段
}
