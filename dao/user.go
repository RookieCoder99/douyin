package dao

import (
	"douyin/common"
	"douyin/model"
	"log"
)

// 查询用户名是否存在
func IsUsernameExist(username string) bool {
	var user model.TUser
	result := common.Db.Where("username=?", username).First(&user)
	if result.RowsAffected >= 1 {
		return true
	}
	return false
}

//根据用户名和密码查询
func QueryForLogin(username string, password string) *model.TUser {
	var user model.TUser
	res := common.Db.Where(" username = ? and password = ? ", username, password).First(&user)
	if res.Error != nil {
		log.Println(res.Error.Error())
		return nil
	}
	return &user
}

func InsertUser(user *model.TUser) *model.TUser {
	res := common.Db.Create(&user)
	if res.Error != nil {
		log.Println(res.Error.Error())
		return nil
	}

	var m model.TUser
	res = common.Db.First(&m, "username = ?", user.Username)
	if res.Error != nil {
		log.Printf("select failed, err is %s", res.Error)
	}
	//mid := strconv.FormatInt(m.ID, 10)
	return &m
}
