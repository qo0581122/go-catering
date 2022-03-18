package logic

import (
	"context"

	pb "catering/proto/user"
	"catering/user/api/internal/svc"
	"catering/user/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) RegisterLogic {
	return RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req types.RegisterReq) (*types.RegisterResp, error) {
	// todo: add your logic here and delete this line
	_, err := l.svcCtx.UserService.Register(l.ctx, &pb.RegisterReq{
		Mobie:    req.Mobie,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}

	return &types.RegisterResp{}, nil
}
