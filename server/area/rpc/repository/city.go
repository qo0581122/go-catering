package repository

import (
	pb "catering/proto/common"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type CityRepository interface {
	SelectByProvinceId() ([]*pb.City, error)
}

type CityRepositoryImpl struct {
	Conn *sqlx.DB
}

func (impl *CityRepositoryImpl) SelectByProvinceId(provinceId int64) ([]*pb.City, error) {
	SQL := "SELECT id, city_name FROM city WHERE province_id = ? AND status = 1"
	stmt, err := impl.Conn.Prepare(SQL)
	if err != nil {
		return nil, err
	}
	rows, err := stmt.Query(provinceId)
	if err != nil {
		return nil, err
	}
	var result []*pb.City
	var id sql.NullInt64
	var cityName sql.NullString
	for rows.Next() {
		rows.Scan(&id, &cityName)
		result = append(result, &pb.City{
			Id:       id.Int64,
			CityName: cityName.String,
		})
	}
	return result, nil
}
