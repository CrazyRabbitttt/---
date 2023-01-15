package route

import (
	"Web-Go/controller"
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	//	公共目录用来服务静态的资源

	r.Static("/static", "./public")

	apiRouter := r.Group("/douyin")

	//basic apis

	//apiRouter.GET("/feed/", controller.Feed) //feed流的handler
	apiRouter.POST("/user/register/", controller.UserRegister)
	apiRouter.POST("/user/login/", controller.UserLogin)
	apiRouter.GET("/user/", controller.UserInfo)
	apiRouter.POST("/publish/action/", controller.Publish)
	//apiRouter.GET("/publish/list/", controller.PublishList)
}
