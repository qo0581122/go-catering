package logic

import (
	"context"

	"catering/advertisement/api/internal/svc"
	"catering/advertisement/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetAdRecommendDailyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAdRecommendDailyLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetAdRecommendDailyLogic {
	return GetAdRecommendDailyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAdRecommendDailyLogic) GetAdRecommendDaily(req types.GetAdRecommendDailyProductReq) (*types.GetAdRecommendDailyProductResp, error) {
	// todo: add your logic here and delete this line

	return &types.GetAdRecommendDailyProductResp{}, nil
}
