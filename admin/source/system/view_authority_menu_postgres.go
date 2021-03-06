package system

import (
	"strings"

	"catering/global"
	"catering/model/system"

	"github.com/pkg/errors"
)

var ViewAuthorityMenuPostgres = new(viewAuthorityMenuPostgres)

type viewAuthorityMenuPostgres struct{}

func (a *viewAuthorityMenuPostgres) TableName() string {
	var entity system.SysMenu
	return entity.TableName()
}

func (a *viewAuthorityMenuPostgres) Initialize() error {
	var entity AuthorityMenus
	sql := `
	CREATE VIEW @table_name as
	select @menus.id                       as id,
		   @menus.path                     as path,
		   @menus.name                     as name,
		   @menus.icon                     as icon,
		   @menus.sort                     as sort,
		   @menus.title                    as title,
		   @menus.hidden                   as hidden,
		   @menus.parent_id                as parent_id,
		   @menus.component                as component,
		   @menus.keep_alive               as keep_alive,
		   @menus.created_time               as created_time,
		   @menus.updated_time               as updated_time,
		   @menus.deleted_time               as deleted_time,
		   @menus.menu_level               as menu_level,
		   @menus.default_menu             as default_menu,
		   @menus.close_tab                as close_tab,
		   @authorities_menus.sys_base_menu_id      as menu_id,
		   @authorities_menus.sys_authority_authority_id as authority_id
	from (@authorities_menus join @menus on ((@authorities_menus.sys_base_menu_id = @menus.id)));`
	sql = strings.ReplaceAll(sql, "@table_name", a.TableName())
	sql = strings.ReplaceAll(sql, "@menus", "sys_base_menus")
	sql = strings.ReplaceAll(sql, "@authorities_menus", entity.TableName())
	if err := global.DB.Exec(sql).Error; err != nil {
		return errors.Wrap(err, a.TableName()+"视图创建失败!")
	}
	return nil
}
