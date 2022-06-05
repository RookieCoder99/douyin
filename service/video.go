package service

import (
	"douyin/common"
	"douyin/dao"
	"douyin/model"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

func UploadVideo(url string, userId int64, coverUrl string, title string) *model.TVideo {
	var video = model.TVideo{
		AuthorID: userId,
		PlayUrl:  url,
		CoverUrl: coverUrl,
		Title:    title,
	}
	v := dao.InsertVideo(&video)

	return v
}

func GetVideoListByUserId(c *gin.Context, userId int64) []*model.Video {
	videoList := dao.GetVideoListByUserId(userId)
	return AddAdditionInfo(c, videoList, userId)
}

func GetVideoList(c *gin.Context, timeUnix string) ([]*model.Video, time.Time) {
	var videoList []*model.Video
	var latestTime time.Time
	if timeUnix == "" {
		videoList, latestTime = dao.GetVideoList()
	} else {
		unix, _ := strconv.ParseInt(timeUnix, 10, 64)
		timeStr := time.UnixMilli(unix).Format("2006-01-02 15:04:05.999")
		videoList, latestTime = dao.GetVideoListByTime(timeStr)
	}
	videoList = AddAdditionInfo(c, videoList, -1)
	return videoList, latestTime
}

func AddAdditionInfo(c *gin.Context, videoList []*model.Video, userId int64) []*model.Video {
	for _, video := range videoList {
		uId := strconv.FormatInt(video.User.Id, 10)
		tf, _ := strconv.ParseInt(common.Rdb.Get(c, common.UserTotalFavoritedPrefix+uId).Val(), 10, 64)
		video.User.TotalFavorited = tf
		fc, _ := strconv.ParseInt(common.Rdb.Get(c, common.UserFavoriteCountPrefix+uId).Val(), 10, 64)
		video.User.FavoriteCount = fc

		followCount, _ := strconv.ParseInt(common.Rdb.Get(c, common.UserFollowCountPrefix+uId).Val(), 10, 64)
		video.User.FollowCount = followCount

		followerCount, _ := strconv.ParseInt(common.Rdb.Get(c, common.UserFollowerCountPrefix+uId).Val(), 10, 64)
		video.User.FollowerCount = followerCount

		video.User.IsFollow = dao.CheckRelation(video.User.Id, userId)

	}
	return videoList
}

//func CheckFavorite(userId int64, videoId int64) bool{
//
//}
