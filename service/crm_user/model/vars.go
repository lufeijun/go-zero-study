package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var ErrNotFound = sqlx.ErrNotFound

type PageList struct {
	Total     int64       `json:"total"`
	Page      int64       `json:"page"`
	Size      int64       `json:"size"`
	TotalPage int64       `json:"total_page"`
	List      interface{} `json:"list"`
}
