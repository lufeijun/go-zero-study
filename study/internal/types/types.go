// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.2

package types

type Request struct {
	Name string `path:"name,options=you|me"`
}

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}