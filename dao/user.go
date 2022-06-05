package dao

import (
	"douyin/common"
	"douyin/model"
	"fmt"
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

// 当前用户id和 查询的用户id
func GetUserById(userId int64) *model.TUser {
	//type User struct {
	//	Id            int64  `json:"id,omitempty"`
	//	Name          string `json:"name,omitempty"`
	//	FollowCount   int64  `json:"follow_count,omitempty"`
	//	FollowerCount int64  `json:"follower_count,omitempty"`
	//	IsFollow      bool   `json:"is_follow,omitempty"`
	//}
	var user model.TUser
	res := common.Db.Table("t_user").Where("id=?", userId).Find(&user)
	if res.Error != nil {
		fmt.Println(res.Error)
	}
	return &user

}

func CheckRelation(otherId int64, userId int64) bool {
	//type User struct {
	//	Id            int64  `json:"id,omitempty"`
	//	Name          string `json:"name,omitempty"`
	//	FollowCount   int64  `json:"follow_count,omitempty"`
	//	FollowerCount int64  `json:"follower_count,omitempty"`
	//	IsFollow      bool   `json:"is_follow,omitempty"`
	//}
	var relation model.TRelation
	count := common.Db.Table("t_relation").Where("follow_id=? and follower_id = ?", otherId, userId).Find(&relation).RowsAffected
	return count > 0

}
