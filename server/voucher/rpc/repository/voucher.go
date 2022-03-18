package repository

import (
	. "catering/pkg/utils"
	. "catering/proto/voucher"
	"database/sql"
	"time"

	. "catering/pkg/log"

	"github.com/jmoiron/sqlx"
)

type VoucherRepository interface {
	Select(opts ...Options) []Voucher
	GetUserVoucher(opts ...Options) []UserVoucher
	InsertUserVoucher(uid, voucher_id int64, useBeginTime, useEndTime string, tx *sql.Tx) error
	GetConn() *sqlx.DB
	UseUserVoucher(userVoucherId int64, tx *sql.Tx) error
	InsertUserVoucherLog(orderId, userVoucherId int64, tx *sql.Tx) error
	DecreseVoucherCount(voucher_id int64, count int32, tx *sql.Tx) error
}

type VoucherRepositoryImpl struct {
	Conn *sqlx.DB
}

func (impl *VoucherRepositoryImpl) GetConn() *sqlx.DB {
	return impl.Conn
}

func (impl *VoucherRepositoryImpl) Select(opts ...Options) []Voucher {
	SQL := "SELECT id, voucher_name, least_use_price, price, pic_url, total_count, remain_count, use_type, get_type, use_begin_time, use_end_time, valid_time_type, valid_day, get_begin_time, get_end_time, get_count, description, status FROM voucher WHERE 1 = 1"
	var queryParam []interface{}
	for _, opt := range opts {
		var param interface{}
		SQL, param = opt(SQL)
		if p, ok := param.(string); ok && len(p) == 0 {
			continue
		}
		queryParam = append(queryParam, param)
	}
	stmt, err := impl.Conn.Prepare(SQL)
	Logger.Info(SQL, queryParam)
	if err != nil {
		return nil
	}
	defer stmt.Close()
	rows, err := stmt.Query(queryParam...)
	if err != nil {
		return nil
	}
	defer rows.Close()
	var result []Voucher
	var id sql.NullInt64
	var voucherName, picUrl, useBeginTime, useEndTime, getBeginTime, getEndTime, description sql.NullString
	var LeastUsePrice, price, totalCount, remainCount, useType, getType, validTimeType, validDay, getCount, status sql.NullInt32
	for rows.Next() {
		rows.Scan(&id, &voucherName, &LeastUsePrice, &price, &picUrl, &totalCount, &remainCount, &useType, &getType, &useBeginTime, &useEndTime, &validTimeType, &validDay, &getBeginTime, &getEndTime, &getCount, &description, &status)
		result = append(result, Voucher{
			Id:            id.Int64,
			VoucherName:   voucherName.String,
			LeastUsePrice: LeastUsePrice.Int32,
			Price:         price.Int32,
			PicUrl:        picUrl.String,
			TotalCount:    totalCount.Int32,
			RemainCount:   remainCount.Int32,
			UseType:       useType.Int32,
			GetType:       getType.Int32,
			UseBeginTime:  useBeginTime.String,
			UseEndTime:    useEndTime.String,
			ValidTimeType: validTimeType.Int32,
			ValidDay:      validDay.Int32,
			GetBeginTime:  getBeginTime.String,
			GetEndTime:    getEndTime.String,
			GetCount:      getCount.Int32,
			Description:   description.String,
			Status:        status.Int32,
		})
	}
	return result
}

