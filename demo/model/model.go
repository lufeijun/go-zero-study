package model

import (
	"reflect"
	"sync"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var (
	once  sync.Once
	beans = make(map[reflect.Type]any) // 表模型单例池
	conn  sqlx.SqlConn
)

func InitDb(dataSource string) {
	once.Do(func() {
		// conn = sqlx.NewMysql( c.Mysql.d )

		// conn.SetConnMaxIdleTime()
		// conn.SetConnMaxLifetime(100 * time.Second)

		// conn.SetMaxOpenConns(c.Mysql.MaxOpenConns)
		// conn.SetMaxIdleConns(c.Mysql.MaxIdleConns)
		// conn.SetConnMaxLifetime(time.Duration(c.Mysql.ConnMaxLifetime) * time.Second)
		// conn.SetConnMaxIdleTime(time.Duration(c.Mysql.ConnMaxIdleTime) * time.Second)

	})
}

// 泛型获取单例
func GetModel[M any]() M {
	var zero M
	t := reflect.TypeOf(zero)
	if v, ok := beans[t]; ok {
		return v.(M)
	}
	// 第一次：调用 NewXxxModel 并缓存
	newFunc := reflect.ValueOf(zero).MethodByName("New") // 约定模型包提供 New 函数
	ins := newFunc.Call([]reflect.Value{
		reflect.ValueOf(conn),
		reflect.ValueOf(nil), // redis 可再传
	})
	m := ins[0].Interface().(M)
	beans[t] = m
	return m
}
