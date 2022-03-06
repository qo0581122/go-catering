package initialize

import (
	"net/http"

	_ "catering/docs"
	"catering/global"
	"catering/middleware"
	"catering/router"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// 初始化总路由

func Routers() *gin.Engine {
	Router := gin.Default()

	// 如果想要不使用nginx代理前端网页，可以修改 web/.env.production 下的
	// VUE_APP_BASE_API = /
	// VUE_APP_BASE_PATH = http://localhost
	// 然后执行打包命令 npm run build。在打开下面4行注释
	// Router.LoadHTMLGlob("./dist/*.html") // npm打包成dist的路径
	// Router.Static("/favicon.ico", "./dist/favicon.ico")
	// Router.Static("/static", "./dist/assets")   // dist里面的静态资源
	// Router.StaticFile("/", "./dist/index.html") // 前端网页入口页面

	Router.StaticFS(global.Config.Local.Path, http.Dir(global.Config.Local.Path)) // 为用户头像和文件提供静态地址
	// Router.Use(middleware.LoadTls())  // 打开就能玩https了
	global.Log.Info("use middleware logger")
	// 跨域，如需跨域可以打开下面的注释
	// Router.Use(middleware.Cors()) // 直接放行全部跨域请求
	//Router.Use(middleware.CorsByRules()) // 按照配置的规则放行跨域请求
	global.Log.Info("use middleware cors")
	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	global.Log.Info("register swagger handler")
	// 方便统一添加路由组前缀 多服务器上线使用

	// 获取路由组实例
	systemRouter := router.RouterGroupApp.System
	exampleRouter := router.RouterGroupApp.Example
	areaRouter := router.RouterGroupApp.Area
	couponRouter := router.RouterGroupApp.Coupon
	voucherRouter := router.RouterGroupApp.Voucher
	productRouter := router.RouterGroupApp.Product
	shopRouter := router.RouterGroupApp.Shop
	userRouter := router.RouterGroupApp.User
	orderRouter := router.RouterGroupApp.Order
	PublicGroup := Router.Group("")
	{
		// 健康监测
		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(200, "ok")
		})
	}
	{
		systemRouter.InitBaseRouter(PublicGroup) // 注册基础功能路由 不做鉴权
		systemRouter.InitInitRouter(PublicGroup) // 自动初始化相关
	}
	PrivateGroup := Router.Group("")
	PrivateGroup.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		systemRouter.InitApiRouter(PrivateGroup)                 // 注册功能api路由
		systemRouter.InitUserRouter(PrivateGroup)                // 注册用户路由
		systemRouter.InitMenuRouter(PrivateGroup)                // 注册menu路由
		systemRouter.InitSystemRouter(PrivateGroup)              // system相关路由
		systemRouter.InitCasbinRouter(PrivateGroup)              // 权限相关路由
		systemRouter.InitAuthorityRouter(PrivateGroup)           // 注册角色路由
		systemRouter.InitSysDictionaryRouter(PrivateGroup)       // 字典管理
		systemRouter.InitSysOperationRecordRouter(PrivateGroup)  // 操作记录
		systemRouter.InitSysDictionaryDetailRouter(PrivateGroup) // 字典详情管理

		exampleRouter.InitExcelRouter(PrivateGroup)                 // 表格导入导出
		exampleRouter.InitCustomerRouter(PrivateGroup)              // 客户路由
		exampleRouter.InitFileUploadAndDownloadRouter(PrivateGroup) // 文件上传下载功能路由

		areaRouter.InitCityRouter(PrivateGroup)
		areaRouter.InitDistrictRouter(PrivateGroup)
		areaRouter.InitProvinceRouter(PrivateGroup)

		couponRouter.InitCouponRouter(PrivateGroup)

		voucherRouter.InitVoucherRouter(PrivateGroup)

		productRouter.InitProductAttributeRouter(PrivateGroup)
		productRouter.InitProductAttributeValueRouter(PrivateGroup)
		productRouter.InitProductBatchRouter(PrivateGroup)
		productRouter.InitProductCategoryRouter(PrivateGroup)
		productRouter.InitProductRouter(PrivateGroup)

		shopRouter.InitShopCategoryRouter(PrivateGroup)
		shopRouter.InitShopRouter(PrivateGroup)

		userRouter.InitIntegrationRouter(PrivateGroup)
		userRouter.InitUserAddressRouter(PrivateGroup)
		userRouter.InitUserAddressTagRouter(PrivateGroup)
		userRouter.InitUserVipLevelRouter(PrivateGroup)

		orderRouter.InitOrderRouter(PrivateGroup)
	}

	InstallPlugin(PublicGroup, PrivateGroup) // 安装插件

	global.Log.Info("router register success")
	return Router
}
