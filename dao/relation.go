package dao

import (
	"douyin/common"
	"douyin/model"
	"log"
)

func AddRelation(relation *model.TRelation) bool {
	res := common.Db.Create(relation)
	if res.Error != nil {
		log.Println(res.Error.Error())
		return false
	}
	return true
}
func DeleteRelation(userId, toUserId int64) bool {
	res := common.Db.Where("follower_id = ? and follow_id = ?", userId, toUserId).Delete(&model.TRelation{})
	if res.Error != nil {
		log.Println(res.Error.Error())
		return false
	}
	return true
}

func GetFollowListByUserId(userId int64) []model.User {
	var users []model.User
	//subQuery := common.Db.Table("t_relation").Where("follower_id = ? and follow_id = t_user.id")
	//res1 := common.Db.Table("t_relation ").
	//	Select("t_user.id, t_user.username as name, t_user.follow_count, t_user.follower_count,  "+
	//		"").
	//	Joins("left join t_user tu on t_relation.follow_id = t_user.id ").
	//	Where("t_relation.follower_id =?", userId).
	//	Find(&users)
	res1 := common.Db.Table("t_relation ").
		Select("tu.id, tu.username as name, tu.follow_count, tu.follower_count, true as is_follow ").
		Joins("left join t_user tu on t_relation.follower_id = tu.id ").
		Where("t_relation.follower_id =?", userId).
		Find(&users)

	if res1.Error != nil {
		log.Println(res1.Error.Error())
		return nil
	}
	return users
}

func GetFollowerListByUserId(userId int64) []model.User {
	var users []model.User
	res1 := common.Db.Raw("select t_user.id, t_user.username as name, t_user.follow_count, t_user.follower_count,"+
		"(select count(*) from t_relation where follower_id = ? and follow_id = t_user.id)>0 as is_follow "+
		" from t_relation "+
		"left join  t_user  on t_relation.follower_id = t_user.id "+
		"where t_relation.follow_id = ?", userId, userId).Find(&users)

	if res1.Error != nil {
		log.Println(res1.Error.Error())
		return nil
	}
	return users
}
