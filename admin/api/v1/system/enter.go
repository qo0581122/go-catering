package system

import "catering/service"

type ApiGroup struct {
	DBApi
	BaseApi
	SystemApi
	CasbinApi
	SystemApiApi
	AuthorityApi
	AuthorityMenuApi
	OperationRecordApi
}

var (
	apiService             = service.ServiceGroupApp.SystemServiceGroup.ApiService
	menuService            = service.ServiceGroupApp.SystemServiceGroup.MenuService
	userService            = service.ServiceGroupApp.SystemServiceGroup.UserService
	initDBService          = service.ServiceGroupApp.SystemServiceGroup.InitDBService
	casbinService          = service.ServiceGroupApp.SystemServiceGroup.CasbinService
	baseMenuService        = service.ServiceGroupApp.SystemServiceGroup.BaseMenuService
	authorityService       = service.ServiceGroupApp.SystemServiceGroup.AuthorityService
	systemConfigService    = service.ServiceGroupApp.SystemServiceGroup.SystemConfigService
	operationRecordService = service.ServiceGroupApp.SystemServiceGroup.OperationRecordService
)
