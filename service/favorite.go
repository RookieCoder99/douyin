package service

import (
	"douyin/dao"
	"github.com/gin-gonic/gin"
)

func Favorite(c *gin.Context, userid string, videoid string) {
	dao.AddFavorite(c, userid, videoid)
}
func CancelFavorite(c *gin.Context, userid string) {
	dao.DeleteFavorite(c, userid)
}
