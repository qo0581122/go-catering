package repository

import "database/sql"

type User struct {
	Id              int64        `db:"id"`
	Mobie           string       `db:"mobie"`
	Password        string       `db:"password"`
	Nickname        string       `db:"nickname"`
	Sex             int64        `db:"sex"`
	Platform        int64        `db:"platform"`
	Avatar          string       `db:"avatar"`
	Uid             int64        `db:"uid"`
	RegistryDid     string       `db:"registry_did"`
	RegistryTime    string       `db:"registry_time"`
	RecentLoginDid  string       `db:"recent_login_did"`
	RecentLoginTime string       `db:"recent_login_time"`
	UpdatedTime     string       `db:"updated_time"`
	DeletedTime     sql.NullTime `db:"deleted_time"`
}
