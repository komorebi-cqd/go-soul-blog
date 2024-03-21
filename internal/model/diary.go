package model

import "github.com/go-soul-blog/pkg/app"

type Diarys struct {
	*Model
	Title   string `json:"title"`
	Content string `json:"content"`
}

type DiarysSwagger struct {
	List  []*Diarys
	Pager *app.Pager
}

func (a Diarys) TableName() string {
	return "diarys"
}
