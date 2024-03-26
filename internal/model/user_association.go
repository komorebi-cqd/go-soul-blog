package model

// 这是把用户关联他作品的表

type UserAssociation struct {
	*Model
	UserId    uint32 `json:"user_id"`
	DiaryId   uint32 `json:"diary_id"`
	ArticleId uint32 `json:"article_id"`
}

func (a UserAssociation) TableName() string {
	return "blog_user_tag_article"
}
