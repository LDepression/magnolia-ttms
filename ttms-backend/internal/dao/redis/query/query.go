/**
 * @Author: lenovo
 * @Description:
 * @File:  query
 * @Version: 1.0.0
 * @Date: 2023/05/29 10:27
 */

package query

import (
	"context"
	"github.com/go-redis/redis/v8"
)

var client = &Queries{}

type Queries struct {
	redis *redis.Client
}

func New(client *redis.Client) *Queries {
	return &Queries{
		redis: client,
	}
}

func init() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379", // ip:端口
		Password: "123456",         // 密码
		PoolSize: 20,               // 连接池
		DB:       0,                // 默认连接数据库
	})
	_, err := rdb.Ping(context.Background()).Result() // 测试连接
	if err != nil {
		panic(err)
	}
	client = New(rdb)
}
