package model

import "github.com/go-soul-blog/pkg/app"

type Diarys struct {
	*Model
	Title   string `json:"title"`
	Content string `json:"content"`
	UserId  string `json:"user_id"`
}

type DiarysSwagger struct {
	List  []*Diarys
	Pager *app.Pager
}

func (a Diarys) TableName() string {
	return "blog_diaries"
}
