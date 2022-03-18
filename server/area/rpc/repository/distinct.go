package repository

import (
	pb "catering/proto/common"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type DistinctRepository interface {
	SelectByCityId(cityId int64) ([]*pb.Distinct, error)
}

type DistinctRepositoryImpl struct {
	Conn *sqlx.DB
}

func (impl *DistinctRepositoryImpl) SelectByCityId(cityId int64) ([]*pb.Distinct, error) {
	SQL := "SELECT id, distinct_name FROM `distinct` WHERE city_id = ? AND status = 1"
	stmt, err := impl.Conn.Prepare(SQL)
	if err != nil {
		return nil, err
	}
	rows, err := stmt.Query(cityId)
	if err != nil {
		return nil, err
	}
	var result []*pb.Distinct
	var id sql.NullInt64
	var distinctName sql.NullString
	for rows.Next() {
		rows.Scan(&id, &distinctName)
		result = append(result, &pb.Distinct{
			Id:           id.Int64,
			DistinctName: distinctName.String,
		})
	}
	return result, nil
}
