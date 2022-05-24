package service

import (
	"douyin/dao"
	"douyin/model"
)

func IsFollow() {

}
func AddFollow(userId, toUserId int64) bool {
	var relation = model.TRelation{
		FollowerId: userId,
		FollowId:   toUserId,
	}
	return dao.AddRelation(&relation)
}

func CancelFollow(userId, toUserId int64) bool {
	return dao.DeleteRelation(userId, toUserId)
}

func GetFollowListByUserId(userId int64) []model.User {
	return dao.GetFollowListByUserId(userId)
}
func GetFollowerListByUserId(userId int64) []model.User {
	return dao.GetFollowerListByUserId(userId)
}
