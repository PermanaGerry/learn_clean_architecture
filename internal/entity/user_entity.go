package entity

type User struct {
	ID       string `gorm:"column:id;primary_key"`
	Password string `gorm:"column:password"`
	Name     string `gorm:"column:name"`
	Token    string `gorm:"token"`
	CreateAt int64  `gorm:"column:created_at;autoCreateTime:milli"`
	UpdateAt int64  `gorm:"column:update_at;autoCreateTime:milli;autoUpdateTime:milli"`
}

func (u *User) TableName() string {
	return "users"
}
