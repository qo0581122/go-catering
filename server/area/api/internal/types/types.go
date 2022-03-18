// Code generated by goctl. DO NOT EDIT.
package types

import (
	pb "catering/proto/common"
)

type GetProvincesReq struct {
}

type GetProvincesResp struct {
	Provinces []*pb.Province
}

type GetCitysReq struct {
	ProvinceId int64 `json:"provinceId"`
}

type GetCitysResp struct {
	Citys []*pb.City
}

type GetDistinctsReq struct {
	CityId int64 `json:"cityId"`
}

type GetDistinctsResp struct {
	Distincts []*pb.Distinct
}
