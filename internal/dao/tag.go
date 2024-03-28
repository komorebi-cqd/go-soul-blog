package dao

import (
	"github.com/go-soul-blog/internal/model"
	"github.com/go-soul-blog/pkg/app"
)

func (d *Dao) GetTagList(name string, page, pageSize int) ([]*model.Tags, error) {
	tag := model.Tags{Name: name}
	pageOffset := app.GetPageOffset(page, pageSize)
	return tag.List(d.engine, pageOffset, pageSize)
}

func (d *Dao) CreateTag(name string) error {
	tag := model.Tags{
		Name: name,
	}
	return tag.Create(d.engine)
}

func (d *Dao) UpdateTag(id uint32, name string) error {
	tag := model.Tags{
		Name:  name,
		Model: &model.Model{ID: id},
	}

	return tag.Update(d.engine)
}

func (d *Dao) DeleteTag(id uint32) error {
	tag := model.Tags{
		Model: &model.Model{ID: id},
	}

	return tag.Delete(d.engine)
}
