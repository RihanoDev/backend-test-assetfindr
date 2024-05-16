package routes

import (
	"post-api/controller"

	"github.com/gin-gonic/gin"
)

func SetupRoute(server *gin.Engine) {

	//routes posts
	server.GET("/api/posts", controller.GetAllPost)
	server.GET("/api/posts/:id", controller.GetPostByID)
	server.POST("/api/posts", controller.CreatePost)
	server.PUT("/api/posts/:id", controller.UpdatePost)
	server.DELETE("/api/posts/:id", controller.DeletePost)

	//routes tags
	server.GET("/api/tags", controller.GetAllTag)
	server.GET("/api/tags/:id", controller.GetTagByID)
	server.POST("/api/tags", controller.CreateTag)
	server.PUT("/api/tags/:id", controller.UpdateTag)
	server.DELETE("/api/tags/:id", controller.DeleteTag)
}
