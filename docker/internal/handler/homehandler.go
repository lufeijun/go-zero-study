package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"study/docker/internal/logic"
	"study/docker/internal/svc"
)

func HomeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewHomeLogic(r.Context(), svcCtx)
		resp, err := l.Home()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
