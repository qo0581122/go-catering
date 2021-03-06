// Code generated by goctl. DO NOT EDIT!
// Source: user.proto

package server

import (
	"context"

	"catering/user/rpc/internal/logic"
	"catering/user/rpc/internal/svc"
	proto "catering/proto/user"
)

type UserHandlerServer struct {
	svcCtx *svc.ServiceContext
}

func NewUserHandlerServer(svcCtx *svc.ServiceContext) *UserHandlerServer {
	return &UserHandlerServer{
		svcCtx: svcCtx,
	}
}

func (s *UserHandlerServer) Register(ctx context.Context, in *proto.RegisterReq) (*proto.RegisterResp, error) {
	l := logic.NewRegisterLogic(ctx, s.svcCtx)
	return l.Register(in)
}

func (s *UserHandlerServer) Login(ctx context.Context, in *proto.LoginReq) (*proto.LoginResp, error) {
	l := logic.NewLoginLogic(ctx, s.svcCtx)
	return l.Login(in)
}

func (s *UserHandlerServer) UserInfo(ctx context.Context, in *proto.UserInfoReq) (*proto.UserInfoResp, error) {
	l := logic.NewUserInfoLogic(ctx, s.svcCtx)
	return l.UserInfo(in)
}

func (s *UserHandlerServer) GetUserAddress(ctx context.Context, in *proto.GetUserAddressReq) (*proto.GetUserAddressResp, error) {
	l := logic.NewGetUserAddressLogic(ctx, s.svcCtx)
	return l.GetUserAddress(in)
}
