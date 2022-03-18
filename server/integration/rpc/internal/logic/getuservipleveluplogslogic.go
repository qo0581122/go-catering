package logic

import (
	"context"

	"catering/integration/rpc/internal/svc"
	. "catering/proto/integration"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetUserVipLevelUpLogsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserVipLevelUpLogsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserVipLevelUpLogsLogic {
	return &GetUserVipLevelUpLogsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserVipLevelUpLogsLogic) GetUserVipLevelUpLogs(in *GetUserVipLevelUpLogsReq) (*GetUserVipLevelUpLogsResp, error) {
	resp := &GetUserVipLevelUpLogsResp{}
	var (
		uid   = in.GetUid()
		begin = in.GetBegin()
		count = in.GetCount()
	)
	logs := l.svcCtx.IntegrationRepository.GetUserVipLevelUpLogs(uid)
	len := int32(len(logs))
	var temp []*UserVipUpLog
	//如果大于最长
	if begin > len {
		logs = []UserVipUpLog{}
	} else if begin+count > len { //begin小于len但begin+count>len
		logs = logs[begin:]
	} else { //begin小于len且begin+count<len
		logs = logs[begin : begin+count]
	}
	for index, _ := range logs {
		temp = append(temp, &logs[index])
	}
	resp.Logs = temp
	return resp, nil
}
