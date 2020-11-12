package entity

type Users struct {
	ID            uint   `gorm:"primary_key;auto_increment"`
	Name          string `gorm:"column:name size:255"`
	Email         string `gorm:"column:email size:255"`
	Phone         string `gorm:"column:phone size:11"`
	TgNo          string `gorm:"column:tg_no size:255"`
	Wechat        string `gorm:"column:wechat size:255"`
	Password      string `gorm:"column:password size:255"`
	LastLoginTime uint   `gorm:"column:last_login_time"`
	RegisterIp    string `gorm:"column:register_ip size:255"`
	Status        uint   `gorm:"column:status"`
	CreateTime    uint   `gorm:"column:create_time"`
	UpdateTime    uint   `gorm:"column:update_time"`
}
