package models

import (
	"time"
)

type User struct {
	ID         uint      `json:"id"`
	Name       string    `json:"name"`
	Gender     string    `json:"gender"`
	Password   string    `json:"-"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
}
