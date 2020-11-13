package entity

type Users struct {
	ID            uint   `gorm:"primary_key;auto_increment"`
	Name          string `gorm:"size:255"`
	Email         string `gorm:"size:255"`
	Phone         string `gorm:"size:11"`
	TgNo          string `gorm:"size:255"`
	Wechat        string `gorm:"size:255"`
	Password      string `gorm:"size:255"`
	LastLoginTime uint   `gorm:"default:0"`
	RegisterIp    string `gorm:"size:255"`
	Status        uint   `gorm:"default:0"`
	CreateTime    uint   `gorm:"default:0"`
	UpdateTime    uint   `gorm:"default:0"`
}
