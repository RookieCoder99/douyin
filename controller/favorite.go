package controller

import (
	"douyin/common"
	"douyin/service"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

type favorite struct {
}

func FavoriteAction(c *gin.Context) {
	//UserId := new(model.TUser.Id)
	//VideoId := new(model.TVideo.Id)
	//var key = int(UserId + VideoId)
	//err := common.RespData(c, common.ErrNo(key), &key)
	//if err != nil {
	//	log.Printf("请求参数解析出错 %s", err)
	//
	//	return
	//}
	//log.Printf("点赞请求参数:%v", key)
	userId := c.Query("user_id")
	videoId := c.Query("video_id")
	var actionType, _ = strconv.ParseInt(c.Query("action_type"), 10, 64)

	if actionType != 1 && actionType != 2 { //状态参数不合法
		common.RespData(c, common.ParamInvalid, common.RespMsg[common.ParamInvalid])
		return
	}
	if actionType == 1 { //进入点赞操作
		//检查是否已经点过赞
		hasFavorite := common.Rdb.SIsMember(c, userId, videoId).Val()
		if hasFavorite {
			log.Printf("用户 %s 已经给 %s 点过赞", userId, videoId)

			return
		}
		service.Favorite(c, userId, videoId) //去service层点赞
	}

	if actionType == 2 { //进入取消点赞操作
		//检查是否取消点赞
		hashFavorite := common.Rdb.SIsMember(c, userId, videoId).Val()
		if hashFavorite == false {
			log.Printf("用户 %s 已经给 %s 取消过了点赞", userId, videoId)

			return
		}
		service.CancelFavorite(c, userId)
	}
}

func FavoriteList(c *gin.Context) {

}
