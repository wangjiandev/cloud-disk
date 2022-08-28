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
	s, err := l.svcCtx.RDB.Get(l.ctx, req.Email).Result()
	if err != nil {
		err = errors.New("验证码不存在")
		return
	}
	if s != req.Code {
		err = errors.New("验证码错误")
		return
	}

	// 判断用户是否存在
	count, err := l.svcCtx.Engine.Where("name = ?", req.Name).Count(&model.UserBasic{})
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
	_, err = l.svcCtx.Engine.Insert(user)
	if err != nil {
		return nil, err
	}
	// 注册成功后删除验证码
	_, err = l.svcCtx.RDB.Del(l.ctx, req.Email).Result()
	if err != nil {
		fmt.Println(err)
	}
	return &types.UserRegisterResponse{
		Message: "注册成功",
	}, nil
}
