package request

import (
	"catering/model/autocode"
	"catering/model/common/request"
)

type {{.StructName}}Search struct{
    autocode.{{.StructName}}
    request.PageInfo
}