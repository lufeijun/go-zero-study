package funcs

import (
	"demo/gozero115/internal/types"
)

func ResponseInit() (res *types.Response) {
	res = new(types.Response)
	res.Status = 0
	res.Message = "success"
	// res.Data = make(map[string]interface{}, 0)
	// res.Data = make(map[string]string)
	return

}

func ResponseSetMsg(res *types.Response, msg string) {
	res.Status = 1
	res.Message = msg
}
