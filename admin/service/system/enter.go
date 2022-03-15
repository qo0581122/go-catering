package system

type ServiceGroup struct {
	ApiService
	MenuService
	UserService
	CasbinService
	InitDBService
	BaseMenuService
	AuthorityService
	SystemConfigService
	OperationRecordService
}
