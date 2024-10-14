package handler

import (
	"net/http"

	"demo/study/internal/logic"
	"demo/study/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func StudyHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewStudyLogic(r.Context(), svcCtx)
		resp, err := l.Study()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
