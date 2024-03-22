package model

type Tags struct {
	*Model
	Name  string `json:"name"`
	State uint8  `json:"state"`
}

type TagsSwagger struct {
	List []*Tags
}

func (t Tags) TableName() string {
	return "tags"
}
