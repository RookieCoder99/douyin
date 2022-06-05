package controller

import (
	"douyin/common"
	"douyin/model"
	"douyin/service"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func FavoriteAction(c *gin.Context) {
	token := c.Query("token")
	videoId := c.Query("video_id")
	actionTYpe := c.Query("action_type")
	userJson := common.Rdb.Get(c, common.UserLoginPrefix+token).Val()
	if userJson == "" {
		c.JSON(http.StatusOK, model.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}
	var tUser model.TUser
	json.Unmarshal([]byte(userJson), &tUser)

	vId, _ := strconv.ParseInt(videoId, 10, 64)

	if actionTYpe == "1" {
		res := service.InsertFavorite(c, tUser.ID, vId)
		if !res {
			log.Printf("%v 视频点赞失败", vId)
			c.JSON(http.StatusOK, model.Response{
				StatusCode: common.FavoriteActionError,
				StatusMsg:  common.RespMsg[common.FavoriteActionError],
			})
			return
		}
		log.Printf("%v 视频点赞成功", vId)
		c.JSON(http.StatusOK, model.Response{
			StatusCode: common.OK,
			StatusMsg:  common.RespMsg[common.OK],
		})
	} else if actionTYpe == "2" {
		res := service.DeleteFavorite(c, tUser.ID, vId)

		if !res {
			log.Printf("%v 视频取消点赞失败", vId)
			c.JSON(http.StatusOK, model.Response{
				StatusCode: common.FavoriteCancelError,
				StatusMsg:  common.RespMsg[common.FavoriteCancelError],
			})
			return
		}
		log.Printf("%v 视频取消点赞成功", vId)
		c.JSON(http.StatusOK, model.Response{
			StatusCode: common.OK,
			StatusMsg:  common.RespMsg[common.OK],
		})

	}

}

func FavoriteList(c *gin.Context) {
	uuid := c.Query("token")
	userId := c.Query("user_id")
	fmt.Println(userId)

	userJson := common.Rdb.Get(c, common.UserLoginPrefix+uuid).Val()
	if userJson == "" {
		c.JSON(http.StatusOK, model.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}
	//dec := json.NewDecoder(strings.NewReader(userJson))\
	var tUser model.TUser
	json.Unmarshal([]byte(userJson), &tUser)
	favList := service.GetFavoriteList(c, tUser.ID)
	c.JSON(http.StatusOK, VideoListResponse{
		Response: model.Response{
			StatusCode: common.OK,
			StatusMsg:  common.RespMsg[common.OK],
		},
		VideoList: favList,
	})
}
