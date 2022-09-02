// Code generated by goctl. DO NOT EDIT.
package types

type Response struct {
	Status  int64       `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type Id struct {
	Id int64 `path:"id"`
}

type RoleAdd struct {
	Id       int64  `json:"id,optional"`
	Name     string `json:"name"`
	IsEnable int64  `json:"is_enable,optional,default=1"`
	Level    int64  `json:"level,optional,default=1"`
	ParentId int64  `json:"parent_id,optional,default=0"`
}

type RoleSearch struct {
	Name     string `json:"name,optional"`
	IsEnable int64  `json:"is_enable,optional,default=-1"`
	Level    int64  `json:"level,optional,default=0"`
	ParentId int64  `json:"parent_id,optional,default=0"`
	Page     int64  `json:"page,optional,default=1"`
}
