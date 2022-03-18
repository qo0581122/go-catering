package logic

import (
	"context"

	"catering/area/pkg/errors"
	"catering/integration/rpc/internal/svc"
	. "catering/proto/integration"

	"github.com/tal-tech/go-zero/core/logx"
)

type ChangeUserIntegrationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewChangeUserIntegrationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChangeUserIntegrationLogic {
	return &ChangeUserIntegrationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ChangeUserIntegrationLogic) ChangeUserIntegration(in *ChangeUserIntegrationReq) (*ChangeUserIntegrationResp, error) {
	resp := &ChangeUserIntegrationResp{}
	var (
		uid         = in.Uid
		changeValue = in.ChangeValue
		changeType  = in.ChangeType
		isAdd       = in.IsAdd
	)
	conn := l.svcCtx.IntegrationRepository.GetConn()
	info := l.svcCtx.IntegrationRepository.GetUserIntegration(uid)
	if isAdd == 1 {
		tx, err := conn.Begin()
		if err != nil {
			return resp, errors.InternalServerError("tx begin fail")
		}
		levelInfo := l.svcCtx.IntegrationRepository.GetUserVipLevel(uid)
		historyIntegration := info.GetHistoryIntegration()
		upNeedIntegration := levelInfo.GetUpNeedIntegration()
		for levelInfo.GetNextLevelId() != 0 && historyIntegration+int64(changeValue) >= upNeedIntegration {
			levelInfo = l.svcCtx.IntegrationRepository.GetVipLevelInfo(levelInfo.NextLevelId)
			upNeedIntegration = levelInfo.GetUpNeedIntegration()
		}
		beforeValue := info.GetCurrentIntegration()
		afterValue := beforeValue + int64(changeValue)
		err = l.svcCtx.IntegrationRepository.InsertUserIntegrationChangeLog(uid, uint32(changeType), int64(changeValue), beforeValue, afterValue, tx)
		if err != nil {
			tx.Rollback()
			return resp, errors.InternalServerError("tx insert integration change log fail")
		}
		err = l.svcCtx.IntegrationRepository.IncreaseUserIntegration(uid, int64(changeValue), levelInfo.Id, tx)
		if err != nil {
			tx.Rollback()
			return resp, errors.InternalServerError("tx increase integration fail")
		}
		if info.GetLevelId() != levelInfo.GetId() {
			err = l.svcCtx.IntegrationRepository.InsertUserVipUpLog(uid, info.GetLevelId(), levelInfo.GetId(), tx)
			if err != nil {
				tx.Rollback()
				return resp, errors.InternalServerError("tx insert user vip up log fail")
			}
		}
		if err = tx.Commit(); err != nil {
			tx.Rollback()
			return resp, errors.InternalServerError("tx commit fail")
		}
	} else {
		tx, err := conn.Begin()
		if err != nil {
			return resp, errors.InternalServerError("tx begin fail")
		}
		beforeValue := info.GetCurrentIntegration()
		afterValue := beforeValue - int64(changeValue)
		err = l.svcCtx.IntegrationRepository.DecreaseUserIntegration(uid, int64(changeValue), tx)
		if err != nil {
			tx.Rollback()
			return resp, errors.InternalServerError("tx decrease integration fail")
		}
		err = l.svcCtx.IntegrationRepository.InsertUserIntegrationChangeLog(uid, uint32(changeType), int64(changeValue), beforeValue, afterValue, tx)
		if err != nil {
			tx.Rollback()
			return resp, errors.InternalServerError("tx insert integration change log fail")
		}
		if err = tx.Commit(); err != nil {
			tx.Rollback()
			return resp, errors.InternalServerError("tx commit fail")
		}
	}
	return resp, nil
}
