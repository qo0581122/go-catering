package repository

import (
	"catering/area/pkg/errors"
	. "catering/proto/coupon"
	"time"
)

func CheckGetCount(coupon *Coupon, count int) error {
	canGetCount := coupon.GetCount
	if int(canGetCount) < count {
		return errors.InternalServerError("can not get coupon")
	}
	return nil
}

func CheckUseTime(useBeginTime, useEndTime string) error {
	currentTime := time.Now().Format("2006-01-02 15:03:04")
	if currentTime < useBeginTime || currentTime > useEndTime {
		return errors.InternalServerError("can not use coupon")
	}
	return nil
}

func CheckGetTime(coupon *Coupon) error {
	currentTime := time.Now().Format("2006-01-02 15:03:04")
	if currentTime < coupon.GetBeginTime || currentTime > coupon.GetEndTime {
		return errors.InternalServerError("can not get coupon")
	}
	return nil
}

func CheckUseStatus(useStatus int32) error {
	if useStatus != 1 {
		return errors.InternalServerError("have bee used")
	}
	return nil
}

func GetUseTime(counpon *Coupon) (string, string) {
	useBeginTime := counpon.GetUseBeginTime()
	useEndTime := counpon.GetUseEndTime()
	if counpon.GetValidTimeType() == 2 {
		day := int64(counpon.GetValidDay())
		useEndTime = time.Now().Add(time.Duration(day * int64(time.Hour) * 24)).Format("2006-01-02 15:03:04")
		useBeginTime = time.Now().Format("2006-01-02 15:03:04")
	}
	return useBeginTime, useEndTime
}
