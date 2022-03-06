package user

import (
	"catering/global"
	"catering/model"
	"catering/model/common/response"
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

var UserIntegrationService userIntegrationService = NewUserIntegrationService()

func NewUserIntegrationService() userIntegrationService {
	return userIntegrationServiceImpl{}
}

type userIntegrationServiceImpl struct {
}

func (impl userIntegrationServiceImpl) IncreaseIntegration(id uint64, changeValue int, changeType int) error {
	global.DB.Transaction(func(tx *gorm.DB) error {
		integration := model.GetUserIntegrationById(id)
		if integration == nil {
			return errors.New("can not find integration info")
		}
		if err := model.IncreaseIntegration(tx, id, changeValue); err != nil {
			return err
		}
		levelInfo := model.GetUserVipLevelById(integration.LevelId)
		sum := integration.CurrentIntegration + changeValue
		for sum <= levelInfo.UpNeedIntegration {
			levelInfo = model.GetUserVipLevelById(levelInfo.NextLevelId)
		}
		if levelInfo.Id != integration.LevelId {
			integration.LevelId = levelInfo.Id
			if err := model.UpdateUserIntegration(integration); err != nil {
				return err
			}
			levelLog := &model.UserVipLevelLog{
				Uid:           integration.Uid,
				BeforeLevelId: integration.LevelId,
				AfterLevelId:  levelInfo.Id,
				UpTime:        model.Time(time.Now()),
			}
			if err := model.AddUserVipLevelLog(levelLog); err != nil {
				return err
			}
		}
		integrationLog := &model.UserIntegrationLog{
			Uid:         integration.Uid,
			ChangeType:  changeType,
			ChangeValue: changeValue,
			BeforeValue: integration.CurrentIntegration,
			AfterValue:  sum,
			ChangeTime:  model.Time(time.Now()),
		}
		if err := model.AddUserIntegrationLog(integrationLog); err != nil {
			return err
		}
		return nil
	})
	return nil
}

func (impl userIntegrationServiceImpl) DecreaseIntegration(id uint64, changeValue int, changeType int) error {
	global.DB.Transaction(func(tx *gorm.DB) error {
		integration := model.GetUserIntegrationById(id)
		if integration == nil {
			return errors.New("can not find integration info")
		}
		if err := model.DecreaseIntegration(tx, id, changeValue); err != nil {
			return err
		}
		sum := integration.CurrentIntegration + changeValue
		integrationLog := &model.UserIntegrationLog{
			Uid:         integration.Uid,
			ChangeType:  changeType,
			ChangeValue: changeValue,
			BeforeValue: integration.CurrentIntegration,
			AfterValue:  sum,
			ChangeTime:  model.Time(time.Now()),
		}
		if err := model.AddUserIntegrationLog(integrationLog); err != nil {
			return err
		}
		return nil
	})
	return nil
}

func (impl userIntegrationServiceImpl) GetById(id uint64) *model.UserIntegration {
	return model.GetUserIntegrationById(id)
}
func (impl userIntegrationServiceImpl) GetByUin(uin uint64) *model.UserIntegration {
	return model.GetUserIntegrationByUin(uin)
}
func (impl userIntegrationServiceImpl) List(params *model.UserIntegration) []*model.UserIntegration {
	return model.ListUserIntegration(params)
}

func (impl userIntegrationServiceImpl) Count() int {
	return model.CountUserAddress()
}

type IntegrationResponse struct {
	*model.UserIntegration
	LevelName string `json:"level_name"`
	Level     int    `json:"level"`
}

func (impl userIntegrationServiceImpl) ListPage(pageNum, pageSize int, params *model.UserIntegration) *response.ApiResponse {
	userIntegrations, err := model.ListUserIntegrationPage(pageNum, pageSize, params)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	levelService := NewUserVipLevelService()
	var list []IntegrationResponse
	for _, item := range userIntegrations {
		levelInfo := levelService.GetById(item.LevelId)
		list = append(list, IntegrationResponse{
			UserIntegration: item,
			LevelName:       levelInfo.LevelName,
			Level:           levelInfo.Level,
		})
	}
	count := model.CountUserIntegration()
	res := &response.ApiResponse{
		List:  list,
		Total: count,
	}
	return res
}
