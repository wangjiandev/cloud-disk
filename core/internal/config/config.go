package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	Postgres struct {
		DataSource string
	}
	Redis struct {
		Addr string
	}
}
