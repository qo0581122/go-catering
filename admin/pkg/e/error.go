package e

import "errors"

var (
	ErrMysqlConfigCheckFail error = errors.New("Mysql配置校验出错")
)
