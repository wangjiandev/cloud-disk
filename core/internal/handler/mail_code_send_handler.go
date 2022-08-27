package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"wangjiandev/cloud-disk/core/internal/logic"
	"wangjiandev/cloud-disk/core/internal/svc"
	"wangjiandev/cloud-disk/core/internal/types"
)

func MailCodeSendHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.MailCodeSendRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewMailCodeSendLogic(r.Context(), svcCtx)
		resp, err := l.MailCodeSend(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
