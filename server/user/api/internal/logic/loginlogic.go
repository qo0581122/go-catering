package logic

import (
	"context"

	pb "catering/proto/user"
	"catering/user/api/internal/svc"
	"catering/user/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) LoginLogic {
	return LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req types.LoginReq) (*types.LoginResp, error) {
	// todo: add your logic here and delete this line
	resp := &types.LoginResp{}
	info, err := l.svcCtx.UserService.Login(l.ctx, &pb.LoginReq{
		Mobie:    req.Mobie,
		Password: req.Password,
	})
	if info == nil || err != nil {
		return nil, err
	}
	resp.Info = info.GetUser()

	return resp, nil
}
