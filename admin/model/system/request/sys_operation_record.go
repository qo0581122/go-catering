package request

import (
	"catering/model/common/request"
	"catering/model/system"
)

type SysOperationRecordSearch struct {
	system.SysOperationRecord
	request.PageInfo
}
