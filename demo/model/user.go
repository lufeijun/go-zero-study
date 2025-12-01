package model

import (
	"time"
)

type User struct {
	ID        uint32     `gorm:"column:id;primaryKey;autoIncrement;comment:用户id" json:"id"`
	Name      *string    `gorm:"column:name;type:varchar(255);comment:用户名称" json:"name"`
	Phone     *string    `gorm:"column:phone;type:varchar(255);comment:手机号" json:"phone"`
	Password  *string    `gorm:"column:password;type:varchar(255)" json:"password"`
	CreatedAt *time.Time `gorm:"column:created_at;comment:创建时间" json:"created_at"`
	UpdatedAt time.Time  `gorm:"column:updated_at;comment:更新时间" json:"updated_at"`
}

func (User) TableName() string {
	return "users"
}
