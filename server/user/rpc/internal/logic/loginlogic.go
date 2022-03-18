package logic

import (
	"catering/area/pkg/bcrypt"
	"catering/area/pkg/errors"
	pb "catering/proto/user"
	"catering/user/rpc/internal/svc"
	"context"

	"github.com/tal-tech/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *pb.LoginReq) (*pb.LoginResp, error) {
	handler := NewHandler(in.LoginType)
	if handler == nil {
		return nil, nil
	}

	info, err := l.svcCtx.UserRepository.FindOneByMobie(in.Mobie)
	if info == nil || err != nil {
		logx.Error(err)
		return nil, errors.InternalServerError("mysql is fail")
	}
	if !handler.check(in.Password, info.Password, in.Smscode) {
		return nil, errors.Unauthorized("mobie or password is wrong")
	}
	return &pb.LoginResp{
		User: info,
	}, nil
}

type CheckHandler interface {
	check(rawPassword, encryPassword, smscode string) bool
}

type DefaultHandler struct {
}

const (
	PasswordLogin int32 = 0 + iota
	SmscodeLogin
)

type PasswordCheckHandler DefaultHandler
type SmscodeCheckHandler DefaultHandler

func (h *PasswordCheckHandler) check(rawPassword, encryPassword, smscode string) bool {
	res := bcrypt.ComparePassword(rawPassword, encryPassword)
	if res {
		return res
	}
	return false
}

func (h *SmscodeCheckHandler) check(rawPassword, encryPassword, smscode string) bool {
	return true
}

func NewHandler(loginType int32) CheckHandler {
	switch loginType {
	case PasswordLogin:
		return &PasswordCheckHandler{}
	case SmscodeLogin:
		return &SmscodeCheckHandler{}
	default:
		break
	}
	return nil
}
