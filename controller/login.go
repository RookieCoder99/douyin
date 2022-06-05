package controller

import (
	"douyin/common"
	"douyin/model"
	"douyin/service"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type LoginRequest struct {
	Username string
	Password string
}

type LoginResponse struct {
	model.Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}
type RegisterRequest struct {
	Username string `json: "username"`
	Password string `json: "password"`
	//Nickname string
}

type RegisterResponse struct {
	model.Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	log.Printf("username=%v, password = %v", username, password)
	tUser := service.Login(username, password)
	if tUser == nil {
		log.Printf("%v 登录密码错误", username)
		common.Resp0(c, common.WrongPassword, common.RespMsg[common.WrongPassword])
		return
	}
	tUser.Password = ""
	tJson, _ := json.Marshal(&tUser)
	tStr := string(tJson)
	common.Rdb.Set(c, common.UserLoginPrefix+tUser.Token, tStr, 0)

	//common.Resp2(c, common.OK, common.RespMsg[common.OK], "user_id", tUser.ID, "token", tUser.Token)
	c.JSON(http.StatusOK, LoginResponse{
		Response: model.Response{StatusCode: common.OK},
		UserId:   tUser.ID,
		Token:    tUser.Token,
	})
	return
}

func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	log.Printf("username=%v, password = %v", username, password)
	if exist := service.IsUsernameExist(username); exist {
		log.Printf("用户名%v已经存在", username)
		c.JSON(http.StatusOK, model.Response{
			StatusCode: common.UserHasExisted,
			StatusMsg:  common.RespMsg[common.UserHasExisted],
		})
		return
	}

	tUser := service.Register(username, password)

	tUser.Password = ""
	tJson, _ := json.Marshal(&tUser)
	tStr := string(tJson)
	common.Rdb.Set(c, common.UserLoginPrefix+tUser.Token, tStr, 0)

	common.Rdb.Set(c, common.UserFollowCountPrefix+strconv.FormatInt(tUser.ID, 10), 0, 0)
	common.Rdb.Set(c, common.UserFollowerCountPrefix+strconv.FormatInt(tUser.ID, 10), 0, 0)

	log.Printf("用户名%v注册成功", username)
	//common.Resp2(c,common.OK, common.RespMsg[common.OK], "user_id" )
	c.JSON(http.StatusOK, RegisterResponse{
		Response: model.Response{StatusCode: common.OK, StatusMsg: common.RespMsg[common.OK]},
		UserId:   tUser.ID,
		Token:    tUser.Token,
	})
	return
}
