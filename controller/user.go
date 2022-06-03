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

type UserResponse struct {
	model.Response
	User model.User `json:"user"`
}

func UserInfo(c *gin.Context) {
	token := c.Query("token")
	userId := c.Query("user_id")
	//fmt.Println(otherId)

	userJson := common.Rdb.Get(c, common.UserLoginPrefix+token).Val()
	//if userJson == "" {
	//	c.JSON(http.StatusOK, UserResponse{
	//		Response: model.Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
	//	})
	//}

	var tUser model.TUser
	json.Unmarshal([]byte(userJson), &tUser)
	uId, _ := strconv.ParseInt(userId, 10, 64)
	user := service.GetUserInfo(c, uId, tUser.ID)

	//var user = model.User{
	//	Id:            tUser.ID,
	//	Name:          tUser.Username,
	//	FollowCount:   tUser.FollowCount,
	//	FollowerCount: tUser.FollowerCount,
	//	//TODO IsFollow 待完善
	//	IsFollow: true,
	//}

	c.JSON(http.StatusOK, UserResponse{
		Response: model.Response{StatusCode: common.OK},
		User:     user,
	})
}
