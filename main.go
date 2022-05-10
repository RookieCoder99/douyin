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
	common.Config = config.InitConfig(parse(os.Args))

	// 初始化数据库等连接
	common.Init()

	router := gin.Default()
	initRouter(router)

	router.Run(common.Config.Server.Port)
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
	ug := g.Group("/douyin")
	ug.GET("/feed", controller.Feed)
	ug.POST("/user/login", controller.Login)
	ug.POST("/user/logout", controller.Logout)
	ug.POST("/user/register", controller.Register)
	ug.GET("/user", controller.UserInfo)

	ug.POST("/publish/action", controller.PublishAction)
	ug.GET("/publish/action", controller.PublishAction)

	ug.POST("/favorite/action", controller.FavoriteAction)
	ug.GET("/favorite/list", controller.FavoriteList)

	ug.POST("/comment/action", controller.CommentAction)
	ug.GET("/comment/list", controller.CommentList)

	ug.POST("/relation/action", controller.RelationAction)
	ug.GET("/relation/follow/list", controller.FollowList)
	ug.GET("/relation/follower/list", controller.FollowerList)

}
