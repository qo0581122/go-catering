package logic

import (
	"context"

	"catering/area/pkg/errors"
	pb "catering/proto/user"
	"catering/user/rpc/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetUserAddressLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserAddressLogic {
	return &GetUserAddressLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserAddressLogic) GetUserAddress(req *pb.GetUserAddressReq) (*pb.GetUserAddressResp, error) {
	resp := &pb.GetUserAddressResp{}
	if req.Uid <= 0 {
		return resp, errors.InternalServerError("user_id is not format")
	}
	allAddress := l.svcCtx.UserAddressRepository.SelectByUserId(req.GetUid())
	var address []*pb.UserAddress
	for index, item := range allAddress {
		if item.DeletedTime != "" {
			continue
		}
		address = append(address, allAddress[index])
	}
	for _, item := range address {
		tags := l.svcCtx.UserAddressRepository.SelectAddressTag(item.GetId())
		item.Tags = tags
	}
	resp.Address = address
	return resp, nil
}
