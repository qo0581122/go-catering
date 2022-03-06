package system

import (
	"catering/global"
	"catering/model/system"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

var BaseMenu = new(menu)

type menu struct{}

func (m *menu) TableName() string {
	return "sys_base_menus"
}

func (m *menu) Initialize() error {
	entities := []system.SysBaseMenu{
		{MenuLevel: 0, Hidden: false, ParentId: "0", Path: "dashboard", Name: "dashboard", Component: "view/dashboard/index.vue", Sort: 1, Meta: system.Meta{Title: "仪表盘", Icon: "odometer"}},
		{MenuLevel: 0, Hidden: false, ParentId: "0", Path: "about", Name: "about", Component: "view/about/index.vue", Sort: 7, Meta: system.Meta{Title: "关于我们", Icon: "info-filled"}},
		{MenuLevel: 0, Hidden: false, ParentId: "0", Path: "admin", Name: "superAdmin", Component: "view/superAdmin/index.vue", Sort: 3, Meta: system.Meta{Title: "超级管理员", Icon: "user"}},
		{MenuLevel: 0, Hidden: false, ParentId: "3", Path: "authority", Name: "authority", Component: "view/superAdmin/authority/authority.vue", Sort: 1, Meta: system.Meta{Title: "角色管理", Icon: "avatar"}},
		{MenuLevel: 0, Hidden: false, ParentId: "3", Path: "menu", Name: "menu", Component: "view/superAdmin/menu/menu.vue", Sort: 2, Meta: system.Meta{Title: "菜单管理", Icon: "tickets", KeepAlive: true}},
		{MenuLevel: 0, Hidden: false, ParentId: "3", Path: "api", Name: "api", Component: "view/superAdmin/api/api.vue", Sort: 3, Meta: system.Meta{Title: "api管理", Icon: "platform", KeepAlive: true}},
		{MenuLevel: 0, Hidden: false, ParentId: "3", Path: "superUser", Name: "superUser", Component: "view/superAdmin/user/user.vue", Sort: 4, Meta: system.Meta{Title: "系统用户管理", Icon: "coordinate"}},
		{MenuLevel: 0, Hidden: true, ParentId: "0", Path: "person", Name: "person", Component: "view/person/person.vue", Sort: 4, Meta: system.Meta{Title: "个人信息", Icon: "message"}},
		{MenuLevel: 0, Hidden: false, ParentId: "0", Path: "example", Name: "example", Component: "view/example/index.vue", Sort: 6, Meta: system.Meta{Title: "示例文件", Icon: "management"}},
		{MenuLevel: 0, Hidden: false, ParentId: "9", Path: "excel", Name: "excel", Component: "view/example/excel/excel.vue", Sort: 4, Meta: system.Meta{Title: "excel导入导出", Icon: "takeaway-box"}},
		{MenuLevel: 0, Hidden: false, ParentId: "9", Path: "upload", Name: "upload", Component: "view/example/upload/upload.vue", Sort: 5, Meta: system.Meta{Title: "媒体库（上传下载）", Icon: "upload"}},
		{MenuLevel: 0, Hidden: false, ParentId: "9", Path: "breakpoint", Name: "breakpoint", Component: "view/example/breakpoint/breakpoint.vue", Sort: 6, Meta: system.Meta{Title: "断点续传", Icon: "upload-filled"}},
		{MenuLevel: 0, Hidden: false, ParentId: "9", Path: "customer", Name: "customer", Component: "view/example/customer/customer.vue", Sort: 7, Meta: system.Meta{Title: "客户列表（资源示例）", Icon: "avatar"}},
		{MenuLevel: 0, Hidden: false, ParentId: "0", Path: "systemTools", Name: "systemTools", Component: "view/systemTools/index.vue", Sort: 5, Meta: system.Meta{Title: "系统工具", Icon: "tools"}},
		{MenuLevel: 0, Hidden: false, ParentId: "14", Path: "autoCode", Name: "autoCode", Component: "view/systemTools/autoCode/index.vue", Sort: 1, Meta: system.Meta{Title: "代码生成器", Icon: "cpu", KeepAlive: true}},
		{MenuLevel: 0, Hidden: false, ParentId: "14", Path: "formCreate", Name: "formCreate", Component: "view/systemTools/formCreate/index.vue", Sort: 2, Meta: system.Meta{Title: "表单生成器", Icon: "magic-stick", KeepAlive: true}},
		{MenuLevel: 0, Hidden: false, ParentId: "14", Path: "system", Name: "system", Component: "view/systemTools/system/system.vue", Sort: 3, Meta: system.Meta{Title: "系统配置", Icon: "operation"}},
		{MenuLevel: 0, Hidden: false, ParentId: "3", Path: "dictionary", Name: "dictionary", Component: "view/superAdmin/dictionary/sysDictionary.vue", Sort: 5, Meta: system.Meta{Title: "字典管理", Icon: "notebook"}},
		{MenuLevel: 0, Hidden: true, ParentId: "3", Path: "dictionaryDetail/:id", Name: "dictionaryDetail", Component: "view/superAdmin/dictionary/sysDictionaryDetail.vue", Sort: 1, Meta: system.Meta{Title: "字典详情", Icon: "order"}},
		{MenuLevel: 0, Hidden: false, ParentId: "3", Path: "operation", Name: "operation", Component: "view/superAdmin/operation/sysOperationRecord.vue", Sort: 6, Meta: system.Meta{Title: "操作历史", Icon: "pie-chart"}},
		{MenuLevel: 0, Hidden: false, ParentId: "9", Path: "simpleUploader", Name: "simpleUploader", Component: "view/example/simpleUploader/simpleUploader", Sort: 6, Meta: system.Meta{Title: "断点续传（插件版）", Icon: "upload"}},
		{MenuLevel: 0, Hidden: false, ParentId: "0", Path: "https://www.gin-vue-admin.com", Name: "https://www.gin-vue-admin.com", Component: "/", Sort: 0, Meta: system.Meta{Title: "官方网站", Icon: "home-filled"}},
		{MenuLevel: 0, Hidden: false, ParentId: "0", Path: "state", Name: "state", Component: "view/system/state.vue", Sort: 6, Meta: system.Meta{Title: "服务器状态", Icon: "cloudy"}},
		{MenuLevel: 0, Hidden: false, ParentId: "14", Path: "autoCodeAdmin", Name: "autoCodeAdmin", Component: "view/systemTools/autoCodeAdmin/index.vue", Sort: 1, Meta: system.Meta{Title: "自动化代码管理", Icon: "magic-stick"}},
		{MenuLevel: 0, Hidden: true, ParentId: "14", Path: "autoCodeEdit/:id", Name: "autoCodeEdit", Component: "view/systemTools/autoCode/index.vue", Sort: 0, Meta: system.Meta{Title: "自动化代码（复用）", Icon: "magic-stick"}},

		{MenuLevel: 0, Hidden: false, ParentId: "0", Path: "area", Name: "area", Component: "view/area/index.vue", Sort: 0, Meta: system.Meta{Title: "区域管理", Icon: "magic-stick"}},
		{MenuLevel: 0, Hidden: false, ParentId: "26", Path: "areaCity", Name: "areaCity", Component: "view/area/city/index.vue", Sort: 0, Meta: system.Meta{Title: "城市管理", Icon: "magic-stick"}},
		{MenuLevel: 0, Hidden: false, ParentId: "26", Path: "areaProvince", Name: "areaProvince", Component: "view/area/province/index.vue", Sort: 0, Meta: system.Meta{Title: "省份管理", Icon: "magic-stick"}},
		{MenuLevel: 0, Hidden: false, ParentId: "26", Path: "areaDistrict", Name: "areaDistrict", Component: "view/area/district/index.vue", Sort: 0, Meta: system.Meta{Title: "区域管理", Icon: "magic-stick"}},

		{MenuLevel: 0, Hidden: false, ParentId: "0", Path: "shop", Name: "shop", Component: "view/shop/index.vue", Sort: 0, Meta: system.Meta{Title: "店铺管理", Icon: "magic-stick"}},
		{MenuLevel: 0, Hidden: false, ParentId: "30", Path: "shopDetail", Name: "shopDetail", Component: "view/shop/detail/index.vue", Sort: 0, Meta: system.Meta{Title: "店铺详情", Icon: "magic-stick"}},
		{MenuLevel: 0, Hidden: false, ParentId: "30", Path: "shopCategory", Name: "shopCategory", Component: "view/shop/category/index.vue", Sort: 0, Meta: system.Meta{Title: "店铺分类", Icon: "magic-stick"}},

		{MenuLevel: 0, Hidden: false, ParentId: "0", Path: "voucher", Name: "voucher", Component: "view/voucher/index.vue", Sort: 0, Meta: system.Meta{Title: "代金券管理", Icon: "magic-stick"}},
		{MenuLevel: 0, Hidden: false, ParentId: "33", Path: "voucherDetail", Name: "voucherDetail", Component: "view/voucher/detail/index.vue", Sort: 0, Meta: system.Meta{Title: "代金券详情", Icon: "magic-stick"}},
		{MenuLevel: 0, Hidden: false, ParentId: "33", Path: "voucherGetLog", Name: "voucherGetLog", Component: "view/voucher/get_log/index.vue", Sort: 0, Meta: system.Meta{Title: "代金券获取明细", Icon: "magic-stick"}},

		{MenuLevel: 0, Hidden: false, ParentId: "0", Path: "coupon", Name: "coupon", Component: "view/coupon/index.vue", Sort: 0, Meta: system.Meta{Title: "优惠券管理", Icon: "magic-stick"}},
		{MenuLevel: 0, Hidden: false, ParentId: "36", Path: "couponDetail", Name: "couponDetail", Component: "view/coupon/detail/index.vue", Sort: 0, Meta: system.Meta{Title: "优惠券详情", Icon: "magic-stick"}},
		{MenuLevel: 0, Hidden: false, ParentId: "36", Path: "couponGetLog", Name: "couponGetLog", Component: "view/coupon/get_log/index.vue", Sort: 0, Meta: system.Meta{Title: "优惠券获取明细", Icon: "magic-stick"}},

		{MenuLevel: 0, Hidden: false, ParentId: "0", Path: "user", Name: "user", Component: "view/user/index.vue", Sort: 0, Meta: system.Meta{Title: "用户管理", Icon: "magic-stick"}},
		{MenuLevel: 0, Hidden: false, ParentId: "39", Path: "userAddress", Name: "userAddress", Component: "view/user/index.vue", Sort: 0, Meta: system.Meta{Title: "用户地址管理", Icon: "magic-stick"}},
		{MenuLevel: 0, Hidden: false, ParentId: "40", Path: "userAddressDetail", Name: "userAddressDetail", Component: "view/user/address/detail/index.vue", Sort: 0, Meta: system.Meta{Title: "用户地址详情", Icon: "magic-stick"}},
		{MenuLevel: 0, Hidden: false, ParentId: "40", Path: "userAddressTag", Name: "userAddressTag", Component: "view/user/address/tag/index.vue", Sort: 0, Meta: system.Meta{Title: "用户地址标签", Icon: "magic-stick"}},

		{MenuLevel: 0, Hidden: false, ParentId: "39", Path: "userIntegration", Name: "userIntegration", Component: "view/user/index.vue", Sort: 0, Meta: system.Meta{Title: "用户积分管理", Icon: "magic-stick"}},
		{MenuLevel: 0, Hidden: false, ParentId: "43", Path: "userIntegrationDetail", Name: "userIntegrationDetail", Component: "view/user/integration/detail/index.vue", Sort: 0, Meta: system.Meta{Title: "用户积分详情", Icon: "magic-stick"}},
		{MenuLevel: 0, Hidden: false, ParentId: "43", Path: "userIntegrationLog", Name: "userIntegrationLog", Component: "view/user/integration/get_log/index.vue", Sort: 0, Meta: system.Meta{Title: "用户积分获取明细", Icon: "magic-stick"}},

		{MenuLevel: 0, Hidden: false, ParentId: "39", Path: "userVipLevel", Name: "userVipLevel", Component: "view/user/index.vue", Sort: 0, Meta: system.Meta{Title: "用户VIP等级管理", Icon: "magic-stick"}},
		{MenuLevel: 0, Hidden: false, ParentId: "46", Path: "userVipLevelDetail", Name: "userVipLevelDetail", Component: "view//vip_level/level/index.vue", Sort: 0, Meta: system.Meta{Title: "VIP等级", Icon: "magic-stick"}},
		{MenuLevel: 0, Hidden: false, ParentId: "46", Path: "userVipLevelLog", Name: "userVipLevelLog", Component: "view/user/vip_level/log/index.vue", Sort: 0, Meta: system.Meta{Title: "用户VIP升级明细", Icon: "magic-stick"}},

		{MenuLevel: 0, Hidden: false, ParentId: "0", Path: "product", Name: "product", Component: "view/product/index.vue", Sort: 0, Meta: system.Meta{Title: "商品管理", Icon: "magic-stick"}},
		{MenuLevel: 0, Hidden: false, ParentId: "49", Path: "productDetail", Name: "productDetail", Component: "view/product/detail/index.vue", Sort: 0, Meta: system.Meta{Title: "商品详情", Icon: "magic-stick"}},
		{MenuLevel: 0, Hidden: false, ParentId: "49", Path: "productCategory", Name: "productCategory", Component: "view/product/category/index.vue", Sort: 0, Meta: system.Meta{Title: "商品分类", Icon: "magic-stick"}},
		{MenuLevel: 0, Hidden: false, ParentId: "49", Path: "productBatch", Name: "productBatch", Component: "view/product/batch/index.vue", Sort: 0, Meta: system.Meta{Title: "商品物料", Icon: "magic-stick"}},
		{MenuLevel: 0, Hidden: false, ParentId: "49", Path: "productAttribute", Name: "productAttribute", Component: "view/product/index.vue", Sort: 0, Meta: system.Meta{Title: "商品属性管理", Icon: "magic-stick"}},
		{MenuLevel: 0, Hidden: false, ParentId: "53", Path: "productAttributeDetail", Name: "productAttributeDetail", Component: "view/product/attribute/detail/index.vue", Sort: 0, Meta: system.Meta{Title: "商品属性详情", Icon: "magic-stick"}},
		{MenuLevel: 0, Hidden: false, ParentId: "53", Path: "productAttributeValue", Name: "productAttributeValue", Component: "view/product/attribute/value/index.vue", Sort: 0, Meta: system.Meta{Title: "商品属性值", Icon: "magic-stick"}},

		{MenuLevel: 0, Hidden: false, ParentId: "0", Path: "order", Name: "order", Component: "view/order/index.vue", Sort: 0, Meta: system.Meta{Title: "订单管理", Icon: "magic-stick"}},
		{MenuLevel: 0, Hidden: false, ParentId: "56", Path: "orderDetail", Name: "orderDetail", Component: "view/order/detail/index.vue", Sort: 0, Meta: system.Meta{Title: "订单详情", Icon: "magic-stick"}},
	}
	if err := global.DB.Create(&entities).Error; err != nil { // 创建 model.User 初始化数据
		return errors.Wrap(err, m.TableName()+"表数据初始化失败!")
	}
	return nil
}

func (m *menu) CheckDataExist() bool {
	if errors.Is(global.DB.Where("path = ?", "autoCodeEdit/:id").First(&system.SysBaseMenu{}).Error, gorm.ErrRecordNotFound) { // 判断是否存在数据
		return false
	}
	return true
}
