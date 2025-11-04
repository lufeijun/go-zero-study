package tool

import (
	"demo/internal/types"

	"github.com/zeromicro/go-zero/core/netx"

	"github.com/zeromicro/x/errors"
)

// 确实有内存逃逸的现象
func ResponseInit() (res *types.Response) {
	return &types.Response{
		Code: 0,
		Msg:  "success",
	}
}

func ResponseSetMsg(res *types.Response, msg string) {
	res.Code = 1
	res.Msg = msg
}

func ResponseError(msg string, code int) error {
	return errors.New(code, msg)
}

func ResponseErrorDefaultCode(msg string) error {
	return errors.New(400, msg)
}

func GetLocalIp() string {
	return netx.InternalIp()
}
