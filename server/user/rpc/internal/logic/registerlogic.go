package logic

import (
	"context"
	"time"

	"catering/area/pkg/bcrypt"
	pb "catering/proto/user"
	"catering/user/rpc/internal/svc"
	"catering/user/rpc/repository"

	"github.com/tal-tech/go-zero/core/logx"
	"golang.org/x/sync/errgroup"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *pb.RegisterReq) (*pb.RegisterResp, error) {
	uid := l.svcCtx.Snowflake.GetUUid()
	pass := bcrypt.EncryptPassword(in.Password)
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	tx, err := l.svcCtx.UserRepository.Conn.Begin()
	if err != nil {
		logx.Error(err, "begin tx fail")
		return nil, err
	}
	g, _ := errgroup.WithContext(l.ctx)
	g.Go(func() error {
		err := l.svcCtx.UserRepository.Insert(repository.User{
			Mobie:           in.Mobie,
			Password:        pass,
			Nickname:        in.Mobie,
			Sex:             1,
			Platform:        1,
			Avatar:          "www.baidu.com",
			Uid:             uid,
			RegistryDid:     "0",
			RegistryTime:    currentTime,
			RecentLoginDid:  "0",
			RecentLoginTime: currentTime,
			UpdatedTime:     currentTime,
		}, tx)
		if err != nil {
			logx.Error(err)
			return err
		}
		return nil
	})
	if err := g.Wait(); err != nil {
		tx.Rollback()
		return nil, err
	}
	if err = tx.Commit(); err != nil {
		tx.Rollback()
		logx.Error(err)
		return nil, err
	}
	return &pb.RegisterResp{
		Code: int32(200),
		Msg:  "success",
	}, nil
}
