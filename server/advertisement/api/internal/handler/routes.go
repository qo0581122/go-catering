// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"catering/advertisement/api/internal/svc"

	"github.com/tal-tech/go-zero/rest"
)

func RegisterHandlers(engine *rest.Server, serverCtx *svc.ServiceContext) {
	engine.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/ad/banner",
				Handler: GetAdBannerHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/ad/hotActivitie",
				Handler: GetHotActivitieHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/ad/dailyProduct",
				Handler: GetAdRecommendDailyHandler(serverCtx),
			},
		},
	)
}