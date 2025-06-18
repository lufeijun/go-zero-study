package tools

import "demo/demo/internal/types"

// 确实有内存逃逸的现象
func ResponseInit() (res *types.Response) {
	return &types.Response{
		Status:  0,
		Message: "success",
	}
}

func ResponseSetMsg(res *types.Response, msg string) {
	res.Status = 1
	res.Message = msg
}
