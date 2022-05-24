package controller

import (
	"douyin/common"
	"douyin/model"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserResponse struct {
	model.Response
	User model.User `json:"user"`
}

func UserInfo(c *gin.Context) {
	token := c.Query("token")
	//otherId := c.Query("user_id")
	userJson := common.Rdb.Get(c, common.UserLoginPrefix+token).Val()
	if userJson == "" {
		c.JSON(http.StatusOK, UserResponse{
			Response: model.Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}
	var tUser model.TUser
	json.Unmarshal([]byte(userJson), &tUser)

	var user = model.User{
		Id:            tUser.ID,
		Name:          tUser.Username,
		FollowCount:   tUser.FollowCount,
		FollowerCount: tUser.FollowerCount,
		//TODO IsFollow 待完善
		IsFollow: true,
	}

	c.JSON(http.StatusOK, UserResponse{
		Response: model.Response{StatusCode: common.OK},
		User:     user,
	})
}
