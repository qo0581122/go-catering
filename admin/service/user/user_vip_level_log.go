package user

import (
	"catering/model"
	"catering/model/common/response"
	"fmt"
)

var UserVipLevelLogService userVipLevelLogService = NewUserVipLevelLogService()

func NewUserVipLevelLogService() userVipLevelLogService {
	return userVipLevelLogServiceImpl{}
}

type userVipLevelLogServiceImpl struct {
}

func (impl userVipLevelLogServiceImpl) GetById(id uint64) *model.UserVipLevelLog {
	return model.GetUserVipLevelLogById(id)
}
func (impl userVipLevelLogServiceImpl) List(params *model.UserVipLevelLog) []*model.UserVipLevelLog {
	return model.ListUserVipLevelLog(params)
}

func (impl userVipLevelLogServiceImpl) Count() int {
	return model.CountUserAddress()
}

type UserVipLevelLogResponse struct {
	*model.UserVipLevelLog
	BeforeLevelName string `json:"before_level_name"`
	BeforeLevel     int    `json:"before_level"`
	AfterLevelName  string `json:"after_level_name"`
	AfterLevel      int    `json:"after_level"`
}

func (impl userVipLevelLogServiceImpl) ListPage(pageNum, pageSize int, params *model.UserVipLevelLog) *response.ApiResponse {
	logs, err := model.ListUserVipLevelLogPage(pageNum, pageSize, params)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	total := impl.Count()
	var result []UserVipLevelLogResponse
	for _, item := range logs {
		before := model.GetUserVipLevelById(item.BeforeLevelId)
		after := model.GetUserVipLevelById(item.AfterLevelId)
		result = append(result, UserVipLevelLogResponse{
			UserVipLevelLog: item,
			BeforeLevelName: before.LevelName,
			BeforeLevel:     before.Level,
			AfterLevelName:  after.LevelName,
			AfterLevel:      after.Level,
		})
	}
	res := &response.ApiResponse{List: result, Total: total}
	return res
}
