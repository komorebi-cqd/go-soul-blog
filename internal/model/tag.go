package model

type Tag struct {
	*Model
	Name  string `json:"name"`
	State uint8  `json:"state"`
}

type TagSwagger struct {
	List []*Tag
}

func (t Tag) TableName() string {
	return "blog_tag"
}
