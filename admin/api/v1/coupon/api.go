package coupon

import (
	"catering/model"
	"catering/model/common/response"
	"catering/pkg/app"
	"strconv"

	"github.com/gin-gonic/gin"
)

type QueryParams struct {
	PageSize   int    `uri:"pageSize" json:"pageSize" form:"pageSize" valid:"Required"`
	PageNum    int    `uri:"pageNum" json:"pageNum" form:"pageNum" valid:"Required"`
	ProductId  uint64 `uri:"product_id"  form:"product_id" json:"product_id"`
	CouponName string `uri:"coupon_name"  form:"coupon_name" json:"coupon_name"`
	Status     uint32 `uri:"status" json:"status" form:"status"`
}

type QueryResponse struct {
	List  []*model.CouponDetail `json:"list,inline"`
	Total int                   `json:"total"`
}

func List(c *gin.Context) {
	var (
		params = QueryParams{}
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

type AddForm struct {
	CouponName    string `json:"coupon_name" form:"coupon_name" valid:"Required"`
	Price         uint64 `json:"price" form:"price" valid:"Required"`
	ProductId     uint64 `json:"product_id" form:"product_id" valid:"Required"`
	LeastUsePrice uint64 `json:"least_use_price" form:"least_use_price" `
	PicUrl        string `json:"pic_url" form:"pic_url" `
	TotalCount    uint64 `json:"total_count" form:"total_count" valid:"Required"`
	GetType       uint32 `json:"get_type" form:"get_type" valid:"Required"`
	ValidTimeType uint32 `json:"valid_time_type" form:"valid_time_type"  `
	ValidDay      uint32 `json:"valid_day" form:"valid_day" `
	UseType       uint32 `json:"use_type" form:"use_type" valid:"Required"`
	Status        uint32 `json:"status" form:"status" `
	GetBeginTime  string `json:"get_begin_time" form:"get_begin_time" valid:"Required"`
	GetEndTime    string `json:"get_end_time" form:"get_end_time" valid:"Required"`
	UseBeginTime  string `json:"use_begin_time" form:"use_begin_time" `
	UseEndTime    string `json:"use_end_time" form:"use_end_time" `
	GetCount      uint32 `json:"get_count" form:"get_count" valid:"Required"`
	Description   string `json:"description" form:"description" `
}

func Add(c *gin.Context) {
	var (
		form = AddForm{}
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

type UpdateForm struct {
	Id            uint64 `form:"id" json:"id" valid:"Required"`
	CouponName    string `json:"coupon_name" form:"coupon_name" valid:"Required"`
	Price         uint64 `json:"price" form:"price" valid:"Required"`
	ProductId     uint64 `json:"product_id" form:"product_id" valid:"Required"`
	LeastUsePrice uint64 `json:"least_use_price" form:"least_use_price" `
	PicUrl        string `json:"pic_url" form:"pic_url" `
	TotalCount    uint64 `json:"total_count" form:"total_count" valid:"Required"`
	RemainCount   uint64 `json:"remain_count" form:"remain_count" valid:"Required"`
	GetType       uint32 `json:"get_type" form:"get_type" valid:"Required"`
	ValidTimeType int    `json:"valid_time_type" form:"valid_time_type"  `
	ValidDay      uint32 `json:"valid_day" form:"valid_day" `
	UseType       uint32 `json:"use_type" form:"use_type" valid:"Required"`
	Status        int    `json:"status" form:"status" `
	GetBeginTime  string `json:"get_begin_time" form:"get_begin_time" valid:"Required"`
	GetEndTime    string `json:"get_end_time" form:"get_end_time" valid:"Required"`
	UseBeginTime  string `json:"use_begin_time" form:"use_begin_time" `
	UseEndTime    string `json:"use_end_time" form:"use_end_time" `
	GetCount      uint32 `json:"get_count" form:"get_count" valid:"Required"`
	Description   string `json:"description" form:"description" `
}

func Update(c *gin.Context) {
	var (
		form = UpdateForm{}
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
