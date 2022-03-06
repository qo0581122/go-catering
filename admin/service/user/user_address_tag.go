package user

import (
	"catering/model"
	"catering/model/common/response"
	"fmt"
)

var UserAddressTagService = NewAddressTagService()

func NewAddressTagService() addressTagService {
	return addressTagServiceImpl{}
}

type addressTagServiceImpl struct {
}

func (impl addressTagServiceImpl) Add(params *model.AddressTag) error {
	return model.AddAddressTag(params)
}
func (impl addressTagServiceImpl) Delete(id uint64) error {
	return model.DeleteAddressTag(id)
}
func (impl addressTagServiceImpl) Update(params *model.AddressTag) error {
	return model.UpdateAddressTag(params)
}

func (impl addressTagServiceImpl) GetById(id uint64) *model.AddressTag {
	return model.GetAddressTagById(id)
}
func (impl addressTagServiceImpl) List(params *model.AddressTag) []*model.AddressTag {
	return model.ListAddressTag(params)
}

func (impl addressTagServiceImpl) Count() int {
	return model.CountUserAddress()
}

func (impl addressTagServiceImpl) ListPage(pageNum, pageSize int, params *model.AddressTag) *response.ApiResponse {
	addr, err := model.ListAddressTagPage(pageNum, pageSize, params)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	total := model.CountAddressTag()
	res := &response.ApiResponse{List: addr, Total: total}
	return res
}
