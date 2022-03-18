package repository

import (
	. "catering/pkg/log"
	. "catering/proto/coupon"
	. "catering/proto/product"
	"database/sql"
	"time"

	"github.com/jmoiron/sqlx"
)

type CouponRepository interface {
	GetUserCoupons(uid int64) []UserCoupon
	GetUserCouponsByCid(uid, cid int64) []UserCoupon
	GetUserCouponsById(id int64) *UserCoupon
	GetCouponInfo(couponId int64) *Coupon
	GetConn() *sqlx.DB
	DecreseCouponCount(couponId int64, count int32, tx *sql.Tx) error
	InserUserCoupon(uid, couponId int64, useBeginTime, useEndTime string, getType int32, tx *sql.Tx) error
	UseUserCoupon(userCouponId int64, tx *sql.Tx) error
	InsertUserCouponLog(orderId, userCouponId int64, tx *sql.Tx) error
}

type CouponRepositoryImpl struct {
	Conn *sqlx.DB
}

func (impl *CouponRepositoryImpl) GetConn() *sqlx.DB {
	return impl.Conn
}

func (impl *CouponRepositoryImpl) GetUserCoupons(uid int64) []UserCoupon {
	SQL := "SELECT id, coupon_id, get_time, get_type, use_begin_time, use_end_time, use_status FROM user_coupon WHERE use_status = 1 AND uid = ? ORDER BY use_status ASC, use_begin_time ASC, use_end_time DESC, id DESC"
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
	var result []UserCoupon
	var id, couponId sql.NullInt64
	var useStatus, getType sql.NullInt32
	var getTime, useBeginTime, useEndTime sql.NullString
	for rows.Next() {
		rows.Scan(&id, &couponId, &getTime, &getType, &useBeginTime, &useEndTime, &useStatus)
		result = append(result, UserCoupon{
			Id:           id.Int64,
			GetTime:      getTime.String,
			GetType:      getType.Int32,
			UseBeginTime: useBeginTime.String,
			UseEndTime:   useEndTime.String,
			UseStatus:    useStatus.Int32,
			Coupon: &Coupon{
				Id: couponId.Int64,
			},
		})
	}
	return result
}

func (impl *CouponRepositoryImpl) GetUserCouponsById(id int64) *UserCoupon {
	SQL := "SELECT coupon_id, get_time, get_type, use_begin_time, use_end_time, use_status FROM user_coupon WHERE use_status = 1 AND id = ? ORDER BY use_status ASC, use_begin_time ASC, use_end_time DESC, id DESC"
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
	var couponId sql.NullInt64
	var useStatus, getType sql.NullInt32
	var getTime, useBeginTime, useEndTime sql.NullString
	for rows.Next() {
		rows.Scan(&id, &couponId, &getTime, &getType, &useBeginTime, &useEndTime, &useStatus)
		return &UserCoupon{
			GetTime:      getTime.String,
			GetType:      getType.Int32,
			UseBeginTime: useBeginTime.String,
			UseEndTime:   useEndTime.String,
			UseStatus:    useStatus.Int32,
			Coupon: &Coupon{
				Id: couponId.Int64,
			},
		}
	}
	return nil
}

func (impl *CouponRepositoryImpl) GetUserCouponsByCid(uid, cid int64) []UserCoupon {
	SQL := "SELECT id, get_time, get_type, use_begin_time, use_end_time, use_status FROM user_coupon WHERE use_status = 1 AND uid = ? AND coupon_id = ? ORDER BY use_status ASC, use_begin_time ASC, use_end_time DESC, id DESC"
	stmt, err := impl.Conn.Prepare(SQL)
	if err != nil {
		return nil
	}
	defer stmt.Close()
	rows, err := stmt.Query(uid, cid)
	if err != nil {
		return nil
	}
	defer rows.Close()
	var result []UserCoupon
	var id, couponId sql.NullInt64
	var useStatus, getType sql.NullInt32
	var getTime, useBeginTime, useEndTime sql.NullString
	for rows.Next() {
		rows.Scan(&id, &couponId, &getTime, &getType, &useBeginTime, &useEndTime, &useStatus)
		result = append(result, UserCoupon{
			Id:           id.Int64,
			GetTime:      getTime.String,
			GetType:      getType.Int32,
			UseBeginTime: useBeginTime.String,
			UseEndTime:   useEndTime.String,
			UseStatus:    useStatus.Int32,
		})
	}
	return result
}

