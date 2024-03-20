package v1

import (
	"github.com/gin-gonic/gin"
)

type Article struct {
}

func NewArticle() Article {
	return Article{}
}

// @Summary 获取单个文章
// @Param id path int true "文章ID"
// @Produce json
// @Success 200 {object} model.ArticleSwagger "成功"
// @Router /api/v1/articles/{id} [get]
func (t Article) Get(c *gin.Context) {

}

// @Summary 获取多个文章
// @Param name query string false "文章名称" maxlength(100)
// @Param state query int false "状态" Enums(0, 1) default(1)
// @Param page query int false "页码"
// @Param page_size query int false "每页数量"
// @Produce json
// @Success 200 {object} model.ArticleSwagger "成功"
// @Router /api/v1/articles [get]
func (t Article) List(c *gin.Context) {

}

// @Summary 新增文章
// @Param name query string true "文章名称" minlength(3) maxlength(100)
// @Param state query int false "状态" Enums(0, 1) default(1)
// @Param created_by body string true "创建者" minlength(3) maxlength(100)
// @Produce json
// @Success 200 {object} model.Article "成功"
// @Router /api/v1/articles [post]
func (t Article) Create(c *gin.Context) {

}

// @Summary 更新文章
// @Param id path int true "文章ID"
// @Param state path int false "状态" Enums(0, 1) default(1)
// @Param name query string false "文章名称" minlength(3) maxlength(100)
// @Param created_by body string true "修改者" minlength(3) maxlength(100)
// @Produce json
// @Success 200 {object} model.Article "成功"
// @Router /api/v1/articles/{id} [put]
func (t Article) Update(c *gin.Context) {

}

// @Summary 删除文章
// @Param id path int true "文章ID"
// @Produce json
// @Success 200 {object} model.Article "成功"
// @Router /api/v1/articles/{id} [delete]
func (t Article) Delete(c *gin.Context) {

}
