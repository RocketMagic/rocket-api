package util

var response map[string]interface{}

// 请求返回json
func RespReturn(code int, data interface{}, msg string) map[string]interface{} {
	response["code"] = code
	response["data"] = data
	response["msg"] = msg

	return response
}
