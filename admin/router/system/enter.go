package system

type RouterGroup struct {
	ApiRouter
	SysRouter
	BaseRouter
	InitRouter
	MenuRouter
	UserRouter
	CasbinRouter
	AuthorityRouter
	OperationRecordRouter
}
