package model

import "github.com/go-soul-blog/pkg/app"

type Articles struct {
	*Model
	Title         string `json:"title"`
	Desc          string `json:"desc"`
	Content       string `json:"content"`
	CoverImageUrl string `json:"cover_image_url"`
	State         uint8  `json:"state"`
	IsDraft       uint8  `json:"is_draft"`
	UserId        string `json:"user_id"`
}

type ArticlesSwagger struct {
	List  []*Articles
	Pager *app.Pager
}

func (a Articles) TableName() string {
	return "blog_articles"
}
