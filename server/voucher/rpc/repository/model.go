package repository

import (
	"catering/area/pkg/errors"
	. "catering/proto/voucher"
	"time"
)

func CheckGetCount(voucher *Voucher, count int) error {
	canGetCount := voucher.GetCount
	if int(canGetCount) < count {
		return errors.InternalServerError("can not get voucher")
	}
	return nil
}

func CheckUseTime(useBeginTime, useEndTime string) error {
	currentTime := time.Now().Format("2006-01-02 15:03:04")
	if currentTime < useBeginTime || currentTime > useEndTime {
		return errors.InternalServerError("can not use voucher")
	}
	return nil
}

func CheckGetTime(voucher *Voucher) error {
	currentTime := time.Now().Format("2006-01-02 15:03:04")
	if currentTime < voucher.GetBeginTime || currentTime > voucher.GetEndTime {
		return errors.InternalServerError("can not get voucher")
	}
	return nil
}

func CheckUseStatus(useStatus int32) error {
	if useStatus != 1 {
		return errors.InternalServerError("have bee used")
	}
	return nil
}
