package user

import (
	"catering/global"
	"catering/model"
	"catering/model/common/response"
	userResponse "catering/model/user/response"
	"fmt"
)

var UserIntegrationLogService = NewUserIntegrationLogService()

func NewUserIntegrationLogService() userIntegrationLogService {
	return userIntegrationLogServiceImpl{}
}

type userIntegrationLogServiceImpl struct {
}

func (impl userIntegrationLogServiceImpl) Add(params *model.UserIntegrationLog) error {
	return global.DB.Create(&params).Error
}

func (impl userIntegrationLogServiceImpl) GetOne(params *model.UserIntegrationLog) *model.UserIntegrationLog {
	var res model.UserIntegrationLog
	err := global.DB.Where(&params).First(&res).Error
	if err != nil {
		return nil
	}
	return &res
}
func (impl userIntegrationLogServiceImpl) List(params *model.UserIntegrationLog) []*model.UserIntegrationLog {
	var tags []*model.UserIntegrationLog
	err := global.DB.Where(&params).Find(&tags).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return tags
}

func (impl userIntegrationLogServiceImpl) Count() int {
	var count int64
	global.DB.Model(&model.UserIntegrationLog{}).Count(&count)
	return int(count)
}

func (impl userIntegrationLogServiceImpl) ListPage(pageNum, pageSize int, params *model.UserIntegrationLog) *response.ApiResponse {
	var logs []*model.UserIntegrationLog
	err := global.DB.Where(&params).Scopes(model.Paginate(pageNum, pageSize)).Find(&logs).Error
	if err != nil {
		return nil
	}
	if err != nil {
		fmt.Println(err)
		return nil
	}
	var result []userResponse.IntegrationLog
	for _, item := range logs {
		userIntegration := UserIntegrationService.GetOne(&model.UserIntegration{
			Uid: item.Uid,
		})
		levelInfo := UserVipLevelService.GetOne(&model.UserVipLevel{
			Id: userIntegration.LevelId,
		})
		result = append(result, userResponse.IntegrationLog{
			UserIntegrationLog: item,
			LevelName:          levelInfo.LevelName,
			Level:              levelInfo.Level,
		})
	}
	var total int64
	global.DB.Model(&model.UserIntegrationLog{}).Where(&params).Count(&total)
	res := &response.ApiResponse{List: result, Total: int(total)}
	return res
}
