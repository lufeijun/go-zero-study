package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ErrormiddlewareMiddleware struct {
}

func NewErrormiddlewareMiddleware() *ErrormiddlewareMiddleware {
	return &ErrormiddlewareMiddleware{}
}

func (m *ErrormiddlewareMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO generate middleware implement function, delete after code implementation

		defer func() {
			// Println executes normally even if there is a panic
			if err := recover(); err != nil {
				fmt.Println("error=", err)
				type CodeErrorResponse struct {
					Code int    `json:"code"`
					Msg  string `json:"msg"`
				}
				w.Header().Add("Content-Type", "application/json")

				result := CodeErrorResponse{
					Code: 1,
					Msg:  fmt.Sprintf("%s", err),
				}

				jsonstr, _ := json.Marshal(result)
				w.Write([]byte(jsonstr))
				return
			}
		}()

		// fmt.Println("11111111111111")
		// Passthrough to next handler if need
		next(w, r)
		// fmt.Println("222222222222")
	}
}
