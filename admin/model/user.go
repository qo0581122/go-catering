package model

import (
	"catering/global"
	"fmt"
)

type User struct {
	Id              uint64 `gorm:"primary_key" json:"id"`
	Mobie           string `json:"mobie"`
	Nickname        string `json:"nickname"`
	Sex             uint32 `json:"sex"`
	Platform        uint32 `json:"platform"`
	Avatar          string `json:"avatar"`
	RegistryDID     string `json:"registry_did"`
	RegistryTime    Time   `json:"registry_time"`
	RecentLoginDID  string `json:"recent_login_did"`
	RecentLoginTime Time   `json:"registry_login_time"`
	ChannelId       uint64 `json:"channel_id"`
	UpdatedTime     Time   `json:"updated_time"`
}

func (User) TableName() string {
	return "user"
}
func AddUser(params *User) error {
	return global.DB.Create(&params).Error
}

func UpdateUser(params *User) error {
	return global.DB.Model(&User{}).Where("id = ?", params.Id).Updates(&params).Error
}

func DeleteUser(id uint64) error {
	return global.DB.Delete(&User{}, id).Error
}

func ListUserPage(pageNum, pageSize int, params *User) ([]*User, error) {
	var tags []*User
	err := global.DB.Where(&params).Scopes(Paginate(pageNum, pageSize)).Find(&tags).Error
	if err != nil {
		return nil, err
	}
	return tags, err
}

func ListUser(params *User) []*User {
	var tags []*User
	err := global.DB.Where(&params).Find(&tags).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return tags
}

func GetUserById(id uint64) *User {
	var res User
	err := global.DB.Where("id = ?", id).First(&res).Error
	if err != nil {
		return nil
	}
	return &res
}

func CountUser() int {
	var count int64
	global.DB.Model(&User{}).Count(&count)
	return int(count)
}
