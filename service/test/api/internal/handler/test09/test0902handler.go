package test09

import (
	"net/http"

	"demo/service/test/api/internal/logic/test09"
	"demo/service/test/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func Test0902Handler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := test09.NewTest0902Logic(r.Context(), svcCtx)
		resp, err := l.Test0902()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
