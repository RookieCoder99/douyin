package service

import (
	"douyin/dao"
	"douyin/model"
)

func GetCommentCount(videoId int64) *int64 {
	count := dao.GetCommentCountByVideoId(videoId)
	return count
}
func CreateComment(userId int64, videoId int64, commentText string) bool {
	var comment = model.TComment{
		VideoID:     videoId,
		CommentText: commentText,
		UserID:      userId,
	}
	return dao.InsertComment(&comment)
}

func DeleteComment(commentId int64) bool {
	return dao.DeleteComment(commentId)
}

func GetCommentList(videoId int64) []model.Comment {
	return dao.GetCommentByVideoId(videoId)
}
