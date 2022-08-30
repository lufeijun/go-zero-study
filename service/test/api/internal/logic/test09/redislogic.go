package test09

import (
	"context"
	"fmt"

	"demo/service/test/api/internal/svc"
	"demo/service/test/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

type RedisLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRedisLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RedisLogic {
	return &RedisLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RedisLogic) Redis() (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	demo01()

	return &types.Response{}, nil
}

func demo01() {
	redisLockKey := fmt.Sprintf("%v%v", "redisTpl", "headId")

	fmt.Println("======")
	fmt.Println(redisLockKey)
	fmt.Println("======")

	redis := redis.New("192.168.0.22:6379")

	// redis.Pass = "123456"

	// err := redis.Set("go-zero", "hello world")

	// if err != nil {
	// 	fmt.Println(err.Error())
	// }

	str, err := redis.Get("go-zero")

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(str)

	// // 1. New redislock
	// redisLock := redis.NewRedisLock(redisConn, redisLockKey)
	// // 2. 可选操作，设置 redislock 过期时间
	// redisLock.SetExpire(redisLockExpireSeconds)
	// if ok, err := redisLock.Acquire(); !ok || err != nil {
	// 	return nil, errors.New("当前有其他用户正在进行操作，请稍后重试")
	// }
	// defer func() {
	// 	recover()
	// 	// 3. 释放锁
	// 	redisLock.Release()
	// }()
}
