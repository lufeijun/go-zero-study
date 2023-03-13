package user

import (
	"net/http"

	"demo/gozero115/internal/logic/user"
	"demo/gozero115/internal/svc"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func DetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := user.NewDetailLogic(r.Context(), svcCtx)
		resp, err := l.Detail()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
