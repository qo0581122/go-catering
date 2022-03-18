package logic

import (
	"context"

	"catering/advertisement/api/internal/svc"
	"catering/advertisement/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetAdBannerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAdBannerLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetAdBannerLogic {
	return GetAdBannerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAdBannerLogic) GetAdBanner(req types.GetAdBannerListReq) (*types.GetAdBannerListResp, error) {
	// todo: add your logic here and delete this line

	return &types.GetAdBannerListResp{}, nil
}
