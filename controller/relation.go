package controller

import (
	"douyin/common"
	"douyin/model"
	"douyin/service"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UserListResponse struct {
	model.Response
	UserList []model.User `json:"user_list"`
}

func RelationAction(c *gin.Context) {
	token := c.Query("token")
	toUserId := c.Query("to_user_id")
	actionType := c.Query("action_type")
	userJson := common.Rdb.Get(c, common.UserLoginPrefix+token).Val()
	if userJson == "" {
		c.JSON(http.StatusOK, model.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}
	var tUser model.TUser
	json.Unmarshal([]byte(userJson), &tUser)
	toUid, _ := strconv.ParseInt(toUserId, 10, 64)
	if actionType == "1" {
		common.Rdb.Incr(c, common.UserFollowCountPrefix+strconv.FormatInt(tUser.ID, 10))
		common.Rdb.Incr(c, common.UserFollowerCountPrefix+toUserId)

		res := service.AddFollow(tUser.ID, toUid)
		if !res {
			c.JSON(http.StatusOK, model.Response{
				StatusCode: common.FollowActionError,
				StatusMsg:  common.RespMsg[common.FollowActionError],
			})
			return
		}
		c.JSON(http.StatusOK, model.Response{
			StatusCode: common.OK,
			StatusMsg:  common.RespMsg[common.OK],
		})
		return
	} else if actionType == "2" {
		res := service.CancelFollow(tUser.ID, toUid)
		if !res {
			c.JSON(http.StatusOK, model.Response{
				StatusCode: common.FollowCancelError,
				StatusMsg:  common.RespMsg[common.FollowCancelError],
			})
			return
		}
		c.JSON(http.StatusOK, model.Response{
			StatusCode: common.OK,
			StatusMsg:  common.RespMsg[common.OK],
		})
		return
	}
}
func FollowList(c *gin.Context) {
	token := c.Query("token")
	userJson := common.Rdb.Get(c, common.UserLoginPrefix+token).Val()
	if userJson == "" {
		c.JSON(http.StatusOK, model.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}
	var tUser model.TUser
	json.Unmarshal([]byte(userJson), &tUser)
	followList := service.GetFollowListByUserId(tUser.ID)
	c.JSON(http.StatusOK, UserListResponse{
		Response: model.Response{
			StatusCode: common.OK,
			StatusMsg:  common.RespMsg[common.OK],
		},
		UserList: followList,
	})
}

func FollowerList(c *gin.Context) {
	token := c.Query("token")
	userJson := common.Rdb.Get(c, common.UserLoginPrefix+token).Val()
	if userJson == "" {
		c.JSON(http.StatusOK, model.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}
	var tUser model.TUser
	json.Unmarshal([]byte(userJson), &tUser)
	followerList := service.GetFollowerListByUserId(tUser.ID)
	c.JSON(http.StatusOK, UserListResponse{
		Response: model.Response{
			StatusCode: common.OK,
			StatusMsg:  common.RespMsg[common.OK],
		},
		UserList: followerList,
	})
}
