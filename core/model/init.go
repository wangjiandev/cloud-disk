package model

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/zeromicro/go-zero/core/logx"
	"xorm.io/xorm"
)

var Engine = Init()

func Init() *xorm.Engine {
	engine, err := xorm.NewEngine("mysql", "root:123456@/cloud-disk?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		logx.Error("xorm NewEngine Fatal %v", err)
		return nil
	}
	return engine
}
