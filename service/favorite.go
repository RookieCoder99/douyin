package service

import (
	"douyin/common"
	"douyin/dao"
	"douyin/model"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

func InsertFavorite(c *gin.Context, userId, videoId int64) bool {
	var favorite = model.TFavorite{
		VideoID: videoId,
		UserID:  userId,
	}
	vId := strconv.FormatInt(videoId, 10)
	// redis增加点赞数量
	common.Rdb.Incr(c, common.VideoFavoriteCountPrefix+vId)

	// 用户redis增加点赞视频
	uId := strconv.FormatInt(userId, 10)
	common.Rdb.SAdd(c, common.UserFavoriteVideoPrefix+uId, vId)

	// 用户点赞过的作品数量+1
	common.Rdb.Incr(c, common.UserFavoriteCountPrefix+uId)

	videoUserId := dao.GetUserIdByVideoId(videoId)
	vuId := strconv.FormatInt(videoUserId, 10)
	// 被点赞用户赞的数量+1
	common.Rdb.Incr(c, common.UserTotalFavoritedPrefix+vuId)

	return dao.InsertFavorite(&favorite)
}

func DeleteFavorite(c *gin.Context, userId, videoId int64) bool {
	// redis减少点赞数量
	vId := strconv.FormatInt(videoId, 10)
	common.Rdb.Decr(c, common.VideoFavoriteCountPrefix+vId)
	// 用户redis增加点赞视频
	uId := strconv.FormatInt(userId, 10)
	common.Rdb.SRem(c, common.UserFavoriteVideoPrefix+uId, vId)

	common.Rdb.Decr(c, common.UserFavoriteCountPrefix+uId)

	videoUserId := dao.GetUserIdByVideoId(videoId)
	vuId := strconv.FormatInt(videoUserId, 10)
	// 被点赞用户赞的数量+1
	common.Rdb.Decr(c, common.UserTotalFavoritedPrefix+vuId)

	return dao.DeleteFavorite(userId, videoId)
}

func GetFavoriteList(c *gin.Context, userId int64) []*model.Video {
	uId := strconv.FormatInt(userId, 10)
	vIds, err := common.Rdb.SMembers(c, common.UserFavoriteVideoPrefix+uId).Result()
	if err != nil {
		log.Println(err)
		return nil
	}
	var videoList = make([]int64, len(vIds))
	for idx, id := range vIds {
		k, _ := strconv.ParseInt(id, 10, 64)
		videoList[idx] = k
	}

	return dao.GetVideosByIds(videoList)
}
