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

type UserRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRegisterLogic {
	return &UserRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRegisterLogic) UserRegister(req *types.UserRegisterRequest) (resp *types.UserRegisterResponse, err error) {
	s, err := model.RDB.Get(l.ctx, req.Email).Result()
	if err != nil {
		return nil, err
	}
	if s != req.Code {
		err = errors.New("验证码错误")
		return
	}

	// 判断用户是否存在
	count, err := model.Engine.Where("name = ?", req.Name).Count(&model.UserBasic{})
	if err != nil {
		return nil, err
	}
	if count > 0 {
		err = errors.New("用户已经存在")
		return
	}

	user := &model.UserBasic{
		Identity: helper.RandUUID(),
		Name:     req.Name,
		Email:    req.Email,
		Password: helper.Md5(req.Password),
	}
	i, err := model.Engine.Insert(user)
	if err != nil {
		return nil, err
	}
	fmt.Println(i)
	return &types.UserRegisterResponse{
		Message: "注册成功",
	}, nil
}
