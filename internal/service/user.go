package service

type SigleUserRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
}

type UserListRequest struct {
	Name  string `form:"name" binding:"max=100"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type CreateUserRequest struct {
	Name   string `form:"name" binding:"required,min=3,max=100"`
	State  uint8  `form:"state,default=1" binding:"oneof=0 1"`
	Avater uint8  `form:"avater" binding:"required"`
}

type UpdateUserRequest struct {
	ID     uint32 `form:"id" binding:"required,gte=1"`
	Name   string `form:"name" binding:"required,min=3,max=100"`
	State  uint8  `form:"state,default=1" binding:"oneof=0 1"`
	Avater uint8  `form:"avater" binding:"required"`
}

type DeleteUserRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
}
