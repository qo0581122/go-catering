package repository

import (
	pb "catering/proto/common"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type ProvinceRepository interface {
	SelectAll() ([]*pb.Province, error)
}

type ProvinceRepositoryImpl struct {
	Conn *sqlx.DB
}

func (impl *ProvinceRepositoryImpl) SelectAll() ([]*pb.Province, error) {
	SQL := "SELECT id, province_name FROM province where status = 1"
	stmt, err := impl.Conn.Prepare(SQL)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	var id sql.NullInt64
	var provinceName sql.NullString
	var result []*pb.Province
	for rows.Next() {
		rows.Scan(&id, &provinceName)
		result = append(result, &pb.Province{
			Id:           id.Int64,
			ProvinceName: provinceName.String,
		})
	}
	return result, nil
}
