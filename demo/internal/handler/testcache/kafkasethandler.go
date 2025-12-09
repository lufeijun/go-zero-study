// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package testcache

import (
	"net/http"

	"demo/internal/logic/testcache"
	"demo/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func KafkaSetHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := testcache.NewKafkaSetLogic(r.Context(), svcCtx)
		resp, err := l.KafkaSet()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
