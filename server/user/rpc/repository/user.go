package repository

import (
	. "catering/pkg/log"
	pb "catering/proto/user"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type UserRepository interface {
	FindOneByMobie(mobie string) (*pb.UserInfo, error)
	Insert(user User, tx *sql.Tx) error
	FindOne(uid int64) (*pb.UserInfo, error)
}

type UserRepositoryImpl struct {
	Conn *sqlx.DB
}

func (impl *UserRepositoryImpl) FindOneByMobie(mobie string) (*pb.UserInfo, error) {
	SQL := "SELECT password, nickname, sex, platform, avatar, uid, registry_did, registry_time, recent_login_did, recent_login_time FROM user where mobie = ? limit 1"
	stmt, err := impl.Conn.Prepare(SQL)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query(mobie)
	Logger.Debug(SQL, mobie)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var result *pb.UserInfo
	var uid sql.NullInt64
	var password, nickname, avatar, registry_did, registry_time, recent_login_did, recent_login_time sql.NullString
	var sex, platform sql.NullInt32
	for rows.Next() {
		rows.Scan(&password, &nickname, &sex, &platform, &avatar, &uid, &registry_did, &registry_time, &recent_login_did, &recent_login_time)
		result = &pb.UserInfo{
			Mobie:           mobie,
			Password:        password.String,
			Nickname:        nickname.String,
			Sex:             sex.Int32,
			Platform:        platform.Int32,
			Avatar:          avatar.String,
			Uid:             uid.Int64,
			RegistryDid:     registry_did.String,
			RegistryTime:    registry_time.String,
			RecentLoginDid:  recent_login_did.String,
			RecentLoginTime: recent_login_time.String,
		}
		return result, nil
	}
	return nil, sql.ErrNoRows
}

func (impl *UserRepositoryImpl) FindOne(uid int64) (*pb.UserInfo, error) {
	SQL := "SELECT mobie, password, nickname, sex, platform, avatar, registry_did, registry_time, recent_login_did, recent_login_time FROM user where uid = ? limit 1"
	stmt, err := impl.Conn.Prepare(SQL)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query(uid)
	Logger.Debug(SQL, uid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var result *pb.UserInfo
	var mobie, password, nickname, avatar, registry_did, registry_time, recent_login_did, recent_login_time sql.NullString
	var sex, platform sql.NullInt32
	for rows.Next() {
		rows.Scan(&mobie, &password, &nickname, &sex, &platform, &avatar, &uid, &registry_did, &registry_time, &recent_login_did, &recent_login_time)
		result = &pb.UserInfo{
			Mobie:           mobie.String,
			Password:        password.String,
			Nickname:        nickname.String,
			Sex:             sex.Int32,
			Platform:        platform.Int32,
			Avatar:          avatar.String,
			Uid:             uid,
			RegistryDid:     registry_did.String,
			RegistryTime:    registry_time.String,
			RecentLoginDid:  recent_login_did.String,
			RecentLoginTime: recent_login_time.String,
		}
		return result, nil
	}
	return nil, sql.ErrNoRows
}

func (impl *UserRepositoryImpl) Insert(user User, tx *sql.Tx) error {
	SQL := `INSERT INTO user (mobie, password, nickname, sex, platform, avatar, uid, registry_did, registry_time, recent_login_did, recent_login_time, updated_time) values(?,?,?,?,?,?,?,?,?,?,?,?)`
	var stmt *sql.Stmt
	var err error
	if tx != nil {
		stmt, err = tx.Prepare(SQL)
		if err != nil {
			return err
		}
	} else {
		stmt, err = impl.Conn.Prepare(SQL)
		if err != nil {
			return err
		}
	}
	defer stmt.Close()
	Logger.Debug(SQL, user.Mobie, user.Password, user.Nickname, user.Sex, user.Platform, user.Avatar, user.Uid, user.RegistryDid, user.RegistryTime, user.RecentLoginDid, user.RecentLoginTime, user.UpdatedTime)
	res, err := stmt.Exec(user.Mobie, user.Password, user.Nickname, user.Sex, user.Platform, user.Avatar, user.Uid, user.RegistryDid, user.RegistryTime, user.RecentLoginDid, user.RecentLoginTime, user.UpdatedTime)
	if err != nil {
		return err
	}
	rows, err := res.RowsAffected()
	if rows <= 0 {
		return sql.ErrNoRows
	}
	if err != nil {
		return err
	}
	return nil
}
