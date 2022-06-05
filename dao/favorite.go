package dao

import (
	"douyin/common"
	"douyin/model"
	"log"
)

func GetFavoriteCountByAuthorId(userId int64) *int64 {
	var count int64
	common.Db.Model(&model.TFavorite{}).Where("user_id = ?", userId).Count(&count)
	return &count
}

func InsertFavorite(favorite *model.TFavorite) bool {
	res := common.Db.Create(favorite)
	if res.Error != nil {
		log.Println(res.Error.Error())
		return false
	}
	return true
}

func DeleteFavorite(userId, videoId int64) bool {
	res := common.Db.Where("user_id= ? and video_id= ?", userId, videoId).Delete(&model.TFavorite{})
	if res.Error != nil {
		log.Println(res.Error.Error())
		return false
	}
	return true
}

func GetVideosByIds(videoIds []int64) []*model.Video {
	var videos []*model.Video
	res := common.Db.Table(" t_video ").
		Select("t_video.id, t_video.play_url, t_video.cover_url,"+
			"(select count(*) from t_favorite f where f.video_id = t_video.id) favorite_count,"+
			"(select count(*) from t_comment tc where tc.video_id = t_video.id) comment_count,"+
			"true  is_favorite,"+
			"t_user.id as id, t_user.username as name, t_user.follow_count as follow_count, t_user.follower_count as follower_count ").
		Joins("left join t_user  on  t_user.id = t_video.author_id ").
		Where("t_video.id in ?", videoIds).Find(&videos)
	if res.Error != nil {
		log.Println(res.Error.Error())
		return nil
	}
	return videos
}

func GetFavoriteByUserId(userId int64) []*model.Video {
	var videos []*model.Video
	//res := common.Db.Table("t_user").
	//	Select("t_user.id, t_user.username, t_user.follow_count, t_user.follower_count, t_user.is_follow, "+
	//		"tv.id, tv.play_url, tv.cover_url").
	//	Joins("left join t_favorite tf on t_user.id = tf.user_id").
	//	Joins("left join t_video tv on tf.video_id = tv.id").
	//	Where("t_user.id=?", userId).
	//	Scan(&videos)

	res := common.Db.Table("t_user").
		Select("t_video.id, t_video.play_url, t_video.cover_url,"+
			"(select count(*) from t_favorite f where f.video_id = t_video.id) favorite_count,"+
			"(select count(*) from t_comment tc where tc.video_id = t_video.id) comment_count,"+
			"true  is_favorite,"+
			"t_user.id as id, t_user.username as name, t_user.follow_count as follow_count, t_user.follower_count as follower_count ").
		Joins("left join t_favorite tf on tf.user_id = t_user.id ").
		Joins("left join t_video  on t_video.id = tf.video_id ").
		Where("t_user.id=?", userId).Find(&videos)
	//Order("t_video.created_at").Find(&videos)

	// 所有视频都是已点赞的状态
	//for _, video := range videos {
	//	video.IsFavorite = true
	//}
	if res.Error != nil {
		log.Println(res.Error.Error())
		return nil
	}
	return videos
}
