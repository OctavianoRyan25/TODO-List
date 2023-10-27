package router

import (
	"github.com/OctavianoRyan25/TODO-List/controller"
	"github.com/OctavianoRyan25/TODO-List/controller/auth"
	"github.com/OctavianoRyan25/TODO-List/middleware"
	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/auth")
	{
		userRouter.POST("/register", auth.UserRegister)
		userRouter.POST("/login", auth.UserLogin)
	}
	noteRouter := r.Group("/note")
	{
		noteRouter.Use(middleware.Authentication())
		noteRouter.GET("/", controller.Index)
		noteRouter.POST("/", controller.Store)
		noteRouter.PUT("/update/:noteId", middleware.Authorization(),controller.Update)
		noteRouter.GET("/note/:noteId", middleware.Authorization(),controller.GetByID)
		noteRouter.DELETE("/delete/:noteId", middleware.Authorization(),controller.Destroy)
	}

	return r
}