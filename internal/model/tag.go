package model

import "gorm.io/gorm"

type Tags struct {
	*Model
	Name string `json:"name"`
}

type TagsList struct {
	ID        uint32
	CreatedAt int64
	UpdatedAt int64
	Name      string
}

type TagsSwagger struct {
	List []*Tags
}

func (t Tags) TableName() string {
	return "blog_tag"
}

func (t Tags) Count(db *gorm.DB) (int64, error) {
	var count int64
	if t.Name != "" {
		db = db.Where("name = ?", t.Name)
	}
	err := db.Model(&t).Count(&count).Error

	if err != nil {
		return 0, err
	}

	return count, nil
}

func (t Tags) List(db *gorm.DB, pageOffset, pageSize int) ([]*Tags, error) {
	var tags []*Tags
	var err error
	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}
	if t.Name != "" {
		db = db.Where("name = ?", t.Name)
	}
	if err = db.Find(&tags).Error; err != nil {
		return nil, err
	}
	return tags, nil
}

func (t Tags) Create(db *gorm.DB) error {

	return db.Create(&t).Error
}

func (t Tags) Update(db *gorm.DB) error {
	return db.Model(&Tags{}).Where("id = ?", t.ID).Updates(t).Error
}

func (t Tags) Delete(db *gorm.DB) error {
	return db.Where("id = ?", t.ID).Delete(&t).Error
}