func (impl *CouponRepositoryImpl) GetCouponInfo(couponId int64) *Coupon {
	SQL := "SELECT coupon_name, least_use_price, price, pic_url, get_type, use_type, use_begin_time, use_end_time, valid_time_type, valid_day, get_begin_time, get_end_time, get_count, description, product_id FROM coupon WHERE id = ? limit 1"
	stmt, err := impl.Conn.Prepare(SQL)
	if err != nil {
		return nil
	}
	defer stmt.Close()
	rows, err := stmt.Query(couponId)
	if err != nil {
		return nil
	}
	var id, productId sql.NullInt64
	var couponName, picUrl, useBeginTime, useEndTime, getBeginTime, getEndTime, description sql.NullString
	var leastUsePrice, price, getType, useType, validTimeType, validDay, getCount sql.NullInt32
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&id, &couponName, &leastUsePrice, &price, &picUrl, &getType, &useType, &useBeginTime, &useEndTime, &validTimeType, &validDay, &getBeginTime, &getEndTime, &getCount, &description, &productId)
		return &Coupon{
			Id:            id.Int64,
			CouponName:    couponName.String,
			LeastUsePrice: leastUsePrice.Int32,
			Price:         price.Int32,
			PicUrl:        picUrl.String,
			GetType:       getType.Int32,
			UseType:       useType.Int32,
			UseBeginTime:  useBeginTime.String,
			UseEndTime:    useEndTime.String,
			ValidTimeType: validTimeType.Int32,
			ValidDay:      validDay.Int32,
			GetBeginTime:  getBeginTime.String,
			GetEndTime:    getEndTime.String,
			GetCount:      getCount.Int32,
			Description:   description.String,
			Product: &ProductBase{
				Id: productId.Int64,
			},
		}
	}
	return nil
}

func (impl *CouponRepositoryImpl) InserUserCoupon(uid, couponId int64, useBeginTime, useEndTime string, getType int32, tx *sql.Tx) error {
	SQL := `INSERT INTO user_conpon (uid, coupon_id, get_time, use_begin_time, use_end_time, use_status, get_type) values(?,?,?,?,?,?,?) `
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
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	Logger.Debug(SQL, uid, couponId, currentTime, useBeginTime, useEndTime, 1, getType)
	res, err := stmt.Exec(uid, couponId, currentTime, useBeginTime, useEndTime, 1, getType)
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

func (impl *CouponRepositoryImpl) DecreseCouponCount(couponId int64, count int32, tx *sql.Tx) error {
	SQL := `UPDATE coupon SET reamin_count = remain_count - ? WHERE reamin_count >= ? AND id = ?`
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
	Logger.Debug(SQL, count, couponId)
	res, err := stmt.Exec(count, couponId)
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

func (impl *CouponRepositoryImpl) UseUserCoupon(userCouponId int64, tx *sql.Tx) error {
	SQL := "UPDATE user_coupon SET use_status = 2 WHERE id = ? AND use_status = 1"
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
	res, err := stmt.Exec(userCouponId)
	if err != nil {
		return err
	}
	rows, err := res.RowsAffected()
	if rows <= 0 {
		return err
	}
	return nil
}

func (impl *CouponRepositoryImpl) InsertUserCouponLog(orderId, userCouponId int64, tx *sql.Tx) error {
	SQL := `INSERT INTO user_coupon_log (order_id, user_coupon_id, created_time) values(?,?,?)`
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
	res, err := stmt.Exec(orderId, userCouponId, time.Now().Format("2006-01-02 15:03:04"))
	if err != nil {
		return err
	}
	rows, err := res.RowsAffected()
	if rows <= 0 {
		return err
	}
	return nil
}
