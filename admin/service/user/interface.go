package user

import (
	"catering/model"
	"catering/model/common/response"
)

type addressTagService interface {
	ListPage(pageNum, pageSize int, params *model.UserAddressTag) *response.ApiResponse
	GetOne(params *model.UserAddressTag) *model.UserAddressTag
	List(params *model.UserAddressTag) []*model.UserAddressTag
	Count() int
	Add(params *model.UserAddressTag) error
	Update(params *model.UserAddressTag) error
	Delete(id uint64) error
}

type userAddressService interface {
	ListPage(pageNum, pageSize int, params *model.UserAddress) *response.ApiResponse
	GetOne(params *model.UserAddress) *model.UserAddress
	List(params *model.UserAddress) []*model.UserAddress
	Count() int
}

type userIntegrationLogService interface {
	ListPage(pageNum, pageSize int, params *model.UserIntegrationLog) *response.ApiResponse
	GetOne(params *model.UserIntegrationLog) *model.UserIntegrationLog
	List(params *model.UserIntegrationLog) []*model.UserIntegrationLog
	Add(params *model.UserIntegrationLog) error
	Count() int
}

type userIntegrationService interface {
	ListPage(pageNum, pageSize int, params *model.UserIntegration) *response.ApiResponse
	GetOne(params *model.UserIntegration) *model.UserIntegration
	List(params *model.UserIntegration) []*model.UserIntegration
	Count() int
	IncreaseIntegration(uid uint64, changeValue int, changeType int) error
	DecreaseIntegration(uid uint64, changeValue int, changeType int) error
}

type userVipLevelLogService interface {
	ListPage(pageNum, pageSize int, params *model.UserVipLevelLog) *response.ApiResponse
	GetOne(params *model.UserVipLevelLog) *model.UserVipLevelLog
	List(params *model.UserVipLevelLog) []*model.UserVipLevelLog
	Add(params *model.UserVipLevelLog) error
	Count() int
}

type userVipLevelService interface {
	ListPage(pageNum, pageSize int, params *model.UserVipLevel) *response.ApiResponse
	GetOne(params *model.UserVipLevel) *model.UserVipLevel
	List(params *model.UserVipLevel) []*model.UserVipLevel
	Count() int
	Add(params *model.UserVipLevel, preLevelId uint64) error
	Adds(params []model.UserVipLevel) error
	Update(params *model.UserVipLevel) error
	Delete(id uint64) error
}

type userService interface {
	ListPage(pageNum, pageSize int, params *model.User) *response.ApiResponse
	GetById(id uint64) *model.User
	List(params *model.User) []*model.User
	Count() int
}
