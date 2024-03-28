package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-soul-blog/pkg/errcode"
)

type Response struct {
	Ctx *gin.Context
}

type Pager struct {
	Page      int `json:"page"`
	PageSize  int `json:"page_size"`
	TotalRows int `json:"total_rows"`
}

func NewResponse(ctx *gin.Context) *Response {
	return &Response{Ctx: ctx}
}

// 普通返回
func (r *Response) ToResponse(data interface{}) {
	if data == nil {
		data = gin.H{}
	}

	r.Ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": nil,
		"msg":  "成功",
	})
}

// 列表返回
func (r *Response) ToResponseList(list interface{}, totalRows int) {
	r.Ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "成功",
		"data": gin.H{
			"list": list,
			"pager": Pager{
				Page:      GetPage(r.Ctx),
				PageSize:  GetPageSize(r.Ctx),
				TotalRows: totalRows,
			},
		},
	})
}

func (r *Response) ToErrorResponse(err *errcode.Error) {
	response := gin.H{"code": err.Code(), "msg": err.Msg()}
	details := err.Details()
	if len(details) > 0 {
		response["details"] = details
	}
	r.Ctx.JSON(err.StatusCode(), response)
}
