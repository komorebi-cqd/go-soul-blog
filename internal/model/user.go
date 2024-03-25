package model

import "github.com/go-soul-blog/pkg/app"

type Users struct {
	*Model
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	State     uint8  `json:"state"`
	Avater    string `json:"avater"`
	Introduce string `json:"introduce"`
}

type UsersSwagger struct {
	List  []*Users
	Pager *app.Pager
}

func (t Users) TableName() string {
	return "blog_users"
}
