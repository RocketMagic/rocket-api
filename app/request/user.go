package request

// 注册接受参数
type RegisterParams struct {
	Name     string `json:"name"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
