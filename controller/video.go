package controller

import (
	"douyin/common"
	"douyin/model"
	"douyin/service"
	"douyin/utils"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"time"
)

type VideoListResponse struct {
	model.Response
	VideoList []model.Video `json:"video_list"`
}

type FeedRequest struct {
	LatestTime int64 // 可选参数，限制返回视频的最新投稿时间戳，精确到s， 不填表示当前时间
}

type FeedResponse struct {
	model.Response
	VideoList []model.Video `json:"video_list,omitempty"`
	NextTime  int64         `json:"next_time,omitempty"`
}

func Feed(c *gin.Context) {
	latestTime := c.Query("latest_time")
	latestTimeVal, _ := strconv.ParseInt(latestTime, 10, 64)
	if latestTimeVal > time.Now().UnixMilli() {
		latestTimeVal = time.Now().UnixMilli()
	}
	videoList, timeRes := service.GetVideoList(strconv.FormatInt(latestTimeVal, 10))

	c.JSON(http.StatusOK, FeedResponse{
		Response:  model.Response{StatusCode: 0},
		VideoList: videoList,
		NextTime:  timeRes.Unix(),
	})
}
func PublishAction(c *gin.Context) {
	//token := c.Query("token")
	//uId := c.Query("user_id")
	//token, _ := c.GetPostForm("user_id")
	title, _ := c.GetPostForm("title")
	token, _ := c.GetPostForm("token")
	log.Println(token)
	//log.Println(uId)
	userJson := common.Rdb.Get(c, common.UserLoginPrefix+token).Val()
	if userJson == "" {
		c.JSON(http.StatusOK, model.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}
	var tUser model.TUser
	json.Unmarshal([]byte(userJson), &tUser)

	data, err := c.FormFile("data")

	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: common.UnknownError,
			StatusMsg:  err.Error(),
		})
		return
	}

	//filename := filepath.Base(data.Filename)
	//finalName := fmt.Sprintf("%d_%s", tUser.ID, filename)
	data.Filename = fmt.Sprintf("%d_%d_%s", tUser.ID, time.Now().UnixNano(), data.Filename)
	status, videoUrl, coverUrl := utils.UploadToQiNiu(data)
	if status != 0 {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: common.UploadError,
			StatusMsg:  data.Filename + common.RespMsg[common.UploadError],
		})
	}
	//TODO 封面url未设置
	service.UploadVideo(videoUrl, tUser.ID, coverUrl, title)
	log.Printf(" 上传状态 %v", status)

	c.JSON(http.StatusOK, model.Response{
		StatusCode: common.OK,
		StatusMsg:  data.Filename + " uploaded successfully",
	})

}

func PublishList(c *gin.Context) {
	token := c.Query("token")
	userJson := common.Rdb.Get(c, common.UserLoginPrefix+token).Val()
	if userJson == "" {
		c.JSON(http.StatusOK, model.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}
	var tUser model.TUser
	json.Unmarshal([]byte(userJson), &tUser)

	videoList := service.GetVideoListByUserId(tUser.ID)

	c.JSON(http.StatusOK, VideoListResponse{
		Response: model.Response{
			StatusCode: 0,
		},
		VideoList: videoList,
	})
}
