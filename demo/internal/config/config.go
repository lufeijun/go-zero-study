// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf

	Mysql struct {
		DataSource string

		MaxOpenConns    int `json:",default=100"`
		MaxIdleConns    int `json:",default=20"`
		ConnMaxLifetime int `json:",default=3600"` // 秒
		ConnMaxIdleTime int `json:",default=600"`  // 秒
	}
	JwtAuth struct {
		AccessSecret string
		AccessExpire int64
	}

	// Auth struct { // JWT 认证需要的密钥和过期时间配置
	// 	AccessSecret string
	// 	AccessExpire int64
	// }
}
