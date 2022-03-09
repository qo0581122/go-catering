package model

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
