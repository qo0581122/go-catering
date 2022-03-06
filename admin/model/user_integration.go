package model

import (
	"catering/global"
	"fmt"

	"gorm.io/gorm"
)

type UserIntegration struct {
	Id                 uint64 `gorm:"primary_key" json:"id"`
	Uid                int    `json:"uid"`
	CurrentIntegration int    `json:"current_integration"`
	HistoryIntegration int    `json:"history_integration"`
	ConsumeIntegration int    `json:"consume_integration"`
	LevelId            uint64 `json:"level_id"`
	CreatedTime        Time   `json:"created_time"`
	UpdatedTime        Time   `json:"updated_time"`
}

func (UserIntegration) TableName() string {
	return "user_integration"
}
func AddUserIntegration(params *UserIntegration) error {
	return global.DB.Create(&params).Error
}

func UpdateUserIntegration(params *UserIntegration) error {
	return global.DB.Model(&UserIntegration{}).Where("id = ?", params.Id).Updates(&params).Error
}

func DeleteUserIntegration(id uint64) error {
	return global.DB.Delete(&UserIntegration{}, id).Error
}

func ListUserIntegrationPage(pageNum, pageSize int, params *UserIntegration) ([]*UserIntegration, error) {
	var tags []*UserIntegration
	err := global.DB.Where(&params).Scopes(Paginate(pageNum, pageSize)).Find(&tags).Error
	if err != nil {
		return nil, err
	}
	return tags, err
}

func ListUserIntegration(params *UserIntegration) []*UserIntegration {
	var tags []*UserIntegration
	err := global.DB.Where(&params).Find(&tags).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return tags
}

func GetUserIntegrationById(id uint64) *UserIntegration {
	var res UserIntegration
	err := global.DB.Where("id = ?", id).First(&res).Error
	if err != nil {
		return nil
	}
	return &res
}

func GetUserIntegrationByUin(uin uint64) *UserIntegration {
	var res UserIntegration
	err := global.DB.Where("uin = ?", uin).First(&res).Error
	if err != nil {
		return nil
	}
	return &res
}

func CountUserIntegration() int {
	var count int64
	global.DB.Model(&UserIntegration{}).Count(&count)
	return int(count)
}

func IncreaseIntegration(tx *gorm.DB, id uint64, changeValue int) error {
	SQL := "UPDATE user_integration SET current_integration = current_integration + ?, history_integration = history_integration + ?  WHERE id = ?"
	db := global.DB
	if tx != nil {
		db = tx
	}
	return db.Exec(SQL, changeValue, id).Error
}

func DecreaseIntegration(tx *gorm.DB, id uint64, changeValue int) error {
	SQL := "UPDATE user_integration SET current_integration = current_integration - ? WHERE id = ? AND current_integration >= ?"
	db := global.DB
	if tx != nil {
		db = tx
	}
	return db.Exec(SQL, changeValue, id, changeValue).Error
}
