package handler

import (
	"net/http"

	"demo/service/crm_user/api/internal/logic"
	"demo/service/crm_user/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func userListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewUserListLogic(r.Context(), svcCtx)
		resp, err := l.UserList()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
