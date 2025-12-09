// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package main

import (
	"context"
	"flag"
	"fmt"
	"go/types"
	"net/http"

	"demo/internal/config"
	"demo/internal/handler"
	"demo/internal/svc"
	"demo/tool"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/zeromicro/x/errors"
	xhttp "github.com/zeromicro/x/http"
)

var configFile = flag.String("f", "etc/demo-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	// server := rest.MustNewServer(c.RestConf)
	// jwt 认证失败处理逻辑
	server := rest.MustNewServer(c.RestConf, rest.WithUnauthorizedCallback(func(w http.ResponseWriter, r *http.Request, err error) {
		fmt.Println(err)

		w.WriteHeader(http.StatusOK)
		httpx.OkJson(w, map[string]any{
			"code": 401,
			"msg":  "令牌无效或已过期",
		})
	}))

	defer server.Stop()

	// model.InitDb(c.Mysql.DataSource)

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	// httpx.SetErrorHandler 仅在调用了 httpx.Error 处理响应时才有效。
	httpx.SetErrorHandler(func(err error) (int, any) {
		switch e := err.(type) {
		case *errors.CodeMsg:
			return http.StatusOK, xhttp.BaseResponse[types.Nil]{
				Code: e.Code,
				Msg:  e.Msg,
			}
		default:
			// return http.StatusInternalServerError, e.Error()
			return http.StatusOK, xhttp.BaseResponse[types.Nil]{
				Code: 400,
				Msg:  e.Error(),
			}
		}
	})

	// 消费队列
	go func() {
		serviceGroup := service.NewServiceGroup()
		defer serviceGroup.Stop()
		for _, mq := range tool.Consumers(c, context.Background(), ctx) {
			serviceGroup.Add(mq)
		}
		fmt.Printf("kafka queue at %s:%d...\n", c.Host, c.Port)
		serviceGroup.Start()
	}()

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
