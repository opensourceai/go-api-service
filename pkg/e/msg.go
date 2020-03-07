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
}

// GetMsg get error information based on Code
func GetMsg(code int) string {
	msg, ok := msgFlags[code]
	if ok {
		return msg
	}

	return msgFlags[ERROR]
}
