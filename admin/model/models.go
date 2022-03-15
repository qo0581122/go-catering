package model

import (
	"gorm.io/gorm"
)

type Model struct {
	ID          uint64         `gorm:"primary_key" json:"id"`
	CreatedAt   Time           `gorm:"column:created_time" json:"created_time"`
	UpdatedAt   Time           `gorm:"column:updated_time" json:"updated_time"`
	DeletedTime gorm.DeletedAt `gorm:"column:deleted_time;index" json:"-"`
}

// type GVA_MODEL struct {
// 	ID        uint           `gorm:"primarykey"` // 主键ID
// 	CreatedAt time.Time      // 创建时间
// 	UpdatedAt time.Time      // 更新时间
// 	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"` // 删除时间
// }

func Paginate(page, size int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page == 0 {
			page = 1
		}
		switch {
		case size > 100:
			size = 100
		case size <= 0:
			size = 10
		}

		offset := (page - 1) * size
		return db.Offset(offset).Limit(size)
	}
}
