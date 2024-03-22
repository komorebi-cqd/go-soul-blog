package v1

import (
	"github.com/gin-gonic/gin"
)

type Diarys struct {
}

func NewDiarys() Diarys {
	return Diarys{}
}

// @Summary 获取单个说说详情
// @Param id path int true "说说ID"
// @Produce json
// @Success 200 {object} model.Diarys "成功"
// @Router /api/v1/diarys/{id} [get]
func (t Diarys) Get(c *gin.Context) {

}

// @Summary 获取说说列表
// @Param name query string false "说说名称" maxlength(100)
// @Param page query int false "页码"
// @Param page_size query int false "每页数量"
// @Produce json
// @Success 200 {object} model.DiarysSwagger "成功"
// @Router /api/v1/diarys [get]
func (t Diarys) List(c *gin.Context) {

}

// @Summary 新增说说
// @Param name query string true "说说名称" minlength(3) maxlength(100)
// @Param content body string true "内容" minlength(3) maxlength(1000)
// @Produce json
// @Success 200 {object} model.Diarys "成功"
// @Router /api/v1/diarys [post]
func (t Diarys) Create(c *gin.Context) {

}

// @Summary 更新说说
// @Param id path int true "说说ID"
// @Param name query string true "说说名称" minlength(3) maxlength(100)
// @Param content body string true "内容" minlength(3) maxlength(1000)
// @Produce json
// @Success 200 {object} model.Diarys "成功"
// @Router /api/v1/diarys/{id} [put]
func (t Diarys) Update(c *gin.Context) {

}

// @Summary 删除说说
// @Param id path int true "说说ID"
// @Produce json
// @Success 200 {object} model.Diarys "成功"
// @Router /api/v1/diarys/{id} [delete]
func (t Diarys) Delete(c *gin.Context) {

}
