package model

import (
	"catering/global"
	"fmt"
)

type UserVipLevelLog struct {
	Id            uint64 `gorm:"primary_key" json:"id"`
	Uid           int    `json:"uid"`
	BeforeLevelId uint64 `json:"before_level_id"`
	AfterLevelId  uint64 `json:"after_level_id"`
	UpTime        Time   `json:"up_time"`
}

func (UserVipLevelLog) TableName() string {
	return "user_vip_level_log"
}

func AddUserVipLevelLog(params *UserVipLevelLog) error {
	return global.DB.Create(&params).Error
}

func UpdateUserVipLevelLog(params *UserVipLevelLog) error {
	return global.DB.Model(&UserVipLevelLog{}).Where("id = ?", params.Id).Updates(&params).Error
}

func DeleteUserVipLevelLog(id uint64) error {
	return global.DB.Delete(&UserVipLevelLog{}, id).Error
}

func ListUserVipLevelLogPage(pageNum, pageSize int, params *UserVipLevelLog) ([]*UserVipLevelLog, error) {
	var tags []*UserVipLevelLog
	err := global.DB.Where(&params).Scopes(Paginate(pageNum, pageSize)).Find(&tags).Error
	if err != nil {
		return nil, err
	}
	return tags, err
}

func ListUserVipLevelLog(params *UserVipLevelLog) []*UserVipLevelLog {
	var tags []*UserVipLevelLog
	err := global.DB.Where(&params).Find(&tags).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return tags
}

func GetUserVipLevelLogById(id uint64) *UserVipLevelLog {
	var res UserVipLevelLog
	err := global.DB.Where("id = ?", id).First(&res).Error
	if err != nil {
		return nil
	}
	return &res
}

func CountUserVipLevelLog() int {
	var count int64
	global.DB.Model(&UserVipLevelLog{}).Count(&count)
	return int(count)
}
