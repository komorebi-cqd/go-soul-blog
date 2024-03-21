package routers

import (
	v1 "github.com/go-soul-blog/internal/routers/api/v1"

	"github.com/gin-gonic/gin"
	_ "github.com/go-soul-blog/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	user := v1.NewTag()
	article := v1.NewArticle()
	tag := v1.NewTag()

	apiv1 := r.Group("/api/v1")
	{
		apiv1.POST("/user", user.Create)       //创建单个用户
		apiv1.DELETE("/user/:id", user.Delete) //删除单个用户
		apiv1.PUT("/user/:id", user.Update)    // 更新单个用户
		apiv1.GET("/users", user.List)         //获取用户列表

		apiv1.POST("/tags", tag.Create)       //创建单个标签
		apiv1.DELETE("/tags/:id", tag.Delete) //删除单个标签
		apiv1.PUT("/tags/:id", tag.Update)    // 更新单个标签
		apiv1.GET("/tags", tag.List)          //获取标签列表

		apiv1.POST("/articles", article.Create)       //创建文章
		apiv1.DELETE("/articles/:id", article.Delete) //删除单个文章
		apiv1.PUT("/articles/:id", article.Update)    // 更新单个文章
		apiv1.GET("/articles/:id", article.Get)       //获取单个文章
		apiv1.GET("/articles", article.List)          //获取文章列表
	}

	return r
}
