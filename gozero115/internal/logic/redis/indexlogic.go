package redis

import (
	"context"
	"fmt"
	"time"

	"demo/gozero115/internal/svc"
	"demo/gozero115/internal/types"
	"demo/gozero115/tool/funcs"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

type IndexLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewIndexLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IndexLogic {
	return &IndexLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *IndexLogic) Index() (resp *types.Response, err error) {
	// todo: add your logic here and delete this line
	resp = funcs.ResponseInit()
	// err = errors.New("报错")

	// l.svcCtx.Config.Jipeng

	redisFunc(l)

	return
}

// 使用原装 redis
func redisFunc(l *IndexLogic) {
	// redis := redis.New("192.168.0.87:6379", func(r *redis.Redis) {
	// 	r.Pass = "123456"
	// 	r.Type = "node"
	// })

	// redis := redis.New("192.168.0.87:6379", func(r *redis.Redis) {
	// 	r.Pass = "123456"
	// 	r.Type = "node"
	// })

	redisconfig := redis.RedisConf{
		Host: l.svcCtx.Config.JpRedisConfig.Host,
		Pass: l.svcCtx.Config.JpRedisConfig.Pass,
		Type: l.svcCtx.Config.JpRedisConfig.Type,
	}
	redisClient, err := redis.NewRedis(redisconfig)
	if err != nil {
		panic("redis 链接失败" + err.Error())
	}

	time.Sleep(time.Second * 10)

	str, err := redisClient.Get("name")
	if err != nil {
		panic("redis 链接失败" + err.Error())
	}

	fmt.Println(str)

}
