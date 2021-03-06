package system

import (
	"catering/global"
	"catering/model/system"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

var Api = new(api)

type api struct{}

func (a *api) TableName() string {
	return "sys_apis"
}

func (a *api) Initialize() error {
	entities := []system.SysApi{
		{ApiGroup: "base", Method: "POST", Path: "/base/login", Description: "用户登录(必选)"},

		{ApiGroup: "jwt", Method: "POST", Path: "/jwt/jsonInBlacklist", Description: "jwt加入黑名单(退出，必选)"},

		{ApiGroup: "系统用户", Method: "DELETE", Path: "/user/deleteUser", Description: "删除用户"},
		{ApiGroup: "系统用户", Method: "POST", Path: "/user/register", Description: "用户注册(必选)"},
		{ApiGroup: "系统用户", Method: "POST", Path: "/user/getUserList", Description: "获取用户列表"},
		{ApiGroup: "系统用户", Method: "PUT", Path: "/user/setUserInfo", Description: "设置用户信息"},
		{ApiGroup: "系统用户", Method: "PUT", Path: "/user/setSelfInfo", Description: "设置自身信息(必选)"},
		{ApiGroup: "系统用户", Method: "GET", Path: "/user/getUserInfo", Description: "获取自身信息(必选)"},
		{ApiGroup: "系统用户", Method: "POST", Path: "/user/setUserAuthorities", Description: "设置权限组"},
		{ApiGroup: "系统用户", Method: "POST", Path: "/user/changePassword", Description: "修改密码（建(选择)"},
		{ApiGroup: "系统用户", Method: "POST", Path: "/user/setUserAuthority", Description: "修改用户角色(必选)"},
		{ApiGroup: "系统用户", Method: "POST", Path: "/user/resetPassword", Description: "重置用户密码"},

		{ApiGroup: "api", Method: "POST", Path: "/api/createApi", Description: "创建api"},
		{ApiGroup: "api", Method: "POST", Path: "/api/deleteApi", Description: "删除Api"},
		{ApiGroup: "api", Method: "POST", Path: "/api/updateApi", Description: "更新Api"},
		{ApiGroup: "api", Method: "POST", Path: "/api/getApiList", Description: "获取api列表"},
		{ApiGroup: "api", Method: "POST", Path: "/api/getAllApis", Description: "获取所有api"},
		{ApiGroup: "api", Method: "POST", Path: "/api/getApiById", Description: "获取api详细信息"},
		{ApiGroup: "api", Method: "DELETE", Path: "/api/deleteApisByIds", Description: "批量删除api"},

		{ApiGroup: "角色", Method: "POST", Path: "/authority/copyAuthority", Description: "拷贝角色"},
		{ApiGroup: "角色", Method: "POST", Path: "/authority/createAuthority", Description: "创建角色"},
		{ApiGroup: "角色", Method: "POST", Path: "/authority/deleteAuthority", Description: "删除角色"},
		{ApiGroup: "角色", Method: "PUT", Path: "/authority/updateAuthority", Description: "更新角色信息"},
		{ApiGroup: "角色", Method: "POST", Path: "/authority/getAuthorityList", Description: "获取角色列表"},
		{ApiGroup: "角色", Method: "POST", Path: "/authority/setDataAuthority", Description: "设置角色资源权限"},

		{ApiGroup: "casbin", Method: "POST", Path: "/casbin/updateCasbin", Description: "更改角色api权限"},
		{ApiGroup: "casbin", Method: "POST", Path: "/casbin/getPolicyPathByAuthorityId", Description: "获取权限列表"},

		{ApiGroup: "菜单", Method: "POST", Path: "/menu/addBaseMenu", Description: "新增菜单"},
		{ApiGroup: "菜单", Method: "POST", Path: "/menu/getMenu", Description: "获取菜单树(必选)"},
		{ApiGroup: "菜单", Method: "POST", Path: "/menu/deleteBaseMenu", Description: "删除菜单"},
		{ApiGroup: "菜单", Method: "POST", Path: "/menu/updateBaseMenu", Description: "更新菜单"},
		{ApiGroup: "菜单", Method: "POST", Path: "/menu/getBaseMenuById", Description: "根据id获取菜单"},
		{ApiGroup: "菜单", Method: "POST", Path: "/menu/getMenuList", Description: "分页获取基础menu列表"},
		{ApiGroup: "菜单", Method: "POST", Path: "/menu/getBaseMenuTree", Description: "获取用户动态路由"},
		{ApiGroup: "菜单", Method: "POST", Path: "/menu/getMenuAuthority", Description: "获取指定角色menu"},
		{ApiGroup: "菜单", Method: "POST", Path: "/menu/addMenuAuthority", Description: "增加menu和角色关联关系"},

		{ApiGroup: "分片上传", Method: "POST", Path: "/fileUploadAndDownload/findFile", Description: "寻找目标文件（秒传）"},
		{ApiGroup: "分片上传", Method: "POST", Path: "/fileUploadAndDownload/breakpointContinue", Description: "断点续传"},
		{ApiGroup: "分片上传", Method: "POST", Path: "/fileUploadAndDownload/breakpointContinueFinish", Description: "断点续传完成"},
		{ApiGroup: "分片上传", Method: "POST", Path: "/fileUploadAndDownload/removeChunk", Description: "上传完成移除文件"},

		{ApiGroup: "文件上传与下载", Method: "POST", Path: "/fileUploadAndDownload/upload", Description: "文件上传示例"},
		{ApiGroup: "文件上传与下载", Method: "POST", Path: "/fileUploadAndDownload/deleteFile", Description: "删除文件"},
		{ApiGroup: "文件上传与下载", Method: "POST", Path: "/fileUploadAndDownload/getFileList", Description: "获取上传文件列表"},

		{ApiGroup: "系统服务", Method: "POST", Path: "/system/getServerInfo", Description: "获取服务器信息"},
		{ApiGroup: "系统服务", Method: "POST", Path: "/system/getSystemConfig", Description: "获取配置文件内容"},
		{ApiGroup: "系统服务", Method: "POST", Path: "/system/setSystemConfig", Description: "设置配置文件内容"},

		{ApiGroup: "客户", Method: "PUT", Path: "/customer/customer", Description: "更新客户"},
		{ApiGroup: "客户", Method: "POST", Path: "/customer/customer", Description: "创建客户"},
		{ApiGroup: "客户", Method: "DELETE", Path: "/customer/customer", Description: "删除客户"},
		{ApiGroup: "客户", Method: "GET", Path: "/customer/customer", Description: "获取单一客户"},
		{ApiGroup: "客户", Method: "GET", Path: "/customer/customerList", Description: "获取客户列表"},

		{ApiGroup: "系统字典详情", Method: "PUT", Path: "/sysDictionaryDetail/updateSysDictionaryDetail", Description: "更新字典内容"},
		{ApiGroup: "系统字典详情", Method: "POST", Path: "/sysDictionaryDetail/createSysDictionaryDetail", Description: "新增字典内容"},
		{ApiGroup: "系统字典详情", Method: "DELETE", Path: "/sysDictionaryDetail/deleteSysDictionaryDetail", Description: "删除字典内容"},
		{ApiGroup: "系统字典详情", Method: "GET", Path: "/sysDictionaryDetail/findSysDictionaryDetail", Description: "根据ID获取字典内容"},
		{ApiGroup: "系统字典详情", Method: "GET", Path: "/sysDictionaryDetail/getSysDictionaryDetailList", Description: "获取字典内容列表"},

		{ApiGroup: "系统字典", Method: "POST", Path: "/sysDictionary/createSysDictionary", Description: "新增字典"},
		{ApiGroup: "系统字典", Method: "DELETE", Path: "/sysDictionary/deleteSysDictionary", Description: "删除字典"},
		{ApiGroup: "系统字典", Method: "PUT", Path: "/sysDictionary/updateSysDictionary", Description: "更新字典"},
		{ApiGroup: "系统字典", Method: "GET", Path: "/sysDictionary/findSysDictionary", Description: "根据ID获取字典"},
		{ApiGroup: "系统字典", Method: "GET", Path: "/sysDictionary/getSysDictionaryList", Description: "获取字典列表"},

		{ApiGroup: "操作记录", Method: "POST", Path: "/sysOperationRecord/createSysOperationRecord", Description: "新增操作记录"},
		{ApiGroup: "操作记录", Method: "GET", Path: "/sysOperationRecord/findSysOperationRecord", Description: "根据ID获取操作记录"},
		{ApiGroup: "操作记录", Method: "GET", Path: "/sysOperationRecord/getSysOperationRecordList", Description: "获取操作记录列表"},
		{ApiGroup: "操作记录", Method: "DELETE", Path: "/sysOperationRecord/deleteSysOperationRecord", Description: "删除操作记录"},
		{ApiGroup: "操作记录", Method: "DELETE", Path: "/sysOperationRecord/deleteSysOperationRecordByIds", Description: "批量删除操作历史"},

		{ApiGroup: "断点续传(插件版)", Method: "POST", Path: "/simpleUploader/upload", Description: "插件版分片上传"},
		{ApiGroup: "断点续传(插件版)", Method: "GET", Path: "/simpleUploader/checkFileMd5", Description: "文件完整度验证"},
		{ApiGroup: "断点续传(插件版)", Method: "GET", Path: "/simpleUploader/mergeFileMd5", Description: "上传完成合并文件"},

		{ApiGroup: "email", Method: "POST", Path: "/email/emailTest", Description: "发送测试邮件"},
		{ApiGroup: "email", Method: "POST", Path: "/email/emailSend", Description: "发送邮件示例"},

		{ApiGroup: "excel", Method: "POST", Path: "/excel/importExcel", Description: "导入excel"},
		{ApiGroup: "excel", Method: "GET", Path: "/excel/loadExcel", Description: "下载excel"},
		{ApiGroup: "excel", Method: "POST", Path: "/excel/exportExcel", Description: "导出excel"},
		{ApiGroup: "excel", Method: "GET", Path: "/excel/downloadTemplate", Description: "下载excel模板"},

		{ApiGroup: "city", Method: "GET", Path: "/citys", Description: "城市"},
		{ApiGroup: "city", Method: "POST", Path: "/city/create", Description: "城市"},
		{ApiGroup: "city", Method: "POST", Path: "/city/update", Description: "城市"},
		{ApiGroup: "city", Method: "DELETE", Path: "/city/delete", Description: "城市"},
		{ApiGroup: "city", Method: "GET", Path: "/citys/province", Description: "城市"},

		{ApiGroup: "province", Method: "GET", Path: "/provinces", Description: "省份"},
		{ApiGroup: "province", Method: "GET", Path: "/provinces/all", Description: "省份"},
		{ApiGroup: "province", Method: "POST", Path: "/province/create", Description: "省份"},
		{ApiGroup: "province", Method: "POST", Path: "/province/update", Description: "省份"},
		{ApiGroup: "province", Method: "DELETE", Path: "/province/delete", Description: "省份"},

		{ApiGroup: "district", Method: "GET", Path: "/districts", Description: "区域"},
		{ApiGroup: "district", Method: "GET", Path: "/districts/city", Description: "区域"},
		{ApiGroup: "district", Method: "POST", Path: "/district/create", Description: "区域"},
		{ApiGroup: "district", Method: "POST", Path: "/district/update", Description: "区域"},
		{ApiGroup: "district", Method: "DELETE", Path: "/district/delete", Description: "区域"},

		{ApiGroup: "shop", Method: "GET", Path: "/shops", Description: "店铺详情"},
		{ApiGroup: "shop", Method: "POST", Path: "/shop/create", Description: "店铺详情"},
		{ApiGroup: "shop", Method: "POST", Path: "/shop/update", Description: "店铺详情"},
		{ApiGroup: "shop", Method: "DELETE", Path: "/shop/delete", Description: "店铺详情"},

		{ApiGroup: "店铺分类", Method: "GET", Path: "/shop-categorys", Description: "店铺分类"},
		{ApiGroup: "店铺分类", Method: "GET", Path: "/shop-categorys/all", Description: "店铺分类"},
		{ApiGroup: "店铺分类", Method: "POST", Path: "/shop-category/create", Description: "店铺分类"},
		{ApiGroup: "店铺分类", Method: "POST", Path: "/shop-category/update", Description: "店铺分类"},
		{ApiGroup: "店铺分类", Method: "DELETE", Path: "/shop-category/delete", Description: "店铺分类"},

		{ApiGroup: "代金券", Method: "GET", Path: "/vouchers", Description: "代金券"},
		{ApiGroup: "代金券", Method: "GET", Path: "/vouchers/get-logs", Description: "代金券"},
		{ApiGroup: "代金券", Method: "POST", Path: "/voucher/create", Description: "代金券"},
		{ApiGroup: "代金券", Method: "POST", Path: "/voucher/update", Description: "代金券"},
		{ApiGroup: "代金券", Method: "DELETE", Path: "/voucher/delete", Description: "代金券"},

		{ApiGroup: "优惠券", Method: "GET", Path: "/coupons", Description: "优惠券"},
		{ApiGroup: "优惠券", Method: "GET", Path: "/coupons/get-logs", Description: "优惠券"},
		{ApiGroup: "优惠券", Method: "POST", Path: "/coupon/create", Description: "优惠券"},
		{ApiGroup: "优惠券", Method: "POST", Path: "/coupon/update", Description: "优惠券"},
		{ApiGroup: "优惠券", Method: "DELETE", Path: "/coupon/delete", Description: "优惠券"},

		{ApiGroup: "用户积分", Method: "GET", Path: "/integrations", Description: "用户积分"},
		{ApiGroup: "用户积分", Method: "GET", Path: "/integrations/logs", Description: "用户积分"},
		{ApiGroup: "用户积分", Method: "POST", Path: "/integration/change", Description: "用户积分"},

		{ApiGroup: "用户地址", Method: "GET", Path: "/user-address", Description: "用户地址"},
		{ApiGroup: "用户地址标签", Method: "GET", Path: "/user-address-tags", Description: "用户地址标签"},
		{ApiGroup: "用户地址标签", Method: "POST", Path: "/user-address-tag/create", Description: "用户地址标签"},
		{ApiGroup: "用户地址标签", Method: "POST", Path: "/user-address-tag/update", Description: "用户地址标签"},
		{ApiGroup: "用户地址标签", Method: "DELETE", Path: "/user-address-tag/delete", Description: "用户地址标签"},

		{ApiGroup: "用户VIP等级", Method: "GET", Path: "/vip-levels", Description: "用户VIP等级"},
		{ApiGroup: "用户VIP等级", Method: "GET", Path: "/vip-levels/logs", Description: "用户VIP等级"},
		{ApiGroup: "用户VIP等级", Method: "POST", Path: "/vip-level/create", Description: "用户VIP等级"},
		{ApiGroup: "用户VIP等级", Method: "POST", Path: "/vip-level/update", Description: "用户VIP等级"},
		{ApiGroup: "用户VIP等级", Method: "DELETE", Path: "/vip-level/delete", Description: "用户VIP等级"},
	}
	if err := global.DB.Create(&entities).Error; err != nil {
		return errors.Wrap(err, a.TableName()+"表数据初始化失败!")
	}
	return nil
}

func (a *api) CheckDataExist() bool {
	if errors.Is(global.DB.Where("path = ? AND method = ?", "/excel/downloadTemplate", "GET").First(&system.SysApi{}).Error, gorm.ErrRecordNotFound) {
		return false
	}
	return true
}
