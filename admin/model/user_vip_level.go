package model

import (
	"catering/global"
	"fmt"

	"gorm.io/gorm"
)

type UserVipLevel struct {
	Id                uint64 `gorm:"primary_key" json:"id"`
	LevelName         string `json:"level_name"`
	Level             int    `json:"level"`
	UpNeedIntegration int    `json:"up_need_integration"`
	LevelDiscount     int    `json:"level_discount"`
	NextLevelId       uint64 `json:"next_level_id"`
	CreatedAt         Time   `gorm:"column:created_time" json:"created_time"`
	UpdatedAt         Time   `gorm:"column:updated_time" json:"updated_time"`
}

func (UserVipLevel) TableName() string {
	return "user_vip_level"
}

func AddUserVipLevel(tx *gorm.DB, params *UserVipLevel) (*UserVipLevel, error) {
	if tx != nil {
		return params, tx.Create(&params).Error
	}
	return params, global.DB.Create(&params).Error
}

func UpdateUserVipLevel(tx *gorm.DB, params *UserVipLevel) error {
	if tx != nil {
		return tx.Model(&UserVipLevel{}).Where("id = ?", params.Id).Updates(&params).Error
	}
	return global.DB.Model(&UserVipLevel{}).Where("id = ?", params.Id).Updates(&params).Error
}

func DeleteUserVipLevel(id uint64) error {
	return global.DB.Delete(&UserVipLevel{}, id).Error
}

func ListUserVipLevelPage(pageNum, pageSize int, params *UserVipLevel) ([]*UserVipLevel, error) {
	var tags []*UserVipLevel
	err := global.DB.Where(&params).Scopes(Paginate(pageNum, pageSize)).Find(&tags).Error
	if err != nil {
		return nil, err
	}
	return tags, err
}

func ListUserVipLevel(params *UserVipLevel) []*UserVipLevel {
	var tags []*UserVipLevel
	err := global.DB.Where(&params).Find(&tags).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return tags
}

func GetUserVipLevelById(id uint64) *UserVipLevel {
	var res UserVipLevel
	err := global.DB.Where("id = ?", id).First(&res).Error
	if err != nil {
		return nil
	}
	return &res
}

func GetUserVipLevelByUin(uin uint64) *UserVipLevel {
	var vipInfo *UserVipLevel
	SQL := "SELECT b.* FROM user_integration a, user_vip_level b WHERE a.level_id = b.id AND a.uid = ?"
	global.DB.Raw(SQL, uin).Scan(&vipInfo)
	return vipInfo
}

func CountUserVipLevel() int {
	var count int64
	global.DB.Model(&UserVipLevel{}).Count(&count)
	return int(count)
}
