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
	levelInfo, err := model.AddUserVipLevel(tx, params)
	if err != nil {
		tx.Rollback()
		return err
	}
	err = model.UpdateUserVipLevel(tx, &model.UserVipLevel{
		Id:          preLevelId,
		NextLevelId: levelInfo.Id,
	})
	for {
		nextLevelId := levelInfo.NextLevelId
		if nextLevelId == 0 {
			break
		}
		temp := &model.UserVipLevel{
			Id:    nextLevelId,
			Level: level + 1,
		}
		model.UpdateUserVipLevel(tx, temp)
		levelInfo = model.GetUserVipLevelById(nextLevelId)
		level++
	}
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (impl userVipLevelServiceImpl) Adds(params []*model.UserVipLevel) error {
	addFailError := errors.New("add levels fail")
	tx := global.DB.Begin()
	sort.Slice(params, func(i, j int) bool {
		return params[i].Level > params[j].Level
	})
	var level *model.UserVipLevel
	var err error
	for i, item := range params {
		if i == 0 {
			item.NextLevelId = 0
		} else {
			item.NextLevelId = level.Id
		}
		level, err = model.AddUserVipLevel(tx, item)
		if err != nil {
			tx.Rollback()
			return addFailError
		}
	}
	tx.Commit()
	return nil
}

func (impl userVipLevelServiceImpl) Delete(id uint64) error {
	return model.DeleteUserVipLevel(id)
}
func (impl userVipLevelServiceImpl) Update(params *model.UserVipLevel) error {
	return model.UpdateUserVipLevel(nil, params)
}

func (impl userVipLevelServiceImpl) GetById(id uint64) *model.UserVipLevel {
	return model.GetUserVipLevelById(id)
}

func (impl userVipLevelServiceImpl) GetByUin(uin uint64) *model.UserVipLevel {
	return model.GetUserVipLevelByUin(uin)
}
func (impl userVipLevelServiceImpl) List(params *model.UserVipLevel) []*model.UserVipLevel {
	return model.ListUserVipLevel(params)
}

func (impl userVipLevelServiceImpl) Count() int {
	return model.CountUserAddress()
}

func (impl userVipLevelServiceImpl) ListPage(pageNum, pageSize int, params *model.UserVipLevel) *response.ApiResponse {
	userVipLevels, err := model.ListUserVipLevelPage(pageNum, pageSize, params)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	count := model.CountUserVipLevel()
	res := &response.ApiResponse{
		List:  userVipLevels,
		Total: count,
	}
	return res
}
