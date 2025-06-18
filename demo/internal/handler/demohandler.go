package handler

import (
	"net/http"

	"demo/demo/internal/logic"
	"demo/demo/internal/svc"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func DemoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewDemoLogic(r.Context(), svcCtx)
		resp, err := l.Demo()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
