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

	users := v1.NewUsers()
	articles := v1.NewArticles()
	tags := v1.NewTags()
	diarys := v1.NewDiarys()
	links := v1.NewLikes()

	upload := v1.NewUpload()
	r.POST("/upload/file", upload.UploadFile)

	apiv1 := r.Group("/api/v1")
	{
		apiv1.POST("/users", users.Create)       //创建单个用户
		apiv1.DELETE("/users/:id", users.Delete) //删除单个用户
		apiv1.PUT("/users", users.Update)        // 更新单个用户
		apiv1.GET("/users", users.List)          //获取用户列表

		apiv1.POST("/tags", tags.Create)       //创建单个标签
		apiv1.DELETE("/tags/:id", tags.Delete) //删除单个标签
		apiv1.PUT("/tags", tags.Update)        // 更新单个标签
		apiv1.GET("/tags", tags.List)          //获取标签列表

		apiv1.POST("/articles", articles.Create)       //创建文章
		apiv1.DELETE("/articles/:id", articles.Delete) //删除单个文章
		apiv1.PUT("/articles", articles.Update)        // 更新单个文章
		apiv1.GET("/articles/:id", articles.Get)       //获取单个文章
		apiv1.GET("/articles", articles.List)          //获取文章列表

		apiv1.POST("/diarys", diarys.Create)       //创建说说
		apiv1.DELETE("/diarys/:id", diarys.Delete) //删除单个说说
		apiv1.PUT("/diarys", diarys.Update)        // 更新单个说说
		apiv1.GET("/diarys/:id", diarys.Get)       //获取单个说说
		apiv1.GET("/diarys", diarys.List)          //获取说说列表

		apiv1.GET("/links", links.List)       //通过blogid获取当前点赞列表
		apiv1.GET("/links/:id", links.Create) //点赞
	}

	return r
}
