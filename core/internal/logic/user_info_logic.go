package logic

import (
	"context"
	"errors"
	"fmt"

	"wangjiandev/cloud-disk/core/internal/svc"
	"wangjiandev/cloud-disk/core/internal/types"
	"wangjiandev/cloud-disk/core/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo(req *types.InfoRequest) (resp *types.InfoResponse, err error) {
	user := model.UserBasic{}
	has, err := model.Engine.Where("identity=?", req.Identity).Get(&user)
	if err != nil {
		return nil, err
	}
	fmt.Println(user)
	if !has {
		return nil, errors.New("用户不存在")
	}
	return &types.InfoResponse{
		Name:  user.Name,
		Email: user.Email,
	}, nil
}
