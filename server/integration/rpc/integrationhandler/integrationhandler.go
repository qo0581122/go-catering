// Code generated by goctl. DO NOT EDIT!
// Source: integration.proto

package integrationhandler

import (
	"context"

	proto "catering/proto/integration"

	"github.com/tal-tech/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	ChangeUserIntegrationReq         = proto.ChangeUserIntegrationReq
	ChangeUserIntegrationResp        = proto.ChangeUserIntegrationResp
	GetUserIntegrationChangeLogsReq  = proto.GetUserIntegrationChangeLogsReq
	GetUserIntegrationChangeLogsResp = proto.GetUserIntegrationChangeLogsResp
	GetUserVipInfoReq                = proto.GetUserVipInfoReq
	GetUserVipInfoResp               = proto.GetUserVipInfoResp
	GetUserVipLevelUpLogsReq         = proto.GetUserVipLevelUpLogsReq
	GetUserVipLevelUpLogsResp        = proto.GetUserVipLevelUpLogsResp
	UserIntegration                  = proto.UserIntegration
	UserIntegrationLog               = proto.UserIntegrationLog
	UserVipLevel                     = proto.UserVipLevel
	UserVipUpLog                     = proto.UserVipUpLog

	IntegrationHandler interface {
		GetUserIntegrationChangeLogs(ctx context.Context, in *GetUserIntegrationChangeLogsReq, opts ...grpc.CallOption) (*GetUserIntegrationChangeLogsResp, error)
		GetUserVipLevelUpLogs(ctx context.Context, in *GetUserVipLevelUpLogsReq, opts ...grpc.CallOption) (*GetUserVipLevelUpLogsResp, error)
		ChangeUserIntegration(ctx context.Context, in *ChangeUserIntegrationReq, opts ...grpc.CallOption) (*ChangeUserIntegrationResp, error)
		GetUserVipInfo(ctx context.Context, in *GetUserVipInfoReq, opts ...grpc.CallOption) (*GetUserVipInfoResp, error)
	}

	defaultIntegrationHandler struct {
		cli zrpc.Client
	}
)

func NewIntegrationHandler(cli zrpc.Client) IntegrationHandler {
	return &defaultIntegrationHandler{
		cli: cli,
	}
}

func (m *defaultIntegrationHandler) GetUserIntegrationChangeLogs(ctx context.Context, in *GetUserIntegrationChangeLogsReq, opts ...grpc.CallOption) (*GetUserIntegrationChangeLogsResp, error) {
	client := proto.NewIntegrationHandlerClient(m.cli.Conn())
	return client.GetUserIntegrationChangeLogs(ctx, in, opts...)
}

func (m *defaultIntegrationHandler) GetUserVipLevelUpLogs(ctx context.Context, in *GetUserVipLevelUpLogsReq, opts ...grpc.CallOption) (*GetUserVipLevelUpLogsResp, error) {
	client := proto.NewIntegrationHandlerClient(m.cli.Conn())
	return client.GetUserVipLevelUpLogs(ctx, in, opts...)
}

func (m *defaultIntegrationHandler) ChangeUserIntegration(ctx context.Context, in *ChangeUserIntegrationReq, opts ...grpc.CallOption) (*ChangeUserIntegrationResp, error) {
	client := proto.NewIntegrationHandlerClient(m.cli.Conn())
	return client.ChangeUserIntegration(ctx, in, opts...)
}

func (m *defaultIntegrationHandler) GetUserVipInfo(ctx context.Context, in *GetUserVipInfoReq, opts ...grpc.CallOption) (*GetUserVipInfoResp, error) {
	client := proto.NewIntegrationHandlerClient(m.cli.Conn())
	return client.GetUserVipInfo(ctx, in, opts...)
}