package example

import (
	"catering/global"
	"catering/model/common/request"
	"catering/model/example"
	"catering/model/system"
	systemService "catering/service/system"
)

type CustomerService struct{}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateExaCustomer
//@description: 创建客户
//@param: e model.ExaCustomer
//@return: err error

func (exa *CustomerService) CreateExaCustomer(e example.ExaCustomer) (err error) {
	err = global.DB.Create(&e).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteFileChunk
//@description: 删除客户
//@param: e model.ExaCustomer
//@return: err error

func (exa *CustomerService) DeleteExaCustomer(e example.ExaCustomer) (err error) {
	err = global.DB.Delete(&e).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateExaCustomer
//@description: 更新客户
//@param: e *model.ExaCustomer
//@return: err error

func (exa *CustomerService) UpdateExaCustomer(e *example.ExaCustomer) (err error) {
	err = global.DB.Save(e).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetExaCustomer
//@description: 获取客户信息
//@param: id uint
//@return: err error, customer model.ExaCustomer

func (exa *CustomerService) GetExaCustomer(id uint) (err error, customer example.ExaCustomer) {
	err = global.DB.Where("id = ?", id).First(&customer).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetCustomerInfoList
//@description: 分页获取客户列表
//@param: sysUserAuthorityID string, info request.PageInfo
//@return: err error, list interface{}, total int64

func (exa *CustomerService) GetCustomerInfoList(sysUserAuthorityID string, info request.PageInfo) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.DB.Model(&example.ExaCustomer{})
	var a system.SysAuthority
	a.AuthorityId = sysUserAuthorityID
	err, auth := systemService.AuthorityServiceApp.GetAuthorityInfo(a)
	if err != nil {
		return
	}
	var dataId []string
	for _, v := range auth.DataAuthorityId {
		dataId = append(dataId, v.AuthorityId)
	}
	var CustomerList []example.ExaCustomer
	err = db.Where("sys_user_authority_id in ?", dataId).Count(&total).Error
	if err != nil {
		return err, CustomerList, total
	} else {
		err = db.Limit(limit).Offset(offset).Preload("SysUser").Where("sys_user_authority_id in ?", dataId).Find(&CustomerList).Error
	}
	return err, CustomerList, total
}
