package main

import (
	"douyin/common"
	"douyin/config"
	"douyin/controller"
	"github.com/gin-gonic/gin"
	"os"
	"strings"
)

func main() {
	common.Config = config.InitConfig(parse(os.Args)) // func() -> int

	// 初始化数据库等连接
	common.Init()

	router := gin.Default()
	initRouter(router)

	router.Run()
}

// parse 解析参数
func parse(args []string) string {
	if len(args) > 1 {
		argArr := strings.Split(os.Args[1], "=")
		if len(argArr) == 2 && argArr[0] == "--config" {
			return argArr[1]
		}
	}
	return ""
}

func initRouter(g *gin.Engine) {
	g.GET("/hello", controller.Hello)
	apiRouter := g.Group("/douyin")

	// basic apis
	apiRouter.GET("/feed/", controller.Feed)
	apiRouter.GET("/user/", controller.UserInfo)
	apiRouter.POST("/user/register/", controller.Register)
	apiRouter.POST("/user/login/", controller.Login)
	apiRouter.POST("/publish/action/", controller.PublishAction)
	apiRouter.GET("/publish/list/", controller.PublishList)

	// extra apis - I
	apiRouter.POST("/favorite/action/", controller.FavoriteAction)
	apiRouter.GET("/favorite/list/", controller.FavoriteList)
	apiRouter.POST("/comment/action/", controller.CommentAction)
	apiRouter.GET("/comment/list/", controller.CommentList)

	// extra apis - II
	apiRouter.POST("/relation/action/", controller.RelationAction)
	apiRouter.GET("/relation/follow/list/", controller.FollowList)
	apiRouter.GET("/relation/follower/list/", controller.FollowerList)

}
