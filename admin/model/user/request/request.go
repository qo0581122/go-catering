package request

import "catering/model/common/request"

type UserAddressTagQueryParams struct {
	request.PageParams
	AddressTagName string `uri:"tag_name"  form:"tag_name" json:"tag_name"`
	Status         uint32 `uri:"status" json:"status" form:"status"`
}

type UserAddressTagAddForm struct {
	TagName         string `json:"tag_name" form:"tag_name" valid:"Required"`
	TextColor       string `json:"text_color" form:"text_color" valid:"Required"`
	BorderColor     string `json:"border_color" form:"border_color" `
	BackgroundColor string `json:"background_color" form:"background_color" `
	Sort            uint32 `json:"sort" form:"sort" valid:"Required"`
	Status          uint32 `json:"status" form:"status" valid:"Required"`
}

type UserAddressTagUpdateForm struct {
	Id              uint64 `form:"id" json:"id" valid:"Required"`
	TagName         string `json:"tag_name" form:"tag_name" valid:"Required"`
	TextColor       string `json:"text_color" form:"text_color" valid:"Required"`
	BorderColor     string `json:"border_color" form:"border_color" `
	BackgroundColor string `json:"background_color" form:"background_color" `
	Sort            uint32 `json:"sort" form:"sort" valid:"Required"`
	Status          uint32 `json:"status" form:"status" valid:"Required"`
}

type UserIntegrationQueryParams struct {
	request.PageParams
	Uid     int    `uri:"uid"  form:"uid" json:"uid"`
	LevelId uint64 `uri:"level_id" json:"level_id" form:"level_id"`
}

type UserIntegrationChangeForm struct {
	Id          uint64 `uri:"id"  form:"id" json:"id"`
	ChangeValue int    `uri:"change_value"  form:"change_value" json:"change_value"`
	ChangeType  int    `uri:"change_type"  form:"change_type" json:"change_type"`
	AddStatus   int    `uri:"add_status"  form:"add_status" json:"add_status"`
}

type UserIntegrationLogQueryParams struct {
	request.PageParams
	Uid        int `uri:"uid" json:"uid" form:"uid"`
	ChangeType int `uri:"change_type" json:"change_type" form:"change_type"`
}

type UserVipLevelQueryParams struct {
	request.PageParams
}

type UserVipLevelAddForm struct {
	PreLevelId        uint64 `json:"pre_level_id" form:"pre_level_id"`
	LevelName         string `json:"level_name" form:"level_name" valid:"Required"`
	UpNeedIntegration int    `json:"up_need_integration" form:"up_need_integration" valid:"Required"`
	LevelDiscount     int    `json:"level_discount" form:"level_discount" `
	NextLevelId       uint64 `json:"next_level_id" form:"next_level_id"`
	Level             int    `json:"level" form:"level"`
}

type UserVipLevelAddForms struct {
	Forms []struct {
		LevelName         string `json:"level_name" form:"level_name" valid:"Required"`
		UpNeedIntegration int    `json:"up_need_integration" form:"up_need_integration" valid:"Required"`
		LevelDiscount     int    `json:"level_discount" form:"level_discount" `
		Level             int    `json:"level" form:"level"`
	} `json:"forms"`
}

type UserVipLevelUpdateForm struct {
	Id                uint64 `form:"id" json:"id" valid:"Required"`
	LevelName         string `json:"level_name" form:"level_name" valid:"Required"`
	UpNeedIntegration int    `json:"up_need_integration" form:"up_need_integration" valid:"Required"`
	LevelDiscount     int    `json:"level_discount" form:"level_discount" `
	NextLevelId       uint64 `json:"next_level_id" form:"next_level_id"`
}

type UserVipLevelLogQueryParams struct {
	request.PageParams
	Uid           int    `uri:"uid" json:"uid" form:"uid"`
	BeforeLevelId uint64 `uri:"before_level_id" json:"before_level_id" form:"before_level_id"`
	AfterLevelId  uint64 `uri:"after_level_id" json:"after_level_id" form:"after_level_id"`
}
