package valid

import (
	"catering/global"

	"github.com/astaxie/beego/validation"
	"go.uber.org/zap"
)

// MarkErrors logs error logs
func MarkErrors(errors []*validation.Error) {
	for _, err := range errors {
		global.Log.Error(err.Key, zap.String("message", err.Message))
	}

	return
}
