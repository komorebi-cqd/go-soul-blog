package v1

import (
	"github.com/gin-gonic/gin"
)

type Likes struct {
}

func NewLikes() Likes {
	return Likes{}
}

// @Summary 获取单个说说
// @Param postId path int true "说说或文章ID"
// @Param type query string true "类型 blog是图文 comment是评论"
// @Produce json
// @Success 200 {object} model.Likes "成功"
// @Router /api/v1/likes/{postId} [get]
func (t Likes) Get(c *gin.Context) {

}

// @Summary 获取点赞列表
// @Param page query int false "页码"
// @Param page_size query int false "每页数量"
// @Produce json
// @Success 200 {object} model.LikesSwagger "成功"
// @Router /api/v1/likes [get]
func (t Likes) List(c *gin.Context) {

}

// @Summary 点赞
// @Param id path int true "说说ID"
// @Param type query string true "类型 blog是图文 comment是评论"
// @Produce json
// @Success 200 {object} model.Likes "成功"
// @Router /api/v1/likes/{id} [put]
func (t Likes) Update(c *gin.Context) {

}
