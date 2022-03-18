package repository

import (
	pb "catering/proto/advertisement"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type AdHotActivitieRepository interface {
	GetList() []*pb.AdHotActivitie
}

type AdHotActivitieRepositoryImpl struct {
	Conn *sqlx.DB
}

func (impl *AdHotActivitieRepositoryImpl) GetList() []*pb.AdHotActivitie {
	SQL := "SELECT id, activatie_name, title, subtitle, content, pic_url, status, sort, begin_time, end_time WHERE status = 1 ORDER BY begin_time DESC, end_time ASC, sort DESC"
	stmt, err := impl.Conn.Prepare(SQL)
	if err != nil {
		return nil
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		return nil
	}
	var result []*pb.AdHotActivitie
	var id sql.NullInt64
	var activatieName, title, subTitle, content, beginTime, endTime, picUrl sql.NullString
	var status, sort sql.NullInt32
	for rows.Next() {
		rows.Scan(&id, &activatieName, &title, &subTitle, &content, &picUrl, &status, &sort, &beginTime, &endTime)
		result = append(result, &pb.AdHotActivitie{
			Id:            id.Int64,
			ActivatieName: activatieName.String,
			Title:         title.String,
			Subtitle:      subTitle.String,
			Content:       content.String,
			Status:        status.Int32,
			Sort:          sort.Int32,
			PicUrl:        picUrl.String,
			BeginTime:     beginTime.String,
			EndTime:       endTime.String,
		})
	}
	return result
}
