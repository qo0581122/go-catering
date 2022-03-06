package request

import (
	"catering/model/common/request"
	"catering/model/system"
)

type SysDictionaryDetailSearch struct {
	system.SysDictionaryDetail
	request.PageInfo
}
