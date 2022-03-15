package user

import (
	"catering/global"
	"catering/model"
	"catering/model/common/response"
	"fmt"
)

var UserAddressTagService = NewUserAddressTagService()

func NewUserAddressTagService() addressTagService {
	return addressTagServiceImpl{}
}

type addressTagServiceImpl struct {
}

func (impl addressTagServiceImpl) Add(params *model.UserAddressTag) error {
	return global.DB.Create(&params).Error
}
func (impl addressTagServiceImpl) Delete(id uint64) error {
	return global.DB.Delete(&model.UserAddressTag{}, id).Error
}
func (impl addressTagServiceImpl) Update(params *model.UserAddressTag) error {
	return global.DB.Model(&model.UserAddressTag{}).Where("id = ?", params.ID).Updates(&params).Error

}

func (impl addressTagServiceImpl) GetOne(params *model.UserAddressTag) *model.UserAddressTag {
	var res model.UserAddressTag
	err := global.DB.Where(params).First(&res).Error
	if err != nil {
		return nil
	}
	return &res
}
func (impl addressTagServiceImpl) List(params *model.UserAddressTag) []*model.UserAddressTag {
	var tags []*model.UserAddressTag
	err := global.DB.Where(&params).Find(&tags).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return tags
}

func (impl addressTagServiceImpl) Count() int {
	var count int64
	global.DB.Model(&model.UserAddressTag{}).Count(&count)
	return int(count)
}

func (impl addressTagServiceImpl) ListPage(pageNum, pageSize int, params *model.UserAddressTag) *response.ApiResponse {
	var tags []*model.UserAddressTag
	err := global.DB.Where(&params).Scopes(model.Paginate(pageNum, pageSize)).Find(&tags).Error
	if err != nil {
		return nil
	}
	if err != nil {
		fmt.Println(err)
		return nil
	}
	var total int64
	global.DB.Model(&model.UserAddressTag{}).Where(&params).Count(&total)
	res := &response.ApiResponse{List: tags, Total: int(total)}
	return res
}
