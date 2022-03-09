package user

import (
	"catering/global"
	"catering/model"
	"catering/model/common/response"
	userResponse "catering/model/user/response"
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
		integration := UserIntegrationService.GetOne(&model.UserIntegration{
			Id: id,
		})
		if integration == nil {
			return errors.New("can not find integration info")
		}
		SQL := "UPDATE user_integration SET current_integration = current_integration + ?, history_integration = history_integration + ?  WHERE id = ?"
		if err := tx.Exec(SQL, changeValue, id).Error; err != nil {
			return err
		}
		levelInfo := UserVipLevelService.GetOne(&model.UserVipLevel{
			Id: integration.LevelId,
		})
		sum := integration.CurrentIntegration + changeValue
		for sum <= levelInfo.UpNeedIntegration {
			levelInfo = UserVipLevelService.GetOne(&model.UserVipLevel{
				Id: levelInfo.NextLevelId,
			})
		}
		if levelInfo.Id != integration.LevelId {
			integration.LevelId = levelInfo.Id
			if err := tx.Model(&model.UserIntegration{}).Where("id = ?", integration.Id).Updates(&integration).Error; err != nil {
				return err
			}
			levelLog := &model.UserVipLevelLog{
				Uid:           integration.Uid,
				BeforeLevelId: integration.LevelId,
				AfterLevelId:  levelInfo.Id,
				UpTime:        model.Time(time.Now()),
			}

			if err := UserVipLevelLogService.Add(levelLog); err != nil {
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
		if err := UserIntegrationLogService.Add(integrationLog); err != nil {
			return err
		}
		return nil
	})
	return nil
}

func (impl userIntegrationServiceImpl) DecreaseIntegration(id uint64, changeValue int, changeType int) error {
	global.DB.Transaction(func(tx *gorm.DB) error {
		integration := UserIntegrationService.GetOne(&model.UserIntegration{
			Id: id,
		})
		if integration == nil {
			return errors.New("can not find integration info")
		}
		SQL := "UPDATE user_integration SET current_integration = current_integration - ? WHERE id = ? AND current_integration >= ?"
		if err := tx.Exec(SQL, changeValue, id, changeValue).Error; err != nil {
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
		if err := tx.Create(&integrationLog).Error; err != nil {
			return err
		}
		return nil
	})
	return nil
}

func (impl userIntegrationServiceImpl) GetOne(params *model.UserIntegration) *model.UserIntegration {
	var res model.UserIntegration
	err := global.DB.Where(&params).First(&res).Error
	if err != nil {
		return nil
	}
	return &res
}

func (impl userIntegrationServiceImpl) List(params *model.UserIntegration) []*model.UserIntegration {
	var integrations []*model.UserIntegration
	err := global.DB.Where(&params).Find(&integrations).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return integrations
}

func (impl userIntegrationServiceImpl) Count() int {
	var count int64
	global.DB.Model(&model.UserIntegration{}).Count(&count)
	return int(count)
}

func (impl userIntegrationServiceImpl) ListPage(pageNum, pageSize int, params *model.UserIntegration) *response.ApiResponse {
	var integrations []*model.UserIntegration
	err := global.DB.Where(&params).Scopes(model.Paginate(pageNum, pageSize)).Find(&integrations).Error
	if err != nil {
		return nil
	}
	if err != nil {
		fmt.Println(err)
		return nil
	}
	var list []userResponse.Integration
	for _, item := range integrations {
		levelInfo := UserVipLevelService.GetOne(&model.UserVipLevel{
			Id: item.LevelId,
		})
		list = append(list, userResponse.Integration{
			UserIntegration: item,
			LevelName:       levelInfo.LevelName,
			Level:           levelInfo.Level,
		})
	}
	var count int64
	global.DB.Model(&model.UserIntegration{}).Where(&params).Count(&count)
	res := &response.ApiResponse{
		List:  list,
		Total: int(count),
	}
	return res
}
