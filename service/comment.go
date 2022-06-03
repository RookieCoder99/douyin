package service

import (
	"douyin/dao"
	"douyin/model"
	"log"
)

func GetCommentCount(videoId int64) *int64 {
	count := dao.GetCommentCountByVideoId(videoId)
	return count
}
func CreateAndReturnComment(userId int64, videoId int64, commentText string) *model.Comment {
	var comment = model.TComment{
		VideoID:     videoId,
		CommentText: commentText,
		UserID:      userId,
	}
	res := dao.InsertComment(&comment)
	if !res {
		log.Println("插入评论失败")
		return nil
	}

	return dao.GetCommentByCommentId(comment.ID)
}

func CreateComment(userId int64, videoId int64, commentText string) bool {
	var comment = model.TComment{
		VideoID:     videoId,
		CommentText: commentText,
		UserID:      userId,
	}
	res := dao.InsertComment(&comment)
	if !res {
		log.Println("插入评论失败")
		return false
	}

	return true
}

func DeleteComment(commentId int64) bool {
	return dao.DeleteComment(commentId)
}

func GetCommentList(videoId int64) []model.Comment {
	return dao.GetCommentByVideoId(videoId)
}

func GetComment(commentId int64) *model.Comment {
	return dao.GetCommentByCommentId(commentId)
}
