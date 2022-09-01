package handler

import (
	"net/http"

	"demo/service/crm_user/api/internal/logic"
	"demo/service/crm_user/api/internal/svc"
	"demo/service/crm_user/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func roleEditHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RoleAdd
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewRoleEditLogic(r.Context(), svcCtx)
		resp, err := l.RoleEdit(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
