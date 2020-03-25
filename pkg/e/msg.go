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

package e

var msgFlags = map[int]string{
	SUCCESS:        "成功",
	ERROR:          "失败",
	INVALID_PARAMS: "请求参数错误",
	NOT_FOUND:      "不存在",

	ERROR_AUTH_CHECK_TOKEN_FAIL:    "Token鉴权失败",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "Token已超时",
	ERROR_AUTH_TOKEN:               "Token生成失败",
	ERROR_AUTH:                     "Token错误",
	ERROR_AUTH_NOT_FOUND_TOKEN:     "Token不存在",

	ERROR_USER_EXIST:     "用户已存在",
	ERROR_USER_NOT_EXIST: "用户不存在",

	ERROR_POST_EXIST:     "主题已存在",
	ERROR_POST_NOT_EXIST: "主题不存在",

	ERROR_COMMENT: "评论错误",
}

// GetMsg get error information based on Code
func GetMsg(code int) string {
	msg, ok := msgFlags[code]
	if ok {
		return msg
	}

	return msgFlags[ERROR]
}
