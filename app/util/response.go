package util

// 请求返回json
func RespReturn(code int, data interface{}, msg string) map[string]interface{} {
	return map[string]interface{}{
		"code": code,
		"data": data,
		"msg":  msg,
	}

}
