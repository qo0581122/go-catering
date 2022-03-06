package model

import (
	"catering/global"
	"fmt"
)

type IntegrationGiftLog struct {
	Id               int    `gorm:"primary_key" json:"id"`
	GiftId           int    `json:"gift_id"`
	IntegrationLogId int    `json:"integration_log_id"`
	Uid              int    `json:"uid"`
	CreatedTime      string `json:"created_time"`
}

func (IntegrationGiftLog) TableName() string {
	return "integration_gift_log"
}

func ListIntegrationGiftLogPage(pageNum, pageSize int, params *IntegrationGiftLog) ([]*IntegrationGiftLog, error) {
	var log []*IntegrationGiftLog
	err := global.DB.Where(&params).Scopes(Paginate(pageNum, pageSize)).Find(&log).Error
	if err != nil {
		return nil, err
	}
	return log, err
}

func ListIntegrationGiftLog(params *IntegrationGiftLog) []*IntegrationGiftLog {
	var log []*IntegrationGiftLog
	err := global.DB.Where(&params).Find(&log).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return log
}

func CountIntegrationGiftLog() int {
	var count int64
	global.DB.Model(&IntegrationGiftLog{}).Count(&count)
	return int(count)
}
