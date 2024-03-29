package v1

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-soul-blog/global"
	"github.com/go-soul-blog/internal/service"
	"github.com/go-soul-blog/pkg/app"
	"github.com/go-soul-blog/pkg/errcode"
)

type Tags struct{}

func NewTags() Tags {
	return Tags{}
}

// @Summary 获取所有标签
// @Produce json
// @Success 200 {object} model.TagsSwagger "成功"
// @Router /api/v1/tags [get]
func (t Tags) List(c *gin.Context) {
	params := service.TagListRequest{}
	fmt.Printf("Create params %v c.params %v", params, c.Param("id"))
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &params)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		errRsp := errcode.InvalidParams.WithDetails(errs.Errors()...)
		response.ToErrorResponse(errRsp)
		return
	}
	svc := service.New(c.Request.Context())
	pager := app.Pager{Page: app.GetPage(c), PageSize: app.GetPageSize(c)}
	totalRows, err := svc.CountTag(&service.CountTagRequest{Name: params.Name})
	if err != nil {
		global.Logger.Errorf("svc.CountTag err: %v", err)
		response.ToErrorResponse(errcode.ErrorCountTagFail)
		return
	}

	tags, err := svc.GetTagList(&params, &pager)
	if err != nil {
		global.Logger.Errorf("svc.GetTagList err: %v", err)
		response.ToErrorResponse(errcode.ErrorGetTagListFail)
		return
	}

	response.ToResponseList(tags, int(totalRows))

}

// @Summary 新增标签
// @Param name query string true "标签名称" minlength(1) maxlength(100)
// @Param state query int false "状态" Enums(0, 1) default(1)
// @Produce json
// @Success 200 {object} model.Tags "成功"
// @Router /api/v1/tags [post]
func (t Tags) Create(c *gin.Context) {
	params := service.CreateTagRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &params)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		errRsp := errcode.InvalidParams.WithDetails(errs.Errors()...)
		response.ToErrorResponse(errRsp)
		return
	}

	svc := service.New(c.Request.Context())

	err := svc.CreateTag(&params)
	if err != nil {
		global.Logger.Errorf("svc.CreateTag errs: %v", err)
		response.ToErrorResponse(errcode.ErrorCreateTagFail)
		return
	}
	response.ToResponse(nil)
}

// @Summary 更新标签
// @Param id path int true "标签ID"
// @Param name query string false "标签名称" minlength(1) maxlength(100)
// @Param state query int false "状态" Enums(0, 1) default(1)
// @Param created_by body string true "修改者" minlength(3) maxlength(100)
// @Produce json
// @Success 200 {object} model.Tags "成功"
// @Router /api/v1/tags/{id} [put]
func (t Tags) Update(c *gin.Context) {
	params := service.UpdateTagRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &params)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		errRsp := errcode.InvalidParams.WithDetails(errs.Errors()...)
		response.ToErrorResponse(errRsp)
		return
	}
	svc := service.New(c.Request.Context())
	err := svc.UpdateTag(&service.UpdateTagRequest{ID: params.ID, Name: params.Name})
	if err != nil {
		global.Logger.Errorf("svc.UpdateTag errs: %v", err)
		response.ToErrorResponse(errcode.ErrorUpdateTagFail)
		return
	}
	response.ToResponse(nil)
}

// @Summary 删除标签
// @Param id path int true "标签ID"
// @Produce json
// @Success 200 {object} model.Tags "成功"
// @Router /api/v1/tags/{id} [delete]
func (t Tags) Delete(c *gin.Context) {
	params := service.DeleteTagRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValidUri(c, &params)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		errRsp := errcode.InvalidParams.WithDetails(errs.Errors()...)
		response.ToErrorResponse(errRsp)
		return
	}
	svc := service.New(c.Request.Context())
	err := svc.DeleteTag(&service.DeleteTagRequest{ID: params.ID})
	if err != nil {
		global.Logger.Errorf("svc.DeleteTag errs: %v", err)
		response.ToErrorResponse(errcode.ErrorDeleteTagFail)
		return
	}
	response.ToResponse(nil)
}
