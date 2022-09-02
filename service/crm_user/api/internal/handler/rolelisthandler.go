package handler

import (
	"net/http"

	"demo/service/crm_user/api/internal/logic"
	"demo/service/crm_user/api/internal/svc"
	"demo/service/crm_user/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func roleListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RoleSearch
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewRoleListLogic(r.Context(), svcCtx)
		resp, err := l.RoleList(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
