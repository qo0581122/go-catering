package product_batch

import (
	"catering/model"
	"catering/model/common/response"
	"catering/model/product/request"
	"catering/pkg/app"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ListPage(c *gin.Context) {
	var (
		form = request.BatchQueryParams{}
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

func Add(c *gin.Context) {
	var (
		form = request.BatchAddForm{}
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

func Update(c *gin.Context) {
	var (
		form = request.BatchUpdateForm{}
	)
	msg, err := app.BindAndValid(c, &form)
	if err != nil {
		response.FailWithMessage(msg, c)
		return
	}
	batch := &model.ProductBatch{
		Model:     model.Model{ID: form.Id},
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

func List(c *gin.Context) {
	batch := &model.ProductBatch{}
	result := productBatchService.List(batch)
	response.OkWithData(response.NewApiResponse(result), c)
}
