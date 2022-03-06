package user

import (
	"catering/model"
	"catering/model/common/response"
	"fmt"
)

type UserAddressDetail struct {
	Addr     *model.UserAddress  `json:"address"`
	Province *model.Province     `json:"province"`
	City     *model.City         `json:"city"`
	District *model.District     `json:"district"`
	Tags     []*model.AddressTag `json:"tags"`
}

var UserAddressService = NewUserAddressService()

func NewUserAddressService() userAddressService {
	return userAddressServiceImpl{}
}

type userAddressServiceImpl struct {
}

func (impl userAddressServiceImpl) GetById(id uint64) *model.UserAddress {
	return model.GetUserAddressById(id)
}
func (impl userAddressServiceImpl) List(params *model.UserAddress) []*model.UserAddress {
	return model.ListUserAddress(params)
}

func (impl userAddressServiceImpl) Count() int {
	return model.CountUserAddress()
}

func (impl userAddressServiceImpl) ListPage(pageNum, pageSize int, params *model.UserAddress) *response.ApiResponse {
	addr, err := model.ListUserAddressPage(pageNum, pageSize, params)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	var details []*UserAddressDetail
	for _, item := range addr {
		tags := model.ListAddressTagByAddressId(int(item.Id))

		pro := model.GetProvinceById(item.ProvinceId)
		ci := model.GetCityById(item.CityId)
		di := model.GetDistrictById(item.DistinctId)
		details = append(details, &UserAddressDetail{
			Addr:     item,
			Tags:     tags,
			Province: pro,
			City:     ci,
			District: di,
		})
	}
	total := impl.Count()
	res := &response.ApiResponse{List: details, Total: total}
	return res
}
