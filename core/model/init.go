package model

import (
	"github.com/go-redis/redis/v9"
	_ "github.com/lib/pq"
	"github.com/zeromicro/go-zero/core/logx"
	"xorm.io/xorm"
)

func Init(dataSource string) *xorm.Engine {
	engine, err := xorm.NewEngine("postgres", dataSource)
	if err != nil {
		logx.Error("xorm NewEngine Fatal %v", err)
		return nil
	}
	return engine
}

func InitRedis(addr string) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return rdb
}
