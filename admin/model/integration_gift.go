package model

import (
	"catering/global"
	"fmt"
)

type IntegrationGift struct {
	Model
	GiftName        string `json:"gift_name"`
	GiftType        int    `json:"gift_type"`
	Description     string `json:"description"`
	OtherId         int    `json:"other_id"`
	CostIntegration int    `json:"cost_integraion"`
	TotalCount      int    `json:"total_count"`
	RemainCount     int    `json:"remain_count"`
	BeginTime       Time   `json:"begin_time"`
	EndTime         Time   `json:"end_time"`
	Status          int    `json:"status"`
	Sort            int    `json:"sort"`
}

func (IntegrationGift) TableName() string {
	return "integration_gift"
}

func AddIntegrationGift(params *IntegrationGift) error {
	return global.DB.Create(&params).Error
}

func UpdateIntegrationGift(params *IntegrationGift) error {
	return global.DB.Model(&IntegrationGift{}).Where("id = ?", params.ID).Updates(&params).Error
}

func DeleteIntegrationGift(id uint64) error {
	return global.DB.Delete(&IntegrationGift{}, id).Error
}

func ListIntegrationGiftPage(pageNum, pageSize int, params *IntegrationGift) ([]*IntegrationGift, error) {
	var gift []*IntegrationGift
	err := global.DB.Where(&params).Scopes(Paginate(pageNum, pageSize)).Find(&gift).Error
	if err != nil {
		return nil, err
	}
	return gift, err
}

func ListIntegrationGift(params *IntegrationGift) []*IntegrationGift {
	var gift []*IntegrationGift
	err := global.DB.Where(&params).Find(&gift).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return gift
}

func CountIntegrationGift() int {
	var count int64
	global.DB.Model(&IntegrationGift{}).Count(&count)
	return int(count)
}
