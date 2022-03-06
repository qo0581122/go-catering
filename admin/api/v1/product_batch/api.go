package product_batch

import (
	"catering/model"
	"catering/model/common/response"
	"catering/pkg/app"
	"strconv"

	"github.com/gin-gonic/gin"
)

type QueryParams struct {
	PageSize  int    `uri:"pageSize" json:"pageSize" form:"pageSize" valid:"Required"`
	PageNum   int    `uri:"pageNum" json:"pageNum" form:"pageNum" valid:"Required"`
	BatchName string `uri:"batch_name"  form:"batch_name" json:"batch_name"`
	Status    int    `uri:"status" json:"status" form:"status"`
}

func List(c *gin.Context) {
	var (
		form = QueryParams{}
	)
	msg, err := app.BindAndValid(c, &form)
	if err != nil {
		response.FailWithMessage(msg, c)
		return
	}
	batch := &model.ProductBatch{
		BatchName: form.BatchName,
		Status:    form.Status,
	}
	result := productBatchService.ListPage(form.PageNum, form.PageSize, batch)
	response.OkWithData(result, c)
}

type AddForm struct {
	Status    int    `form:"status" json:"status" valid:"Range(1,2)"`
	BatchName string `json:"batch_name" form:"batch_name" valid:"Required"`
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
	batch := &model.ProductBatch{
		BatchName: form.BatchName,
		Status:    form.Status,
	}
	err = productBatchService.Add(batch)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.Ok(c)
}

type UpdateForm struct {
	Id        uint64 `form:"id" json:"id" valid:"Required"`
	Status    int    `form:"status" json:"status" valid:"Range(1,2)"`
	BatchName string `json:"batch_name" form:"batch_name" valid:"Required"`
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
	batch := &model.ProductBatch{
		Model:     model.Model{Id: form.Id},
		BatchName: form.BatchName,
		Status:    form.Status,
	}
	err = productBatchService.Update(batch)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.Ok(c)
}

func Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := productBatchService.Delete(uint64(id))
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.Ok(c)
}

func ListAll(c *gin.Context) {
	batch := &model.ProductBatch{}
	result := productBatchService.List(batch)
	response.OkWithData(response.NewApiResponse(result), c)
}