func (impl *VoucherRepositoryImpl) GetUserVoucher(opts ...Options) []UserVoucher {
	SQL := "SELECT a.id as user_voucher_id, a.get_time, a.use_begin_time, a.use_end_time, a.use_status, b.id, b.voucher_name, b.least_use_price, b.price, b.pic_url, b.total_count, b.remain_count, b.use_type, b.get_type, b.use_begin_time, b.use_end_time, b.valid_time_type, b.valid_day, b.get_begin_time, b.get_end_time, b.get_count, b.description, b.status FROM user_voucher a, voucher b WHERE a.voucher_id = b.id "
	ORDER := " ORDER BY a.use_begin_time desc, a.get_time desc, a.id desc"
	var queryParam []interface{}
	for _, opt := range opts {
		var param interface{}
		SQL, param = opt(SQL)
		if p, ok := param.(string); ok && len(p) == 0 {
			continue
		}
		queryParam = append(queryParam, param)
	}
	SQL += ORDER
	stmt, err := impl.Conn.Prepare(SQL)
	if err != nil {
		return nil
	}
	defer stmt.Close()
	rows, err := stmt.Query(queryParam...)
	if err != nil {
		return nil
	}
	defer rows.Close()
	Logger.Info(SQL, queryParam)
	var result []UserVoucher
	var id, userVocuherId sql.NullInt64
	var voucherName, picUrl, useBeginTime, useEndTime, getBeginTime, getEndTime, description, userGetTime, userUseBeginTime, userUseEndTime sql.NullString
	var LeastUsePrice, price, totalCount, remainCount, useType, getType, validTimeType, validDay, getCount, status, userUseStatus sql.NullInt32
	for rows.Next() {
		err = rows.Scan(&userVocuherId, &userGetTime, &userUseBeginTime, &userUseEndTime, &userUseStatus, &id, &voucherName, &LeastUsePrice, &price, &picUrl, &totalCount, &remainCount, &useType, &getType, &useBeginTime, &useEndTime, &validTimeType, &validDay, &getBeginTime, &getEndTime, &getCount, &description, &status)
		if err != nil {
			Logger.Info()
			return result
		}
		result = append(result, UserVoucher{
			UserVoucherId: userVocuherId.Int64,
			GetTime:       userGetTime.String,
			UseBeginTime:  useBeginTime.String,
			UseEndTime:    useEndTime.String,
			UseStatus:     userUseStatus.Int32,
			Voucher: &Voucher{
				Id:            id.Int64,
				VoucherName:   voucherName.String,
				LeastUsePrice: LeastUsePrice.Int32,
				Price:         price.Int32,
				PicUrl:        picUrl.String,
				TotalCount:    totalCount.Int32,
				RemainCount:   remainCount.Int32,
				UseType:       useType.Int32,
				GetType:       getType.Int32,
				UseBeginTime:  useBeginTime.String,
				UseEndTime:    useEndTime.String,
				ValidTimeType: validTimeType.Int32,
				ValidDay:      validDay.Int32,
				GetBeginTime:  getBeginTime.String,
				GetEndTime:    getEndTime.String,
				GetCount:      getCount.Int32,
				Description:   description.String,
				Status:        status.Int32,
			},
		})
	}
	return result
}

func (impl *VoucherRepositoryImpl) InsertUserVoucher(uid, voucher_id int64, useBeginTime, useEndTime string, tx *sql.Tx) error {
	SQL := "INSERT INTO user_voucher(uid, voucher_id, get_time, use_begin_time, use_end_time, use_status) values(?,?,?,?,?,?)"
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
	if err != nil {
		return err
	}
	defer stmt.Close()
	time := GetCurrentTime()
	Logger.Info(SQL, uid, voucher_id, time, useBeginTime, useEndTime)
	res, err := stmt.Exec(uid, voucher_id, time, useBeginTime, useEndTime, 1)
	if err != nil {
		return err
	}
	rows, err := res.RowsAffected()
	if rows <= 0 {
		return err
	}
	return nil
}

func (impl *VoucherRepositoryImpl) DecreseVoucherCount(voucher_id int64, count int32, tx *sql.Tx) error {
	SQL := "UPDATE voucher SET remain_count = remain_count - ? WHERE remain_count >= ? AND id = ?"
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
	Logger.Info(SQL, voucher_id, count)
	defer stmt.Close()
	res, err := stmt.Exec(count, count, voucher_id)
	if err != nil {
		return err
	}
	rows, err := res.RowsAffected()
	if rows <= 0 {
		return err
	}
	return nil
}

func (impl *VoucherRepositoryImpl) UseUserVoucher(userVoucherId int64, tx *sql.Tx) error {
	SQL := "UPDATE user_voucher SET use_status = 2 WHERE id = ? AND use_status = 1"
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
	Logger.Info(SQL, userVoucherId)
	defer stmt.Close()
	res, err := stmt.Exec(userVoucherId)
	if err != nil {
		return err
	}
	rows, err := res.RowsAffected()
	if rows <= 0 {
		return err
	}
	return nil
}

func (impl *VoucherRepositoryImpl) InsertUserVoucherLog(orderId, userVoucherId int64, tx *sql.Tx) error {
	SQL := `INSERT INTO user_voucher_log (order_id, user_voucher_id, created_time) values(?,?,?)`
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
	Logger.Info(SQL, orderId, userVoucherId)
	res, err := stmt.Exec(orderId, userVoucherId, time.Now().Format("2006-01-02 15:03:04"))
	if err != nil {
		return err
	}
	rows, err := res.RowsAffected()
	if rows <= 0 {
		return err
	}
	return nil
}
