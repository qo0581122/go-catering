package logic

import (
	"context"

	"catering/integration/rpc/internal/svc"
	. "catering/proto/integration"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetUserIntegrationChangeLogsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserIntegrationChangeLogsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserIntegrationChangeLogsLogic {
	return &GetUserIntegrationChangeLogsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserIntegrationChangeLogsLogic) GetUserIntegrationChangeLogs(in *GetUserIntegrationChangeLogsReq) (*GetUserIntegrationChangeLogsResp, error) {
	resp := &GetUserIntegrationChangeLogsResp{}
	var (
		uid   = in.GetUid()
		begin = in.GetBegin()
		count = in.GetCount()
	)
	logs := l.svcCtx.IntegrationRepository.GetUserIntegrationChangeLogs(uid)
	len := int32(len(logs))
	var temp []*UserIntegrationLog
	//如果大于最长
	if begin > len {
		logs = []UserIntegrationLog{}
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
