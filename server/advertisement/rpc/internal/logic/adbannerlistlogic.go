package logic

import (
	"context"

	"catering/advertisement/rpc/internal/svc"
	proto "catering/proto/advertisement"

	"github.com/tal-tech/go-zero/core/logx"
)

type AdBannerListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAdBannerListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdBannerListLogic {
	return &AdBannerListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AdBannerListLogic) AdBannerList(in *proto.GetAdBannerListReq) (*proto.GetAdBannerListResp, error) {
	var resp proto.GetAdBannerListResp
	banners := l.svcCtx.AdBannerRepository.GetList()
	resp.Banners = banners
	return &resp, nil
}
