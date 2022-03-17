package valid

import (
	"errors"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

// BindAndValid binds and validates data
func BindAndValid(c *gin.Context, form interface{}) (string, error) {
	err := c.ShouldBind(form)
	if err != nil {
		return "参数校验错误", err
	}

	valid := validation.Validation{}
	check, err := valid.Valid(form)
	if err != nil {
		return "服务器处理错误500", err
	}
	if !check {
		MarkErrors(valid.Errors)
		errMessage := ""
		for _, err := range valid.Errors {
			errMessage += err.Key + "=>" + err.Message + "  "
		}
		return errMessage, errors.New("")
	}

	return "", nil
}
