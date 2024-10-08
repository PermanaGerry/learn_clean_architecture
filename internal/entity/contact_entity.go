package entity

type Contact struct {
	ID        string `gorm:"column:id;primary_key"`
	FirstName string `gorm:"column:first_name"`
	LastName  string `gorm:"column:last_name"`
	Email     string `gorm:"column:email"`
	Phone     string `gorm:"column:phone"`
	UserId    string `gorm:"column:user_id"`
	CreateAt  int64  `gorm:"column:create_at;autoCreateTime:milli"`
	UpdateAt  int64  `gorm:"column:update_at;autoCreateTime:milli;autoUpdateTime:milli"`
	User      User   `gorm:"foreignKey:user_id;reference:id"`
}

func (c *Contact) TableName() string {
	return "contacts"
}
