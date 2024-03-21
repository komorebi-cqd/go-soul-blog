package model

import "github.com/go-soul-blog/pkg/app"

type User struct {
	*Model
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	State     uint8  `json:"state"`
	Avater    string `json:"avater"`
	Introduce string `json:"introduce"`
}

type UserSwagger struct {
	List  []*User
	Pager *app.Pager
}

func (t User) TableName() string {
	return "user"
}
