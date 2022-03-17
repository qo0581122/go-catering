package voucher

import (
	"catering/model"
	"catering/model/common/response"
	"catering/model/voucher/request"
	"catering/pkg/valid"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ListPage(c *gin.Context) {
	var (
		form = request.VoucherQueryParams{}
	)
	msg, err := valid.BindAndValid(c, &form)
	if err != nil {
		response.FailWithMessage(msg, c)
		return
	}
	model := &model.Voucher{
		Name:   form.VoucherName,
		Status: form.Status,
	}
	res := voucherService.ListPage(form.PageNum, form.PageSize, model)
	response.OkWithData(res, c)
}

func Add(c *gin.Context) {
	var (
		form = request.VoucherAddForm{}
	)
	msg, err := valid.BindAndValid(c, &form)
	if err != nil {
		response.FailWithMessage(msg, c)
		return
	}
	getBeginTime := model.ParseTime(form.GetBeginTime)
	getEndTime := model.ParseTime(form.GetEndTime)
	useBeginTime := model.ParseTime(form.UseBeginTime)
	useEndTime := model.ParseTime(form.UseEndTime)
	voucher := &model.Voucher{
		Name:          form.VoucherName,
		Price:         form.Price,
		LeastUsePrice: form.LeastUsePrice,
		PicUrl:        form.PicUrl,
		TotalCount:    form.TotalCount,
		GetType:       form.GetType,
		ValidTimeType: form.ValidTimeType,
		ValidDay:      form.ValidDay,
		UseType:       form.UseType,
		Status:        form.Status,
		GetBeginTime:  getBeginTime,
		GetEndTime:    getEndTime,
		UseBeginTime:  useBeginTime,
		UseEndTime:    useEndTime,
		GetCount:      form.GetCount,
		Description:   form.Description,
		RemainCount:   form.TotalCount,
	}

	err = voucherService.Add(voucher)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.Ok(c)
}

func Update(c *gin.Context) {
	var (
		form = request.VoucherUpdateForm{}
	)
	msg, err := valid.BindAndValid(c, &form)
	if err != nil {
		response.FailWithMessage(msg, c)
		return
	}
	getBeginTime := model.ParseTime(form.GetBeginTime)
	getEndTime := model.ParseTime(form.GetEndTime)
	useBeginTime := model.ParseTime(form.UseBeginTime)
	useEndTime := model.ParseTime(form.UseEndTime)
	voucher := &model.Voucher{
		Name:          form.VoucherName,
		Price:         form.Price,
		LeastUsePrice: form.LeastUsePrice,
		PicUrl:        form.PicUrl,
		TotalCount:    form.TotalCount,
		GetType:       form.GetType,
		ValidTimeType: uint32(form.ValidTimeType),
		ValidDay:      form.ValidDay,
		UseType:       form.UseType,
		Status:        uint32(form.Status),
		GetBeginTime:  getBeginTime,
		GetEndTime:    getEndTime,
		UseBeginTime:  useBeginTime,
		UseEndTime:    useEndTime,
		GetCount:      form.GetCount,
		Description:   form.Description,
		RemainCount:   0,
	}
	err = voucherService.Update(voucher)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.Ok(c)
}

func Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.DefaultQuery("id", ""))
	err := voucherService.Delete(uint64(id))
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.Ok(c)
}
