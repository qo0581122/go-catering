// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"catering/integration/api/internal/svc"

	"github.com/tal-tech/go-zero/rest"
)

func RegisterHandlers(engine *rest.Server, serverCtx *svc.ServiceContext) {
	engine.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/integration/logs",
				Handler: GetUserIntegrationLogsHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/integration/vipLogs",
				Handler: GetUserVipUpLogsHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/integration/vip",
				Handler: GetUserVipInfoHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/integration/change",
				Handler: ChangeUserIntegrationHandler(serverCtx),
			},
		},
	)
}