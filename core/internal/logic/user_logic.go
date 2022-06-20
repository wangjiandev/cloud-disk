package logic

import (
	"context"
	"errors"
	"fmt"

	"wangjiandev/cloud-disk/core/helper"
	"wangjiandev/cloud-disk/core/internal/svc"
	"wangjiandev/cloud-disk/core/internal/types"
	"wangjiandev/cloud-disk/core/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLogic {
	return &UserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLogic) User(req *types.LoginRequest) (resp *types.LoginResponse, err error) {

	fmt.Println(req.Name)
	fmt.Println(req.Password)

	// 数据库查询用户
	user := new(model.UserBasic)
	has, err := model.Engine.Where("name=? AND password=?", req.Name, helper.Md5(req.Password)).Get(user)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, errors.New("用户名或密码错误")
	}

	// 生成token返回
	token, err := helper.GenerateToken(user.Id, user.Identity, user.Name)
	if err != nil {
		return nil, err
	}
	return &types.LoginResponse{
		Token: token,
	}, nil
}
