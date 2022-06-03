package dao

import (
	"douyin/common"
	"douyin/model"
	"log"
	"time"
)

//根据用户名和密码查询
func InsertVideo(video *model.TVideo) *model.TVideo {
	res := common.Db.Create(video)
	if res.Error != nil {
		log.Println(res.Error.Error())
		return nil
	}
	var m model.TVideo
	res = common.Db.First(&m, "play_url = ?", video.PlayUrl)
	if res.Error != nil {
		log.Printf("select failed, err is %s", res.Error)
		return nil
	}
	return &m
}
func GetVideoListByUserId(userId int64) []model.Video {
	var videos []model.Video
	res1 := common.Db.Table("t_video").
		Select("t_video.id, t_video.play_url, t_video.cover_url, t_video.title, "+
			"(select count(*) from t_favorite f where f.video_id = t_video.id) favorite_count,"+
			"(select count(*) from t_comment tc where tc.video_id = t_video.id) comment_count,"+
			"(select count(*) from t_favorite tf where tf.user_id = tu.id and tf.video_id = t_video.id)>0  is_favorite, "+
			"tu.id as id, tu.username as name, tu.follow_count as follow_count, tu.follower_count as follower_count ").
		Joins("left join t_user tu on t_video.author_id = tu.id ").
		Where("tu.id=?", userId).
		Order("t_video.created_at").Limit(30).Find(&videos)

	if res1.Error != nil {
		log.Println(res1.Error.Error())
		return nil
	}
	return videos
}

/**
res := common.Db.Table("t_user").
		Select("t_user.id, t_user.username, t_user.follow_count, t_user.follower_count, t_user.is_follow, "+
			"tv.id, tv.play_url, tv.cover_url").
		Joins("left join t_favorite tf on t_user.id = tf.user_id").
		Joins("left join t_video tv on tf.video_id = tv.id").
		Where("t_user.id=?", userId).
		Scan(&videos)
*/

func GetVideoList() ([]model.Video, time.Time) {
	var videos []model.Video
	res1 := common.Db.Table("t_video").
		Select("t_video.id, t_video.play_url, t_video.cover_url, t_video.title, " +
			"tu.id as id,  tu.username as name , tu.follow_count as follow_count , tu.follower_count as follower_count," +
			"(select count(*) from t_favorite f where f.video_id = t_video.id)  favorite_count, " +
			"(select count(*) from t_comment tc where tc.video_id = t_video.id)  comment_count, " +
			"(select count(*) from t_favorite tf where tf.user_id = tu.id and tf.video_id = t_video.id)>0  is_favorite ").
		Joins("left join t_user tu on t_video.author_id = tu.id").
		Order("t_video.created_at").
		Find(&videos)
	var v model.TVideo
	res2 := common.Db.Where("id=?", videos[0].Id).First(&v)

	if res1.Error != nil {
		log.Println(res1.Error.Error())
		return nil, time.Now()
	}
	if res2.Error != nil {
		log.Println("查询时间错误", res2.Error.Error())
	}
	return videos, v.CreatedAt
}

func GetVideoListByTime(timeStr string) ([]model.Video, time.Time) {
	var videos []model.Video
	res1 := common.Db.Table("t_video").
		Select("t_video.id, t_video.play_url, t_video.cover_url, tu.id as id,t_video.title,"+
			"tu.username as name , "+
			"tu.follow_count as follow_count , tu.follower_count as follower_count,"+
			"(select count(*) from t_favorite tf where tf.video_id = t_video.id)  favorite_count, "+
			"(select count(*) from t_comment tc where tc.video_id = t_video.id)  comment_count, "+
			"(select count(*) from t_favorite tf where tf.user_id = tu.id and tf.video_id = t_video.id)>0  is_favorite ").
		Joins("left join t_user tu on t_video.author_id = tu.id ").
		Where("t_video.created_at <=? ", timeStr).
		Order("t_video.created_at desc").
		Find(&videos)

	//res1 := common.Db.Raw("select t_video.id, t_video.play_url, t_video.cover_url, t_video.title, " +
	//	"tu.id as id, tu.username as name, tu.follow_count as follow_count, tu.follower_count as follower_count," +
	//	"(select count(*) from t_favorite f where f.video_id = t_video.id)  favorite_count, " +
	//	"(select count(*) from t_comment tc where tc.video_id = t_video.id)  comment_count," +
	//	"(select count(*) from t_favorite tf where tf.user_id = tu.id and tf.video_id = t_video.id)>0  is_favorite " +
	//	"from t_video " +
	//	"left join t_user tu on t_video.author_id = tu.id " +
	//	"where t_video.created_at < '2330-08-22 16:05:14.9998' " +
	//	"order by t_video.created_at " +
	//	"limit 30").Scan(&videos)

	var v model.TVideo
	res2 := common.Db.Where("id=?", videos[0].Id).First(&v)
	if res1.Error != nil {
		log.Println(res1.Error.Error())
		return nil, time.Now()
	}
	if res2.Error != nil {
		log.Println("查询时间错误", res2.Error.Error())
	}
	return videos, v.CreatedAt
}
