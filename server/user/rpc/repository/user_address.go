package repository

import (
	pb "catering/proto/user"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type UserAddressRepository interface {
	SelectByUserId(userId int64) []*pb.UserAddress
	SelectAddressTag(addressId int64) []*pb.AddressTag
}

type UserAddressRepositoryImpl struct {
	Conn *sqlx.DB
}

func (impl *UserAddressRepositoryImpl) SelectByUserId(userId int64) []*pb.UserAddress {
	SQL := "SELECT id, receive_name, receive_sex, receive_phone, province_id, city_id, distinct_id, detail_address, is_default, deleted_time FROM user_address WHERE uid = ? order by is_default desc, sort desc"
	stmt, err := impl.Conn.Prepare(SQL)
	if err != nil {
		return nil
	}
	defer stmt.Close()
	rows, err := stmt.Query(userId)
	defer rows.Close()
	if err != nil {
		return nil
	}
	var result []*pb.UserAddress
	var id, distinctId, provinceId, cityId, isDefault, receiveSex sql.NullInt64
	var receiveName, receivePhone, detailAddress, deletedTime sql.NullString
	for rows.Next() {
		rows.Scan(&id, &receiveName, &receiveSex, &receivePhone, &provinceId, &cityId, &distinctId, &detailAddress, &isDefault, &deletedTime)
		result = append(result, &pb.UserAddress{
			Id:            id.Int64,
			ReceiveName:   receiveName.String,
			ReceiveSex:    receiveSex.Int64,
			ReceivePhone:  receivePhone.String,
			ProvinceId:    provinceId.Int64,
			CityId:        cityId.Int64,
			DistinctId:    distinctId.Int64,
			DetailAddress: detailAddress.String,
			IsDefault:     int32(isDefault.Int64),
			DeletedTime:   deletedTime.String,
		})
	}
	return result
}
func (impl *UserAddressRepositoryImpl) SelectAddressTag(addressId int64) []*pb.AddressTag {
	SQL := "SELECT a.id, a.tag_name, a.sort FROM address_tag a LEFT JOIN address_tag_relation b ON a.id = b.tag_id WHERE b.address_id = ? ORDER BY sort desc"
	stmt, err := impl.Conn.Prepare(SQL)
	if err != nil {
		return nil
	}
	defer stmt.Close()
	rows, err := stmt.Query(addressId)
	defer rows.Close()
	if err != nil {
		return nil
	}
	var result []*pb.AddressTag
	var id, sort sql.NullInt64
	var tagName sql.NullString
	for rows.Next() {
		rows.Scan(&id, &tagName, &sort)
		result = append(result, &pb.AddressTag{
			Id:      id.Int64,
			TagName: tagName.String,
			Sort:    int32(sort.Int64),
		})
	}
	return result
}
