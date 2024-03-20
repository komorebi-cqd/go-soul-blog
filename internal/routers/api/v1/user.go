package v1

import "github.com/gin-gonic/gin"

type User struct{}

func NewUser() Tag {
	return Tag{}
}

func (t User) Get(c *gin.Context) {

}

// @Summary 获取用户列表
// @Param name query string false "用户名称" maxlength(100)
// @Param state query int false "状态" Enums(0, 1) default(1)
// @Param page query int false "页码"
// @Param page_size query int false "每页数量"
// @Produce json
// @Success 200 {object} model.TagSwagger "成功"
// @Router /api/v1/users [get]
func (t User) List(c *gin.Context) {

}

// @Summary 新增用户
// @Param name body string true "用户名称" minlength(1) maxlength(100)
// @Param name body file true "头像文件"
// @Param state body int false "状态" Enums(0, 1) default(1)
// @Produce json
// @Success 200 {object} model.Tag "成功"
// @Router /api/v1/user [post]
func (t User) Create(c *gin.Context) {

}

// @Summary 更新标签
// @Param id path int true "用户ID"
// @Param name body string true "用户名称" minlength(1) maxlength(100)
// @Param state query int false "状态" Enums(0, 1) default(1)
// @Param name body file true "头像文件"
// @Produce json
// @Success 200 {object} model.Tag "成功"
// @Router /api/v1/user/{id} [put]
func (t User) Update(c *gin.Context) {

}

// @Summary 删除标签
// @Param id path int true "标签ID"
// @Produce json
// @Success 200 {object} model.Tag "成功"
// @Router /api/v1/user/{id} [delete]
func (t User) Delete(c *gin.Context) {

}
