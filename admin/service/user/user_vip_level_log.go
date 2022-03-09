package user

import (
	"catering/global"
	"catering/model"
	"catering/model/common/response"
	userResponse "catering/model/user/response"
	"fmt"
)

var UserVipLevelLogService userVipLevelLogService = NewUserVipLevelLogService()

func NewUserVipLevelLogService() userVipLevelLogService {
	return userVipLevelLogServiceImpl{}
}

type userVipLevelLogServiceImpl struct {
}

func (impl userVipLevelLogServiceImpl) Add(params *model.UserVipLevelLog) error {
	return global.DB.Create(&params).Error
}

func (impl userVipLevelLogServiceImpl) GetOne(params *model.UserVipLevelLog) *model.UserVipLevelLog {
	var res model.UserVipLevelLog
	err := global.DB.Where(&params).First(&res).Error
	if err != nil {
		return nil
	}
	return &res
}

func (impl userVipLevelLogServiceImpl) List(params *model.UserVipLevelLog) []*model.UserVipLevelLog {
	var logs []*model.UserVipLevelLog
	err := global.DB.Where(&params).Find(&logs).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return logs
}

func (impl userVipLevelLogServiceImpl) Count() int {
	var count int64
	global.DB.Model(&model.UserVipLevelLog{}).Count(&count)
	return int(count)
}

func (impl userVipLevelLogServiceImpl) ListPage(pageNum, pageSize int, params *model.UserVipLevelLog) *response.ApiResponse {
	var logs []*model.UserVipLevelLog
	err := global.DB.Where(&params).Scopes(model.Paginate(pageNum, pageSize)).Find(&logs).Error
	if err != nil {
		return nil
	}
	var result []userResponse.VipLevelLog
	for _, item := range logs {
		before := UserVipLevelService.GetOne(&model.UserVipLevel{
			Id: item.BeforeLevelId,
		})
		after := UserVipLevelService.GetOne(&model.UserVipLevel{
			Id: item.AfterLevelId,
		})
		result = append(result, userResponse.VipLevelLog{
			UserVipLevelLog: item,
			BeforeLevelName: before.LevelName,
			BeforeLevel:     before.Level,
			AfterLevelName:  after.LevelName,
			AfterLevel:      after.Level,
		})
	}
	var total int64
	global.DB.Model(&model.UserVipLevelLog{}).Where(&params).Count(&total)
	res := &response.ApiResponse{List: result, Total: int(total)}
	return res
}
