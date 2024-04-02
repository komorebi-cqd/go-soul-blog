package model

import (
	"github.com/go-soul-blog/pkg/app"
	"gorm.io/gorm"
)

type Users struct {
	*Model
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	State    uint8  `json:"state"`
	Avater   string `json:"avater"`
	Desc     string `json:"desc"`
}

type UsersSwagger struct {
	List  []*Users
	Pager *app.Pager
}

func (t Users) TableName() string {
	return "blog_users"
}

func (t Users) Count(db *gorm.DB) (int64, error) {
	var count int64
	if t.Username != "" {
		db = db.Where("username = ?", t.Username)
	}
	err := db.Model(&t).Count(&count).Error

	if err != nil {
		return 0, err
	}

	return count, nil
}

func (t Users) List(db *gorm.DB, pageOffset, pageSize int) ([]*Users, error) {
	var users []*Users
	var err error
	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}
	if t.Username != "" {
		db = db.Where("username = ?", t.Username)
	}
	if err = db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (t Users) Create(db *gorm.DB) error {
	return db.Create(&t).Error
}

func (t Users) Update(db *gorm.DB) error {
	return db.Model(&Tags{}).Where("id = ?", t.ID).Updates(t).Error
}

func (t Users) Delete(db *gorm.DB) error {
	return db.Where("id = ?", t.ID).Delete(&t).Error
}
