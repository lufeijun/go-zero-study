// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package middleware

import (
	"context"
	"net/http"
	"time"

	"github.com/zeromicro/go-zero/core/collection"
	"github.com/zeromicro/go-zero/rest/httpx"
)

type TestCacheMiddleware struct {
}

func NewTestCacheMiddleware() *TestCacheMiddleware {
	return &TestCacheMiddleware{}
}

func (m *TestCacheMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO generate middleware implement function, delete after code implementation

		cache, err := collection.NewCache(time.Minute, collection.WithLimit(10))
		if err != nil {
			// 返回错误：直接写入响应，不再调用 next
			w.WriteHeader(http.StatusOK)
			httpx.OkJson(w, map[string]any{
				"code": 401,
				"msg":  "cache init error",
			})
			return // ⚠️ 关键：不要调用 next()
		}

		// Passthrough to next handler if need
		ctx := context.WithValue(r.Context(), "test_cache", cache)
		newReq := r.WithContext(ctx)
		next(w, newReq)
	}
}
