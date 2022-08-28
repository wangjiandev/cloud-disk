package svc

import (
	"wangjiandev/cloud-disk/core/internal/config"
	"wangjiandev/cloud-disk/core/model"

	"github.com/go-redis/redis/v9"
	"xorm.io/xorm"
)

type ServiceContext struct {
	Config config.Config
	Engine *xorm.Engine
	RDB    *redis.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Engine: model.Init(c.Postgres.DataSource),
		RDB:    model.InitRedis(c.Redis.Addr),
	}
}
