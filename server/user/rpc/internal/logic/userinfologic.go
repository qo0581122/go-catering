package logic

import (
	"context"

	pb "catering/proto/user"
	"catering/user/rpc/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type UserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserInfoLogic) UserInfo(in *pb.UserInfoReq) (*pb.UserInfoResp, error) {
	info, err := l.svcCtx.UserRepository.FindOne(in.Uid)
	if err != nil {
		logx.Error(err)
		return nil, err
	}
	return &pb.UserInfoResp{
		User: info,
	}, nil
}
