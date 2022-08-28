package svc

import (
	"wangjiandev/cloud-disk/core/internal/config"
	"wangjiandev/cloud-disk/core/internal/middleware"
	"wangjiandev/cloud-disk/core/model"

	"github.com/go-redis/redis/v9"
	"github.com/zeromicro/go-zero/rest"
	"xorm.io/xorm"
)

type ServiceContext struct {
	Config config.Config
	Engine *xorm.Engine
	RDB    *redis.Client
	Auth   rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Engine: model.Init(c.Postgres.DataSource),
		RDB:    model.InitRedis(c.Redis.Addr),
		Auth:   middleware.NewAuthMiddleware().Handle,
	}
}
