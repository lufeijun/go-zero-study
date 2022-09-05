package main

import (
	"flag"
	"fmt"
	"net/http"

	"demo/service/test/api/internal/config"
	"demo/service/test/api/internal/handler"
	"demo/service/test/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
)

var configFile = flag.String("f", "etc/test.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	// 错误处理
	httpx.SetErrorHandler(func(err error) (int, interface{}) {
		fmt.Println("======")

		msg := err.Error()
		fmt.Println(msg)

		type CodeErrorResponse struct {
			Code int    `json:"code"`
			Msg  string `json:"msg"`
		}

		return http.StatusOK, &CodeErrorResponse{
			Code: 1,
			Msg:  "error",
		}

	})

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
