package logic

import (
	"context"
	"errors"
	"fmt"
	"time"

	"wangjiandev/cloud-disk/core/helper"
	"wangjiandev/cloud-disk/core/internal/svc"
	"wangjiandev/cloud-disk/core/internal/types"
	"wangjiandev/cloud-disk/core/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type MailCodeSendLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMailCodeSendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MailCodeSendLogic {
	return &MailCodeSendLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MailCodeSendLogic) MailCodeSend(req *types.MailCodeSendRequest) (resp *types.MailCodeSendResponse, err error) {
	code := helper.RandCode()
	model.RDB.Set(l.ctx, req.Email, code, time.Second*time.Duration(300))
	err = helper.EmailSendCode(req.Email, code)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("验证码发送失败")
	}
	return &types.MailCodeSendResponse{Message: "验证码发送成功"}, nil
}
