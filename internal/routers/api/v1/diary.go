package v1

import (
	"github.com/gin-gonic/gin"
)

type Diary struct {
}

func NewDiary() Diary {
	return Diary{}
}

// @Summary 获取单个说说
// @Param id path int true "说说ID"
// @Produce json
// @Success 200 {object} model.Diary "成功"
// @Router /api/v1/diarys/{id} [get]
func (t Diary) Get(c *gin.Context) {

}

// @Summary 获取说说列表
// @Param name query string false "说说名称" maxlength(100)
// @Param state query int false "状态" Enums(0, 1) default(1)
// @Param page query int false "页码"
// @Param page_size query int false "每页数量"
// @Produce json
// @Success 200 {object} model.DiarysSwagger "成功"
// @Router /api/v1/diarys [get]
func (t Diary) List(c *gin.Context) {

}

// @Summary 新增说说
// @Param name query string true "说说名称" minlength(3) maxlength(100)
// @Param state query int false "状态" Enums(0, 1) default(1)
// @Param created_by body string true "创建者" minlength(3) maxlength(100)
// @Produce json
// @Success 200 {object} model.Diarys "成功"
// @Router /api/v1/diarys [post]
func (t Diary) Create(c *gin.Context) {

}

// @Summary 更新说说
// @Param id path int true "说说ID"
// @Param state path int false "状态" Enums(0, 1) default(1)
// @Param name query string false "说说名称" minlength(3) maxlength(100)
// @Produce json
// @Success 200 {object} model.Article "成功"
// @Router /api/v1/articles/{id} [put]
func (t Diary) Update(c *gin.Context) {

}

// @Summary 删除文章
// @Param id path int true "文章ID"
// @Produce json
// @Success 200 {object} model.Diarys "成功"
// @Router /api/v1/diarys/{id} [delete]
func (t Diary) Delete(c *gin.Context) {

}
