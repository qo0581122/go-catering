package repository

import (
	pb "catering/proto/advertisement"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type AdBannerRepository interface {
	GetList() []*pb.AdBanner
}

type AdBannerRepositoryImpl struct {
	Conn *sqlx.DB
}

func (impl *AdBannerRepositoryImpl) GetList() []*pb.AdBanner {
	SQL := "SELECT id, banner_type, title, pic_url, remark, status, sort, begin_time, end_time FROM ad_banner WHERE status = 1 ORDER BY begin_time DESC, end_time ASC, sort DESC"
	stmt, err := impl.Conn.Prepare(SQL)
	if err != nil {
		return nil
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		return nil
	}
	var result []*pb.AdBanner
	var id sql.NullInt64
	var title, picUrl, remark, beginTime, endTime sql.NullString
	var bannerType, sort sql.NullInt32
	for rows.Next() {
		rows.Scan(&id, &bannerType, &title, &picUrl, &remark, &sort, &beginTime, &endTime)
		result = append(result, &pb.AdBanner{
			Id:         id.Int64,
			BannerType: bannerType.Int32,
			Title:      title.String,
			PicUrl:     picUrl.String,
			Remark:     remark.String,
			Sort:       sort.Int32,
			BeginTime:  beginTime.String,
			EndTime:    endTime.String,
		})
	}
	return result
}
