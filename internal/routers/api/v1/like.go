package v1

import (
	"github.com/gin-gonic/gin"
)

type Likes struct {
}

func NewLikes() Likes {
	return Likes{}
}

// @Summary 通过博客或说说ID获取点赞列表
// @Param postId path int true "博客或说说ID"
// @Param type query string true "Enums(blog, common) default(blog)"
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
// @Router /api/v1/likes/{id} [post]
func (t Likes) Create(c *gin.Context) {

}
