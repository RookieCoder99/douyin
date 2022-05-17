package dao

import (
	"douyin/common"
	"github.com/gin-gonic/gin"
)

func AddFavorite(c *gin.Context, userid string, videoid string) {
	common.Rdb.SAdd(c, userid, videoid)
}
func DeleteFavorite(c *gin.Context, userid string) {
	common.Rdb.Del(c, userid)
}
