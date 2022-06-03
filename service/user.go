package service

import (
	"douyin/common"
	"douyin/dao"
	"douyin/model"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetUserInfo(c *gin.Context, userId int64, currentUserId int64) model.User {
	user := dao.GetUserById(userId)
	followCount, _ := strconv.ParseInt(common.Rdb.Get(c, common.UserFollowCountPrefix+strconv.FormatInt(userId, 10)).Val(), 10, 64)
	followerCount, _ := strconv.ParseInt(common.Rdb.Get(c, common.UserFollowerCountPrefix+strconv.FormatInt(userId, 10)).Val(), 10, 64)
	user.FollowCount = followCount
	user.FollowerCount = followerCount
	user.IsFollow = dao.CheckRelation(userId, currentUserId)
	return user
}
