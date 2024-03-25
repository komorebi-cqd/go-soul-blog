package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/go-soul-blog/global"
	"github.com/go-soul-blog/pkg/app"
	"github.com/go-soul-blog/pkg/errcode"
)

type Articles struct {
}

func NewArticles() Articles {
	return Articles{}
}

// @Summary 获取单个文章
// @Param id path int true "文章ID"
// @Produce json
// @Success 200 {object} model.Articles "成功"
// @Router /api/v1/articles/{id} [get]
func (t Articles) Get(c *gin.Context) {

}

// @Summary 获取文章列表
// @Param name query string false "文章名称" maxlength(100)
// @Param is_draft query int false "是否是草稿" Enums(0, 1) default(1)
// @Param page query int false "页码"
// @Param page_size query int false "每页数量"
// @Produce json
// @Success 200 {object} model.ArticlesSwagger "成功"
// @Router /api/v1/articles [get]
func (t Articles) List(c *gin.Context) {
	params := struct {
		Name    string `form:"name" binding:"max=100"`
		IsDraft uint8  `form:"is_draft,default=1" binding:"oneof=0 1"`
	}{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &params)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		errRsp := errcode.InvalidParams.WithDetails(errs.Errors()...)
		response.ToErrorResponse(errRsp)
		return
	}

	response.ToResponse(gin.H{})
}

// @Summary 新增文章
// @Param name query string true "文章名称" minlength(3) maxlength(100)
// @Param is_draft query int false "是否是草稿" Enums(0, 1) default(1)
// @Produce json
// @Success 200 {object} model.Articles "成功"
// @Router /api/v1/articles [post]
func (t Articles) Create(c *gin.Context) {

}

// @Summary 更新文章
// @Param id path int true "文章ID"
// @Param is_draft path int false "是否是草稿" Enums(0, 1) default(1)
// @Param name query string false "文章名称" minlength(3) maxlength(100)
// @Produce json
// @Success 200 {object} model.Articles "成功"
// @Router /api/v1/articles/{id} [put]
func (t Articles) Update(c *gin.Context) {

}

// @Summary 删除文章
// @Param id path int true "文章ID"
// @Produce json
// @Success 200 {object} model.Article "成功"
// @Router /api/v1/articles/{id} [delete]
func (t Articles) Delete(c *gin.Context) {

}
