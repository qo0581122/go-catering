package database

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Option func(db *gorm.DB)

func WithMaxIdleConns(idle int) Option {
	return func(db *gorm.DB) {
		if idle > 0 {
			db.DB().SetMaxIdleConns(idle)
		}
	}
}

func WithMaxOpenConns(open int) Option {
	return func(db *gorm.DB) {
		if open > 0 {
			db.DB().SetMaxOpenConns(open)
		}
	}
}

func WithConnMaxLifetime(t int) Option {
	return func(db *gorm.DB) {
		if t > 0 {
			db.DB().SetConnMaxLifetime(time.Duration(t * int(time.Second)))
		}
	}
}
