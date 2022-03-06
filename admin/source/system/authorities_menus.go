package system

import (
	"reflect"

	"catering/global"
	"catering/model/system"

	"github.com/pkg/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var AuthoritiesMenus = new(authoritiesMenus)

type authoritiesMenus struct{}

func (a *authoritiesMenus) TableName() string {
	var entity AuthorityMenus
	return entity.TableName()
}

func (a *authoritiesMenus) Initialize() error {
	entities := []AuthorityMenus{
		{BaseMenuId: 1, AuthorityId: "888"},
		{BaseMenuId: 2, AuthorityId: "888"},
		{BaseMenuId: 3, AuthorityId: "888"},
		{BaseMenuId: 4, AuthorityId: "888"},
		{BaseMenuId: 5, AuthorityId: "888"},
		{BaseMenuId: 6, AuthorityId: "888"},
		{BaseMenuId: 7, AuthorityId: "888"},
		{BaseMenuId: 8, AuthorityId: "888"},
		{BaseMenuId: 9, AuthorityId: "888"},
		{BaseMenuId: 10, AuthorityId: "888"},
		{BaseMenuId: 11, AuthorityId: "888"},
		{BaseMenuId: 12, AuthorityId: "888"},
		{BaseMenuId: 13, AuthorityId: "888"},
		{BaseMenuId: 14, AuthorityId: "888"},
		{BaseMenuId: 15, AuthorityId: "888"},
		{BaseMenuId: 16, AuthorityId: "888"},
		{BaseMenuId: 17, AuthorityId: "888"},
		{BaseMenuId: 18, AuthorityId: "888"},
		{BaseMenuId: 19, AuthorityId: "888"},
		{BaseMenuId: 20, AuthorityId: "888"},
		{BaseMenuId: 22, AuthorityId: "888"},
		{BaseMenuId: 23, AuthorityId: "888"},
		{BaseMenuId: 24, AuthorityId: "888"},
		{BaseMenuId: 25, AuthorityId: "888"},
		{BaseMenuId: 27, AuthorityId: "888"},
		{BaseMenuId: 28, AuthorityId: "888"},
		{BaseMenuId: 29, AuthorityId: "888"},
		{BaseMenuId: 30, AuthorityId: "888"},
		{BaseMenuId: 31, AuthorityId: "888"},
		{BaseMenuId: 32, AuthorityId: "888"},
		{BaseMenuId: 33, AuthorityId: "888"},
		{BaseMenuId: 34, AuthorityId: "888"},
		{BaseMenuId: 35, AuthorityId: "888"},
		{BaseMenuId: 36, AuthorityId: "888"},
		{BaseMenuId: 37, AuthorityId: "888"},
		{BaseMenuId: 38, AuthorityId: "888"},
		{BaseMenuId: 39, AuthorityId: "888"},
		{BaseMenuId: 40, AuthorityId: "888"},
		{BaseMenuId: 41, AuthorityId: "888"},
		{BaseMenuId: 42, AuthorityId: "888"},
		{BaseMenuId: 43, AuthorityId: "888"},
		{BaseMenuId: 44, AuthorityId: "888"},
		{BaseMenuId: 45, AuthorityId: "888"},
		{BaseMenuId: 46, AuthorityId: "888"},
		{BaseMenuId: 47, AuthorityId: "888"},
		{BaseMenuId: 48, AuthorityId: "888"},
		{BaseMenuId: 49, AuthorityId: "888"},
		{BaseMenuId: 50, AuthorityId: "888"},
		{BaseMenuId: 51, AuthorityId: "888"},
		{BaseMenuId: 52, AuthorityId: "888"},
		{BaseMenuId: 53, AuthorityId: "888"},
		{BaseMenuId: 54, AuthorityId: "888"},
		{BaseMenuId: 55, AuthorityId: "888"},

		{BaseMenuId: 1, AuthorityId: "8881"},
		{BaseMenuId: 2, AuthorityId: "8881"},
		{BaseMenuId: 8, AuthorityId: "8881"},

		{BaseMenuId: 1, AuthorityId: "9528"},
		{BaseMenuId: 2, AuthorityId: "9528"},
		{BaseMenuId: 3, AuthorityId: "9528"},
		{BaseMenuId: 4, AuthorityId: "9528"},
		{BaseMenuId: 5, AuthorityId: "9528"},
		{BaseMenuId: 6, AuthorityId: "9528"},
		{BaseMenuId: 7, AuthorityId: "9528"},
		{BaseMenuId: 8, AuthorityId: "9528"},
		{BaseMenuId: 9, AuthorityId: "9528"},
		{BaseMenuId: 10, AuthorityId: "9528"},
		{BaseMenuId: 11, AuthorityId: "9528"},
		{BaseMenuId: 12, AuthorityId: "9528"},
		{BaseMenuId: 14, AuthorityId: "9528"},
		{BaseMenuId: 15, AuthorityId: "9528"},
		{BaseMenuId: 16, AuthorityId: "9528"},
		{BaseMenuId: 17, AuthorityId: "9528"},
	}
	if err := global.DB.Create(&entities).Error; err != nil {
		return errors.Wrap(err, a.TableName()+"表数据初始化失败!")
	}
	return nil
}

func (a *authoritiesMenus) CheckDataExist() bool {
	if errors.Is(global.DB.Where("sys_base_menu_id = ? AND sys_authority_authority_id = ?", 17, "9528").First(&AuthorityMenus{}).Error, gorm.ErrRecordNotFound) { // 判断是否存在数据
		return false
	}
	return true
}

type AuthorityMenus struct {
	AuthorityId string `gorm:"column:sys_authority_authority_id"`
	BaseMenuId  uint   `gorm:"column:sys_base_menu_id"`
}

func (a *AuthorityMenus) TableName() string {
	var entity system.SysAuthority
	types := reflect.TypeOf(entity)
	if s, o := types.FieldByName("SysBaseMenus"); o {
		m1 := schema.ParseTagSetting(s.Tag.Get("gorm"), ";")
		return m1["MANY2MANY"]
	}
	return ""
}
