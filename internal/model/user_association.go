package model

type UserAssociation struct {
	*Model
	UserId    uint32 `json:"user_id"`
	TagId     uint32 `json:"tag_id"`
	ArticleId uint32 `json:"article_id"`
}

func (a UserAssociation) TableName() string {
	return "blog_user_association"
}
