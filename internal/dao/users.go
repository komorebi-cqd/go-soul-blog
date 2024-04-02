package dao

import (
	"github.com/go-soul-blog/internal/model"
	"github.com/go-soul-blog/pkg/app"
)

func (d *Dao) CountUsers(username string) (int64, error) {
	users := model.Users{Username: username}
	return users.Count(d.engine)
}

func (d *Dao) GetUsersList(username string, page, pageSize int) ([]*model.Users, error) {
	users := model.Users{Username: username}
	pageOffset := app.GetPageOffset(page, pageSize)
	return users.List(d.engine, pageOffset, pageSize)
}

func (d *Dao) CreateUsers(username string) error {

	users := model.Users{
		Username: username,
	}

	return users.Create(d.engine)
}

func (d *Dao) UpdateUsers(id uint32, username string) error {
	users := model.Users{
		Username: username,
		Model:    &model.Model{ID: id},
	}
	return users.Update(d.engine)
}

func (d *Dao) DeleteUsers(id uint32) error {
	users := model.Users{
		Model: &model.Model{ID: id},
	}

	return users.Delete(d.engine)
}
