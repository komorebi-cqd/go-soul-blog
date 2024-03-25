package service

type SigleDiaryRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
}

type DiaryListRequest struct {
	Name string `form:"name" binding:"max=100"`
}

type CreateDiaryRequest struct {
	Name    string `form:"name" binding:"required,min=3,max=100"`
	Content string `form:"name" binding:"required,min=3,max=100"`
}

type UpdateDiaryRequest struct {
	ID      uint32 `form:"id" binding:"required,gte=1"`
	Name    string `form:"name" binding:"required,min=3,max=100"`
	Content string `form:"name" binding:"required,min=3,max=100"`
}

type DeleteDiaryRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
}
