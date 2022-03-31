package area

import (
	"catering/global"
	"catering/model"
	"catering/model/common/response"
	"errors"
)

var ProvinceService provinceService = NewProvinceService()

func NewProvinceService() provinceService {
	return provinceServiceImpl{}
}

type provinceServiceImpl struct {
}

func (impl provinceServiceImpl) Add(params *model.Province) error {
	return global.DB.Create(&params).Error
}
func (impl provinceServiceImpl) Delete(id uint64) error {
	query := model.City{
		ProvinceId: id,
	}
	var city model.City
	if err := global.DB.Where(&query).Find(&city).Error; err != nil {
		return err
	}
	if city.ID > 0 {
		return errors.New("存在关联，无法删除")
	}
	return global.DB.Delete(&model.Province{}, id).Error
}
func (impl provinceServiceImpl) Update(params *model.Province) error {
	return global.DB.Model(&model.Province{}).Where("id = ?", params.ID).Updates(&params).Error
}

func (impl provinceServiceImpl) Get(params *model.Province) *model.Province {
	var province model.Province
	err := global.DB.Where(&params).First(&province).Error
	if err != nil {
		return nil
	}
	return &province
}

func (impl provinceServiceImpl) List(params *model.Province) []*model.Province {
	var provinces []*model.Province
	if params.ProvinceName != "" {
		global.DB = global.DB.Where("province_name LIKE ?", "%"+params.ProvinceName+"%")
		params.ProvinceName = ""
	}
	err := global.DB.Where(&params).Find(&provinces).Error
	if err != nil {
		return provinces
	}
	return provinces
}

func (impl provinceServiceImpl) Count() int {
	var total int64
	global.DB.Model(&model.Province{}).Count(&total)
	return int(total)
}

func (impl provinceServiceImpl) ListPage(pageNum, pageSize int, params *model.Province) *response.ApiResponse {
	res := &response.ApiResponse{}

	var provinces []*model.Province
	if params.ProvinceName != "" {
		global.DB = global.DB.Where("province_name LIKE ?", "%"+params.ProvinceName+"%")
		params.ProvinceName = ""
	}
	err := global.DB.Where(&params).Scopes(model.Paginate(pageNum, pageSize)).Order("id desc").Find(&provinces).Error
	if err != nil {
		return res
	}
	var total int64
	global.DB.Model(&model.Province{}).Where(&params).Count(&total)
	res.List = provinces
	res.Total = int(total)
	return res
}
