package dao

import (
	"douyin/common"
	"douyin/model"
	"log"
)

func GetCommentCountByAuthorId(authorId int64) *int64 {
	var count int64
	common.Db.Model(&model.TComment{}).Where("author_id = ?", authorId).Count(&count)
	return &count
}
func GetCommentCountByVideoId(videoId int64) *int64 {
	var count int64
	common.Db.Model(&model.TComment{}).Where("video_id = ?", videoId).Count(&count)
	return &count
}

func InsertComment(comment *model.TComment) bool {
	res := common.Db.Create(comment)
	if res.Error != nil {
		log.Println(res.Error.Error())
		return false
	}
	return true
}

func DeleteComment(commentId int64) bool {
	res := common.Db.Delete(&model.TComment{}, commentId)
	if res.Error != nil {
		log.Println(res.Error.Error())
		return false
	}
	return true
}

func GetCommentByVideoId(videoId int64) []model.Comment {
	var comments []model.Comment
	res := common.Db.Table("t_comment").
		Select(" t_comment.id,  t_comment.comment_text as content, t_comment.created_at as create_date, "+
			"tu.id as id, tu.username as name, tu.follow_count as follow_count, tu.follower_count as follower_count").
		Joins("left join t_user tu on t_comment.user_id = tu.id ").
		Where("t_comment.video_id = ? ", videoId).
		Find(&comments)
	if res.Error != nil {
		log.Println(res.Error.Error())
		return nil
	}
	return comments
}
