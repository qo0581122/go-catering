package repository

import (
	. "catering/proto/integration"
	"database/sql"
	"time"

	"github.com/jmoiron/sqlx"
)

type IntegrationRepository interface {
	GetConn() *sqlx.DB
	GetUserIntegration(uid int64) *UserIntegration
	GetUserVipLevel(uid int64) *UserVipLevel
	GetVipLevelInfo(id int64) *UserVipLevel
	GetUserIntegrationChangeLogs(uid int64) []UserIntegrationLog
	GetUserVipLevelUpLogs(uid int64) []UserVipUpLog
	InsertUserVipUpLog(uid, beforeLevelId, afterLevelId int64, tx *sql.Tx) error
	InsertUserIntegrationChangeLog(uid int64, changeType uint32, changeValue int64, beforeValue int64, afterValue int64, tx *sql.Tx) error
	DecreaseUserIntegration(uid, changeValue int64, tx *sql.Tx) error
	IncreaseUserIntegration(uid, changeValue, levelId int64, tx *sql.Tx) error
}

type IntegrationRepositoryImpl struct {
	Conn *sqlx.DB
}

func (impl *IntegrationRepositoryImpl) GetConn() *sqlx.DB {
	return impl.Conn
}

func (impl *IntegrationRepositoryImpl) GetUserIntegration(uid int64) *UserIntegration {
	SQL := "SELECT id, current_integration, history_integration, consume_integration, level_id, created_time, updated_time FROM user_integration WHERE uid = ? limit 1"

	stmt, err := impl.Conn.Prepare(SQL)
	if err != nil {
		return nil
	}
	defer stmt.Close()
	rows, err := stmt.Query(uid)
	if err != nil {
		return nil
	}
	defer rows.Close()
	var id, current, history, consume, levelId sql.NullInt64
	var updatedTime, createdTime sql.NullString
	for rows.Next() {
		rows.Scan(&id, &current, &history, &consume, &levelId, &createdTime, &updatedTime)
		return &UserIntegration{
			Id:                 id.Int64,
			Uid:                uid,
			CurrentIntegration: current.Int64,
			HistoryIntegration: history.Int64,
			ConsumeIntegration: consume.Int64,
			LevelId:            levelId.Int64,
			UpdatedTime:        updatedTime.String,
			CreatedTime:        createdTime.String,
		}
	}
	return nil
}

func (impl *IntegrationRepositoryImpl) GetUserVipLevel(uid int64) *UserVipLevel {
	SQL := "SELECT b.id, b.level_name, b.up_need_integration, b.level_discount, b.updated_time, b.next_level_id,  c.level_name as next_level_name FROM user_integration a INNER JOIN user_vip_level b ON a.level_id = b.id INNER JOIN user_vip_level c ON b.next_level_id = c.id  WHERE a.uid = ?"
	stmt, err := impl.Conn.Prepare(SQL)
	if err != nil {
		return nil
	}
	defer stmt.Close()
	rows, err := stmt.Query(uid)
	if err != nil {
		return nil
	}
	defer rows.Close()
	var id, upNeed, levelDiscount, nextLevelId sql.NullInt64
	var updatedTime, levelName, nextLevelName sql.NullString
	for rows.Next() {
		rows.Scan(&id, &levelName, &upNeed, &levelDiscount, &updatedTime, &nextLevelId, &nextLevelName)
		return &UserVipLevel{
			Id:                id.Int64,
			LevelName:         levelName.String,
			UpNeedIntegration: upNeed.Int64,
			LevelDiscount:     int32(levelDiscount.Int64),
			UpdatedTime:       updatedTime.String,
			NextLevelId:       nextLevelId.Int64,
			NextLevelName:     nextLevelName.String,
		}
	}
	return nil
}

func (impl *IntegrationRepositoryImpl) GetVipLevelInfo(id int64) *UserVipLevel {
	SQL := "SELECT b.level_name, b.up_need_integration, b.level_discount, b.updated_time, b.next_level_id,  c.level_name as next_level_name FROM user_vip_level b INNER JOIN user_vip_level c ON b.next_level_id = c.id  WHERE b.id = ?"
	stmt, err := impl.Conn.Prepare(SQL)
	if err != nil {
		return nil
	}
	defer stmt.Close()
	rows, err := stmt.Query(id)
	if err != nil {
		return nil
	}
	defer rows.Close()
	var upNeed, levelDiscount, nextLevelId sql.NullInt64
	var updatedTime, levelName, nextLevelName sql.NullString
	for rows.Next() {
		rows.Scan(&levelName, &upNeed, &levelDiscount, &updatedTime, &nextLevelId, &nextLevelName)
		return &UserVipLevel{
			Id:                id,
			LevelName:         levelName.String,
			UpNeedIntegration: upNeed.Int64,
			LevelDiscount:     int32(levelDiscount.Int64),
			UpdatedTime:       updatedTime.String,
			NextLevelId:       nextLevelId.Int64,
			NextLevelName:     nextLevelName.String,
		}
	}
	return nil
}

