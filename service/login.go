package service

import (
	"douyin/dao"
	"douyin/model"
	"douyin/utils"
	"log"
)

func Login(username string, password string) *model.TUser {
	//1、使用用户名 + 密码 查询用户信息 如果存在用户 直接返回
	pwd := utils.EncoderSha256(password)
	user := dao.QueryForLogin(username, pwd)
	return user
}

func IsUsernameExist(username string) bool {
	return dao.IsUsernameExist(username)
}

func Register(username string, password string) *model.TUser {

	var user model.TUser
	user.Username = username
	pwd := utils.EncoderSha256(password)
	user.Password = pwd
	user.Token = utils.GetUUID()
	log.Println("用户token: ", user.Token)
	newUser := dao.InsertUser(&user)
	return newUser
}
