package logic

import (
	pb "catering/proto/user"
	"catering/user/api/internal/svc"
	"catering/user/api/internal/types"
	"context"

	"github.com/tal-tech/go-zero/core/logx"
)

type InfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) InfoLogic {
	return InfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InfoLogic) Info(req types.UserInfoReq) (*types.UserInfoResp, error) {
	resp := &types.UserInfoResp{}
	info, err := l.svcCtx.UserService.UserInfo(l.ctx, &pb.UserInfoReq{
		Uid: req.Uid,
	})
	if info == nil || err != nil {
		return nil, err
	}
	resp.Info = info.GetUser()
	return resp, nil
}
