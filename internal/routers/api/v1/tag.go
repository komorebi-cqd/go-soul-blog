package v1

import "github.com/gin-gonic/gin"

type Tags struct{}

func NewTags() Tags {
	return Tags{}
}

// @Summary 获取所有标签
// @Produce json
// @Success 200 {object} model.TagsSwagger "成功"
// @Router /api/v1/tags [get]
func (t Tags) List(c *gin.Context) {

}

// @Summary 新增标签
// @Param name query string true "标签名称" minlength(3) maxlength(100)
// @Param state query int false "状态" Enums(0, 1) default(1)
// @Produce json
// @Success 200 {object} model.Tags "成功"
// @Router /api/v1/tags [post]
func (t Tags) Create(c *gin.Context) {

}

// @Summary 更新标签
// @Param id path int true "标签ID"
// @Param name query string false "标签名称" minlength(3) maxlength(100)
// @Param state query int false "状态" Enums(0, 1) default(1)
// @Param created_by body string true "修改者" minlength(3) maxlength(100)
// @Produce json
// @Success 200 {object} model.Tags "成功"
// @Router /api/v1/tags/{id} [put]
func (t Tags) Update(c *gin.Context) {

}

// @Summary 删除标签
// @Param id path int true "标签ID"
// @Produce json
// @Success 200 {object} model.Tags "成功"
// @Router /api/v1/tags/{id} [delete]
func (t Tags) Delete(c *gin.Context) {

}
