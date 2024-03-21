package v1

import "github.com/gin-gonic/gin"

type Tag struct{}

func NewTag() Tag {
	return Tag{}
}

// @Summary 获取所有标签
// @Produce json
// @Success 200 {object} model.TagSwagger "成功"
// @Router /api/v1/tags [get]
func (t Tag) List(c *gin.Context) {

}

// @Summary 新增标签
// @Param name query string true "标签名称" minlength(3) maxlength(100)
// @Param state query int false "状态" Enums(0, 1) default(1)
// @Param created_by body string true "创建者" minlength(3) maxlength(100)
// @Produce json
// @Success 200 {object} model.Tag "成功"
// @Router /api/v1/tags [post]
func (t Tag) Create(c *gin.Context) {

}

// @Summary 更新标签
// @Param id path int true "标签ID"
// @Param name query string false "标签名称" minlength(3) maxlength(100)
// @Param state query int false "状态" Enums(0, 1) default(1)
// @Param created_by body string true "修改者" minlength(3) maxlength(100)
// @Produce json
// @Success 200 {object} model.Tag "成功"
// @Router /api/v1/tags/{id} [put]
func (t Tag) Update(c *gin.Context) {

}

// @Summary 删除标签
// @Param id path int true "标签ID"
// @Produce json
// @Success 200 {object} model.Tag "成功"
// @Router /api/v1/tags/{id} [delete]
func (t Tag) Delete(c *gin.Context) {

}
