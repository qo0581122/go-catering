package user_address

import (
	"catering/model"
	"catering/model/common/response"
	"catering/pkg/app"

	"github.com/gin-gonic/gin"
)

type UserAddressQueryParams struct {
	PageSize int    `uri:"pageSize" json:"pageSize" form:"pageSize" valid:"Required"`
	PageNum  int    `uri:"pageNum" json:"pageNum" form:"pageNum" valid:"Required"`
	Uid      uint64 `uri:"uid" json:"uid" form:"uid"`
}

func ListPage(c *gin.Context) {
	var (
		form = UserAddressQueryParams{}
	)
	msg, err := app.BindAndValid(c, &form)
	if err != nil {
		response.FailWithMessage(msg, c)
		return
	}
	model := &model.UserAddress{
		Uid: form.Uid,
	}
	res := userAddressService.ListPage(form.PageNum, form.PageSize, model)
	response.OkWithData(res, c)
}
