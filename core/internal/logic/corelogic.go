package logic

import (
	"bytes"
	"context"
	"encoding/json"

	"wangjiandev/cloud-disk/core/internal/svc"
	"wangjiandev/cloud-disk/core/internal/types"
	"wangjiandev/cloud-disk/core/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type CoreLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCoreLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CoreLogic {
	return &CoreLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CoreLogic) Core(req *types.Request) (resp *types.Response, err error) {
	// 获取用户列表
	data := make([]*model.UserBasic, 0)
	err = model.Engine.Find(&data)
	if err != nil {
		logx.Error("数据获取失败")
		return nil, err
	}
	b, err := json.Marshal(data)
	if err != nil {
		logx.Error("数据转化失败")
		return nil, err
	}
	dst := new(bytes.Buffer)
	err = json.Indent(dst, b, "", " ")
	if err != nil {
		logx.Error("数据格式化失败")
		return nil, err
	}
	return &types.Response{
		Message: dst.String(),
	}, nil
}
