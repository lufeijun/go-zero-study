package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf

	Mysql struct {
		DataSource string
	}

	CacheRedis cache.CacheConf

	// roc 客户端配置
	UserRpc    zrpc.RpcClientConf
	ProductRpc zrpc.RpcClientConf
}
