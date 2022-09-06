package models

type User struct {
	ID         uint   `json:"id"`
	Name       string `json:"name"`
	Gender     string `json:"gender"`
	Password   string `json:"-"`
	CreateTime string `json:"create_time"`
	UpdateTime string `json:"update_time"`
}
