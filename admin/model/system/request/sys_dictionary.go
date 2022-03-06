package request

import (
	"catering/model/common/request"
	"catering/model/system"
)

type SysDictionarySearch struct {
	system.SysDictionary
	request.PageInfo
}
