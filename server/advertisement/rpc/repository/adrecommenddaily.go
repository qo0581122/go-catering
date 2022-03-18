package repository

import (
	pb "catering/proto/advertisement"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type AdRecommendDailyRepository interface {
	GetList() []*pb.AdRecommendDaily
}

type AdRecommendDailyRepositoryImpl struct {
	Conn *sqlx.DB
}

func (impl *AdRecommendDailyRepositoryImpl) GetList() []*pb.AdRecommendDaily {
	SQL := "SELECT id, product_id, title, subtitle, content, pic_url, status, sort, begin_time, end_time WHERE status = 1 ORDER BY begin_time DESC, end_time ASC, sort DESC"
	stmt, err := impl.Conn.Prepare(SQL)
	if err != nil {
		return nil
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		return nil
	}
	var result []*pb.AdRecommendDaily
	var id, productId sql.NullInt64
	var title, subTitle, content, beginTime, endTime, picUrl sql.NullString
	var status, sort sql.NullInt32
	for rows.Next() {
		rows.Scan(&id, &productId, &title, &subTitle, &content, &picUrl, &status, &sort, &beginTime, &endTime)
		result = append(result, &pb.AdRecommendDaily{
			Id:        id.Int64,
			ProductId: productId.Int64,
			Title:     title.String,
			Subtitle:  subTitle.String,
			Content:   content.String,
			Status:    status.Int32,
			Sort:      sort.Int32,
			PicUrl:    title.String,
			BeginTime: beginTime.String,
			EndTime:   endTime.String,
		})
	}
	return result
}
