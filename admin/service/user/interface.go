package user

import (
	"catering/model"
	"catering/model/common/response"
)

type addressTagService interface {
	ListPage(pageNum, pageSize int, params *model.AddressTag) *response.ApiResponse
	GetById(id uint64) *model.AddressTag
	List(params *model.AddressTag) []*model.AddressTag
	Count() int
	Add(params *model.AddressTag) error
	Update(params *model.AddressTag) error
	Delete(id uint64) error
}

type userAddressService interface {
	ListPage(pageNum, pageSize int, params *model.UserAddress) *response.ApiResponse
	GetById(id uint64) *model.UserAddress
	List(params *model.UserAddress) []*model.UserAddress
	Count() int
}

type userIntegrationLogService interface {
	ListPage(pageNum, pageSize int, params *model.UserIntegrationLog) *response.ApiResponse
	GetById(id uint64) *model.UserIntegrationLog
	List(params *model.UserIntegrationLog) []*model.UserIntegrationLog
	Count() int
}

type userIntegrationService interface {
	ListPage(pageNum, pageSize int, params *model.UserIntegration) *response.ApiResponse
	GetById(id uint64) *model.UserIntegration
	GetByUin(uin uint64) *model.UserIntegration
	List(params *model.UserIntegration) []*model.UserIntegration
	Count() int
	IncreaseIntegration(uid uint64, changeValue int, changeType int) error
	DecreaseIntegration(uid uint64, changeValue int, changeType int) error
}

type userVipLevelLogService interface {
	ListPage(pageNum, pageSize int, params *model.UserVipLevelLog) *response.ApiResponse
	GetById(id uint64) *model.UserVipLevelLog
	List(params *model.UserVipLevelLog) []*model.UserVipLevelLog
	Count() int
}

type userVipLevelService interface {
	ListPage(pageNum, pageSize int, params *model.UserVipLevel) *response.ApiResponse
	GetById(id uint64) *model.UserVipLevel
	List(params *model.UserVipLevel) []*model.UserVipLevel
	Count() int
	Add(params *model.UserVipLevel, preLevelId uint64) error
	Adds(params []*model.UserVipLevel) error
	Update(params *model.UserVipLevel) error
	Delete(id uint64) error
	GetByUin(uin uint64) *model.UserVipLevel
}

type userService interface {
	ListPage(pageNum, pageSize int, params *model.User) *response.ApiResponse
	GetById(id uint64) *model.User
	List(params *model.User) []*model.User
	Count() int
}
