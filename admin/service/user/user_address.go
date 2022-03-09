package user

import (
	"catering/global"
	"catering/model"
	"catering/model/common/response"
	"fmt"
)

type UserAddressDetail struct {
	Addr     *model.UserAddress      `json:"address"`
	Province *model.Province         `json:"province"`
	City     *model.City             `json:"city"`
	District *model.District         `json:"district"`
	Tags     []*model.UserAddressTag `json:"tags"`
}

var UserAddressService = NewUserAddressService()

func NewUserAddressService() userAddressService {
	return userAddressServiceImpl{}
}

type userAddressServiceImpl struct {
}

func (impl userAddressServiceImpl) GetOne(params *model.UserAddress) *model.UserAddress {
	var res model.UserAddress
	err := global.DB.Where(params).First(&res).Error
	if err != nil {
		return nil
	}
	return &res
}
func (impl userAddressServiceImpl) List(params *model.UserAddress) []*model.UserAddress {
	var address []*model.UserAddress
	err := global.DB.Where(&params).Find(&address).Error
	if err != nil {
		return nil
	}
	return address
}

func (impl userAddressServiceImpl) Count() int {
	var count int64
	global.DB.Model(&model.UserAddress{}).Count(&count)
	return int(count)
}

func (impl userAddressServiceImpl) ListPage(pageNum, pageSize int, params *model.UserAddress) *response.ApiResponse {
	var address []*model.UserAddress
	err := global.DB.Where(&params).Scopes(model.Paginate(pageNum, pageSize)).Preload("Province").Preload("City").Preload("District").Preload("AddressTag").Find(&address).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	var total int64
	global.DB.Model(&model.UserAddress{}).Count(&total)
	res := &response.ApiResponse{List: address, Total: int(total)}
	return res
}
