package service

import (
	"github.com/go-soul-blog/internal/model"
	"github.com/go-soul-blog/pkg/app"
)

type CountTagRequest struct {
	Name string `form:"name" binding:"min=1,max=100"`
}

type TagListRequest struct {
	Name string `form:"name" binding:"max=100"`
}

type CreateTagRequest struct {
	Name string `form:"name" binding:"required,min=1,max=100"`
}

type UpdateTagRequest struct {
	ID   uint32 `form:"id" binding:"required,gte=1"`
	Name string `form:"name" binding:"required,min=1,max=100"`
}

type DeleteTagRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
}

func (svc *Service) CountTag(param *CountTagRequest) (int64, error) {
	return svc.dao.CountTag(param.Name)
}

func (svc *Service) GetTagList(param *TagListRequest, pager *app.Pager) ([]*model.Tags, error) {
	return svc.dao.GetTagList(param.Name, pager.Page, pager.PageSize)
}
func (svc *Service) CreateTag(param *CreateTagRequest) error {
	return svc.dao.CreateTag(param.Name)
}

func (svc *Service) UpdateTag(param *UpdateTagRequest) error {
	return svc.dao.UpdateTag(param.ID, param.Name)
}

func (svc *Service) DeleteTag(param *DeleteTagRequest) error {
	return svc.dao.DeleteTag(param.ID)
}
