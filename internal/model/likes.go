package model

import "github.com/go-soul-blog/pkg/app"

type Type string

const (
	Blog    Type = "blog"
	Comment Type = "comment"
)

type Links struct {
	*Model
	UserId uint32 `json:"user_id"`
	PostId uint32 `json:"post_id"`
	Type   Type   `json:"type"`
}

type LinksSwagger struct {
	List  []*Links
	Pager *app.Pager
}

func (a Links) TableName() string {
	return "links"
}
