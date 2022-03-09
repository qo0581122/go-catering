package user

import (
	"catering/global"
	"catering/model"
	"catering/model/common/response"
	"errors"
	"fmt"
	"sort"
)

var UserVipLevelService userVipLevelService = NewUserVipLevelService()

func NewUserVipLevelService() userVipLevelService {
	return userVipLevelServiceImpl{}
}

type userVipLevelServiceImpl struct {
}

func (impl userVipLevelServiceImpl) Add(params *model.UserVipLevel, preLevelId uint64) error {
	tx := global.DB.Begin()
	level := params.Level
	if err := tx.Create(&params).Error; err != nil {
		tx.Rollback()
		return err
	}
	levelInfo := params
	if err := tx.Model(&model.UserVipLevel{}).Where("id = ?", params.Id).Updates(&model.UserVipLevel{
		Id:          preLevelId,
		NextLevelId: levelInfo.Id,
	}).Error; err != nil {
		tx.Rollback()
		return err
	}
	for {
		nextLevelId := levelInfo.NextLevelId
		if nextLevelId == 0 {
			break
		}
		temp := &model.UserVipLevel{
			Id:    nextLevelId,
			Level: level + 1,
		}
		if err := tx.Model(&model.UserVipLevel{}).Where("id = ?", params.Id).Updates(&temp).Error; err != nil {
			tx.Rollback()
			return err
		}
		if err := global.DB.Where("id = ?", nextLevelId).First(&levelInfo).Error; err != nil {
			tx.Rollback()
			return err
		}
		level++
	}
	tx.Commit()
	return nil
}

func (impl userVipLevelServiceImpl) Adds(params []model.UserVipLevel) error {
	fmt.Println(params)
	addFailError := errors.New("add levels fail")
	tx := global.DB.Begin()
	sort.Slice(params, func(i, j int) bool {
		return params[i].Level > params[j].Level
	})
	var level model.UserVipLevel
	var err error
	for i, item := range params {
		if i == 0 {
			item.NextLevelId = 0
		} else {
			item.NextLevelId = level.Id
		}
		if err = tx.Create(&item).Error; err != nil {
			tx.Rollback()
			return addFailError
		}
		level = item
	}
	tx.Commit()
	return nil
}

func (impl userVipLevelServiceImpl) Delete(id uint64) error {
	return global.DB.Delete(&model.UserVipLevel{}, id).Error

}
func (impl userVipLevelServiceImpl) Update(params *model.UserVipLevel) error {
	return global.DB.Model(&model.UserVipLevel{}).Where("id = ?", params.Id).Updates(&params).Error
}

func (impl userVipLevelServiceImpl) GetOne(params *model.UserVipLevel) *model.UserVipLevel {
	var res model.UserVipLevel
	err := global.DB.Where(&params).First(&res).Error
	if err != nil {
		return nil
	}
	return &res
}

func (impl userVipLevelServiceImpl) List(params *model.UserVipLevel) []*model.UserVipLevel {
	var vips []*model.UserVipLevel
	err := global.DB.Where(&params).Find(&vips).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return vips
}

func (impl userVipLevelServiceImpl) Count() int {
	var count int64
	global.DB.Model(&model.UserVipLevel{}).Count(&count)
	return int(count)
}

func (impl userVipLevelServiceImpl) ListPage(pageNum, pageSize int, params *model.UserVipLevel) *response.ApiResponse {
	var userVipLevels []*model.UserVipLevel
	err := global.DB.Where(&params).Scopes(model.Paginate(pageNum, pageSize)).Find(&userVipLevels).Error
	if err != nil {
		return nil
	}
	var count int64
	global.DB.Model(&model.UserVipLevel{}).Count(&count)
	res := &response.ApiResponse{
		List:  userVipLevels,
		Total: int(count),
	}
	return res
}