func (impl *IntegrationRepositoryImpl) GetUserIntegrationChangeLogs(uid int64) []UserIntegrationLog {
	SQL := "SELECT id, change_type, change_value, before_value, after_value, change_time FROM user_integration_log WHERE uid = ? ORDER BY change_time desc"
	stmt, err := impl.Conn.Prepare(SQL)
	if err != nil {
		return nil
	}
	defer stmt.Close()
	rows, err := stmt.Query(uid)
	if err != nil {
		return nil
	}
	defer rows.Close()
	var (
		result                                     []UserIntegrationLog
		id, changeType, changeValue, before, after sql.NullInt64
		changeTime                                 sql.NullString
	)
	for rows.Next() {
		rows.Scan(&id, &changeType, &changeValue, &before, &after, &changeTime)
		result = append(result, UserIntegrationLog{
			Id:          id.Int64,
			ChangeType:  changeType.Int64,
			ChangeValue: changeValue.Int64,
			BeforeValue: before.Int64,
			AfterValue:  after.Int64,
			ChangeTime:  changeTime.String,
		})
	}
	return result
}

func (impl *IntegrationRepositoryImpl) GetUserVipLevelUpLogs(uid int64) []UserVipUpLog {
	SQL := "SELECT id, before_level_id, after_level_id, up_time FROM user_vip_log WHERE uid = ? ORDER BY up_time desc"
	stmt, err := impl.Conn.Prepare(SQL)
	if err != nil {
		return nil
	}
	defer stmt.Close()
	rows, err := stmt.Query(uid)
	if err != nil {
		return nil
	}
	defer rows.Close()
	var (
		result            []UserVipUpLog
		id, before, after sql.NullInt64
		upTime            sql.NullString
	)
	for rows.Next() {
		rows.Scan(&id, &before, &after, &before, &upTime)
		result = append(result, UserVipUpLog{
			Id:            id.Int64,
			Uid:           uid,
			BeforeLevelId: before.Int64,
			AfterLevelId:  after.Int64,
			UpTime:        upTime.String,
		})
	}
	return result
}

func (impl *IntegrationRepositoryImpl) InsertUserIntegrationChangeLog(uid int64, changeType uint32, changeValue int64, beforeValue int64, afterValue int64, tx *sql.Tx) error {
	SQL := "INSERT INTO user_integration_log (uid, change_type, change_value, before_value, after_value, change_time) values(?,?,?,?,?,?)"
	var (
		stmt *sql.Stmt
		err  error
	)
	if tx != nil {
		stmt, err = tx.Prepare(SQL)
	} else {
		stmt, err = impl.Conn.Prepare(SQL)
	}
	if err != nil {
		return err
	}
	defer stmt.Close()
	updatedTime := time.Now().Format("2006-01-02 15:03:04")
	rs, err := stmt.Exec(uid, changeType, changeValue, beforeValue, afterValue, updatedTime)
	if err != nil {
		return err
	}
	rows, err := rs.RowsAffected()
	if err != nil || rows <= 0 {
		return err
	}
	return nil
}

func (impl *IntegrationRepositoryImpl) InsertUserVipUpLog(uid, beforeLevelId, afterLevelId int64, tx *sql.Tx) error {
	SQL := "INSERT INTO user_vip_log (uid, before_level_id, after_level_id, up_time) values(?,?,?,?)"
	var (
		stmt *sql.Stmt
		err  error
	)
	if tx != nil {
		stmt, err = tx.Prepare(SQL)
	} else {
		stmt, err = impl.Conn.Prepare(SQL)
	}
	if err != nil {
		return err
	}
	defer stmt.Close()
	updatedTime := time.Now().Format("2006-01-02 15:03:04")
	rs, err := stmt.Exec(uid, beforeLevelId, afterLevelId, updatedTime)
	if err != nil {
		return err
	}
	rows, err := rs.RowsAffected()
	if err != nil || rows <= 0 {
		return err
	}
	return nil
}

func (impl *IntegrationRepositoryImpl) IncreaseUserIntegration(uid, changeValue, levelId int64, tx *sql.Tx) error {
	SQL := "UPDATE user_integration SET current_integration = current_integration + ?, history_integration = history_integration + ?, level_id = ? WHERE uid = ?"
	var (
		stmt *sql.Stmt
		err  error
	)
	if tx != nil {
		stmt, err = tx.Prepare(SQL)
	} else {
		stmt, err = impl.Conn.Prepare(SQL)
	}
	if err != nil {
		return err
	}
	defer stmt.Close()
	rs, err := stmt.Exec(changeValue, changeValue, levelId, uid)
	if err != nil {
		return err
	}
	rows, err := rs.RowsAffected()
	if err != nil || rows <= 0 {
		return err
	}
	return nil
}

func (impl *IntegrationRepositoryImpl) DecreaseUserIntegration(uid, changeValue int64, tx *sql.Tx) error {
	SQL := "UPDATE user_integration SET current_integration = current_integration - ?, consume_integration = consume_integration + ? WHERE uid = ? AND current_integration >= ?"
	var (
		stmt *sql.Stmt
		err  error
	)
	if tx != nil {
		stmt, err = tx.Prepare(SQL)
	} else {
		stmt, err = impl.Conn.Prepare(SQL)
	}
	if err != nil {
		return err
	}
	defer stmt.Close()
	rs, err := stmt.Exec(changeValue, changeValue, uid, changeValue)
	if err != nil {
		return err
	}
	rows, err := rs.RowsAffected()
	if err != nil || rows <= 0 {
		return err
	}
	return nil
}
