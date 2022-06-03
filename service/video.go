package service

import (
	"douyin/dao"
	"douyin/model"
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

func GetVideoListByUserId(userId int64) []model.Video {
	//1、使用用户名 + 密码 查询用户信息 如果存在用户 直接返回
	return dao.GetVideoListByUserId(userId)
}

func GetVideoList(timeUnix string) ([]model.Video, time.Time) {
	var videoList []model.Video
	var latestTime time.Time
	if timeUnix == "" {
		videoList, latestTime = dao.GetVideoList()
	} else {
		unix, _ := strconv.ParseInt(timeUnix, 10, 64)
		timeStr := time.UnixMilli(unix).Format("2006-01-02 15:04:05.999")
		videoList, latestTime = dao.GetVideoListByTime(timeStr)
	}

	return videoList, latestTime
}

//func CheckFavorite(userId int64, videoId int64) bool{
//
//}
