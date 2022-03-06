package user

import (
	"catering/model"
	"catering/model/common/response"
	"fmt"
)

var UserIntegrationLogService = NewUserIntegrationLogService()

func NewUserIntegrationLogService() userIntegrationLogService {
	return userIntegrationLogServiceImpl{}
}

type userIntegrationLogServiceImpl struct {
}

func (impl userIntegrationLogServiceImpl) GetById(id uint64) *model.UserIntegrationLog {
	return model.GetUserIntegrationLogById(id)
}
func (impl userIntegrationLogServiceImpl) List(params *model.UserIntegrationLog) []*model.UserIntegrationLog {
	return model.ListUserIntegrationLog(params)
}

func (impl userIntegrationLogServiceImpl) Count() int {
	return model.CountUserAddress()
}

type UserIntegrationLogResponse struct {
	*model.UserIntegrationLog
	LevelName string `json:"level_name"`
	Level     int    `json:"level"`
}

func (impl userIntegrationLogServiceImpl) ListPage(pageNum, pageSize int, params *model.UserIntegrationLog) *response.ApiResponse {
	logs, err := model.ListUserIntegrationLogPage(pageNum, pageSize, params)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	var result []UserIntegrationLogResponse
	vipLevelService := NewUserVipLevelService()
	for _, item := range logs {
		levelInfo := vipLevelService.GetByUin(uint64(item.Uid))
		result = append(result, UserIntegrationLogResponse{
			UserIntegrationLog: item,
			LevelName:          levelInfo.LevelName,
			Level:              levelInfo.Level,
		})
	}
	total := impl.Count()
	res := &response.ApiResponse{List: result, Total: total}
	return res
}
