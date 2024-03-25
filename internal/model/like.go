package model

import "github.com/go-soul-blog/pkg/app"

type Type string

const (
	Blog    Type = "blog"
	Comment Type = "comment"
)

type Likes struct {
	*Model
	UserId uint32 `json:"user_id"`
	PostId uint32 `json:"post_id"`
	Type   Type   `json:"type"`
}

type LikesSwagger struct {
	List  []*Likes
	Pager *app.Pager
}

func (a Likes) TableName() string {
	return "blog_likes"
}
