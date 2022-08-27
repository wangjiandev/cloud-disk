package model

import (
	"github.com/go-redis/redis/v9"
	_ "github.com/lib/pq"
	"github.com/zeromicro/go-zero/core/logx"
	"xorm.io/xorm"
)

var Engine = Init()
var RDB = InitRedis()

func Init() *xorm.Engine {
	engine, err := xorm.NewEngine("postgres", "postgresql://wangjian:123456@127.0.0.1:5432/cloud-disk?sslmode=disable")
	if err != nil {
		logx.Error("xorm NewEngine Fatal %v", err)
		return nil
	}
	return engine
}

func InitRedis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return rdb
}
