package coupon

import (
	"catering/model"
	"catering/model/common/response"
	"catering/model/coupon/request"
	"catering/pkg/app"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ListPage(c *gin.Context) {
	var (
		params = request.CouponQueryParams{}
	)
	msg, err := app.BindAndValid(c, &params)
	if err != nil {
		response.FailWithMessage(msg, c)
		return
	}
	model := &model.Coupon{
		Name:      params.CouponName,
		ProductId: params.ProductId,
		Status:    params.Status,
	}
	res := couponService.ListPage(params.PageNum, params.PageSize, model)
	response.OkWithData(res, c)
}

func Add(c *gin.Context) {
	var (
		form = request.CouponAddForm{}
	)
	msg, err := app.BindAndValid(c, &form)
	if err != nil {
		response.FailWithMessage(msg, c)
		return
	}
	// currentTime := time.Now().Format("2006-01-02 15:04:05")
	getBeginTime := model.ParseTime(form.GetBeginTime)
	getEndTime := model.ParseTime(form.GetEndTime)
	useBeginTime := model.ParseTime(form.UseBeginTime)
	useEndTime := model.ParseTime(form.UseEndTime)
	coupon := &model.Coupon{
		Name:          form.CouponName,
		Price:         form.Price,
		ProductId:     form.ProductId,
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

	err = couponService.Add(coupon)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.Ok(c)
}

func Update(c *gin.Context) {
	var (
		form = request.CouponUpdateForm{}
	)
	msg, err := app.BindAndValid(c, &form)
	if err != nil {
		response.FailWithMessage(msg, c)
		return
	}
	getBeginTime := model.ParseTime(form.GetBeginTime)
	getEndTime := model.ParseTime(form.GetEndTime)
	useBeginTime := model.ParseTime(form.UseBeginTime)
	useEndTime := model.ParseTime(form.UseEndTime)
	coupon := &model.Coupon{
		Model: model.Model{
			Id: form.Id,
		},
		Name:          form.CouponName,
		Price:         form.Price,
		ProductId:     form.ProductId,
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
	err = couponService.Update(coupon)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.Ok(c)
}

func Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.DefaultQuery("id", ""))
	err := couponService.Delete(uint64(id))
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
}
