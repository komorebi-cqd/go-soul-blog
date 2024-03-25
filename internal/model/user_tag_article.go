package model

// 这是把用户ID和文章ID关联的表

type UserTagArticle struct {
	*Model
	UserId    uint32 `json:"user_id"`
	TagId     uint32 `json:"tag_id"`
	ArticleId uint32 `json:"article_id"`
}

func (a UserTagArticle) TableName() string {
	return "blog_user_tag_article"
}
