package model

type Tags struct {
	*Model
	Name string `json:"name"`
}

type TagsSwagger struct {
	List []*Tags
}

func (t Tags) TableName() string {
	return "blog_tags"
}
