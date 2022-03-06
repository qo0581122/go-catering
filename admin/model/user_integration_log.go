package model

import (
	"catering/global"
	"fmt"
)

type UserIntegrationLog struct {
	Id          uint64 `gorm:"primary_key" json:"id"`
	Uid         int    `json:"uid"`
	ChangeType  int    `json:"change_type"`
	ChangeValue int    `json:"change_value"`
	BeforeValue int    `json:"before_value"`
	AfterValue  int    `json:"after_value"`
	ChangeTime  Time   `json:"change_time"`
}

func (UserIntegrationLog) TableName() string {
	return "user_integration_log"
}

func AddUserIntegrationLog(params *UserIntegrationLog) error {
	return global.DB.Create(&params).Error
}

func UpdateUserIntegrationLog(params *UserIntegrationLog) error {
	return global.DB.Model(&UserIntegrationLog{}).Where("id = ?", params.Id).Updates(&params).Error
}

func DeleteUserIntegrationLog(id uint64) error {
	return global.DB.Delete(&UserIntegrationLog{}, id).Error
}

func ListUserIntegrationLogPage(pageNum, pageSize int, params *UserIntegrationLog) ([]*UserIntegrationLog, error) {
	var tags []*UserIntegrationLog
	err := global.DB.Where(&params).Scopes(Paginate(pageNum, pageSize)).Find(&tags).Error
	if err != nil {
		return nil, err
	}
	return tags, err
}

func ListUserIntegrationLog(params *UserIntegrationLog) []*UserIntegrationLog {
	var tags []*UserIntegrationLog
	err := global.DB.Where(&params).Find(&tags).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return tags
}

func GetUserIntegrationLogById(id uint64) *UserIntegrationLog {
	var res UserIntegrationLog
	err := global.DB.Where("id = ?", id).First(&res).Error
	if err != nil {
		return nil
	}
	return &res
}

func CountUserIntegrationLog() int {
	var count int64
	global.DB.Model(&UserIntegrationLog{}).Count(&count)
	return int(count)
}
