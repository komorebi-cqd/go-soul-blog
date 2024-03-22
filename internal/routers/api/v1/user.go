package v1

import "github.com/gin-gonic/gin"

type Users struct{}

func NewUsers() Users {
	return Users{}
}

func (t Users) Get(c *gin.Context) {

}

// @Summary 获取用户列表
// @Param name query string false "用户名称" maxlength(100)
// @Param state query int false "状态" Enums(0, 1) default(1)
// @Param page query int false "页码"
// @Param page_size query int false "每页数量"
// @Produce json
// @Success 200 {object} model.UsersSwagger "成功"
// @Router /api/v1/users [get]
func (t Users) List(c *gin.Context) {

}

// @Summary 新增用户
// @Param name body string true "用户名称" minlength(1) maxlength(100)
// @Param avater formData file true "头像文件"
// @Param state body int false "状态" Enums(0, 1) default(1)
// @Produce json
// @Success 200 {object} model.Users "成功"
// @Router /api/v1/users [post]
func (t Users) Create(c *gin.Context) {

}

// @Summary 更新用户信息
// @Param id path int true "用户ID"
// @Param name body string true "用户名称" minlength(1) maxlength(100)
// @Param state query int false "状态" Enums(0, 1) default(1)
// @Param avater formData file true "头像文件"
// @Produce json
// @Success 200 {object} model.Users "成功"
// @Router /api/v1/users/{id} [put]
func (t Users) Update(c *gin.Context) {

}

// @Summary 删除标签
// @Param id path int true "标签ID"
// @Produce json
// @Success 200 {object} model.Users "成功"
// @Router /api/v1/users/{id} [delete]
func (t Users) Delete(c *gin.Context) {

}
