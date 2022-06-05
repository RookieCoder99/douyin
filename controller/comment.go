package controller

import (
	"douyin/common"
	"douyin/model"
	"douyin/service"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

var commentPrev = "comment_"

type CommentListResponse struct {
	model.Response
	CommentList []model.Comment `json:"comment_list,omitempty"`
}

type CommentResponse struct {
	model.Response
	Comment model.Comment `json:"comment,omitempty"`
}

func CommentAction(c *gin.Context) {
	token := c.Query("token")
	videoId := c.Query("video_id")
	actionType := c.Query("action_type")
	commentText := c.Query("comment_text")
	commentId := c.Query("comment_id")
	userJson := common.Rdb.Get(c, common.UserLoginPrefix+token).Val()
	if userJson == "" {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: common.UserNotExisted,
			StatusMsg:  common.RespMsg[common.UserNotExisted],
		})
		return
	}
	var tUser model.TUser
	json.Unmarshal([]byte(userJson), &tUser)
	vId, err := strconv.ParseInt(videoId, 10, 64)

	if err != nil {
		log.Println("videoId params error!")
		log.Println(err)
		c.JSON(http.StatusOK, model.Response{
			StatusCode: common.ParamInvalid,
			StatusMsg:  common.RespMsg[common.ParamInvalid],
		})
		return
	}
	if actionType == "1" {
		// 添加评论
		comment := service.CreateAndReturnComment(tUser.ID, vId, commentText)
		log.Println(comment)
		if comment == nil {
			c.JSON(http.StatusOK, model.Response{
				StatusCode: common.CommentActionError,
				StatusMsg:  common.RespMsg[common.CommentActionError],
			})
			return
		}

		// 添加评论到redis中
		commentJson, _ := json.Marshal(comment)
		common.Rdb.SAdd(c, commentPrev+videoId, commentJson)

		c.JSON(http.StatusOK, CommentResponse{
			Response: model.Response{
				StatusCode: common.OK,
				StatusMsg:  common.RespMsg[common.OK],
			},
			Comment: *comment,
		})
	} else if actionType == "2" {
		// 删除评论
		cId, _ := strconv.ParseInt(commentId, 10, 64)

		// 在mysql中查询该条评论
		// TODO: 这里需要经过两次数据库操作
		comment := service.GetComment(cId)
		res := service.DeleteComment(cId)
		if !res {

			c.JSON(http.StatusOK, model.Response{
				StatusCode: common.CommentDeleteError,
				StatusMsg:  common.RespMsg[common.CommentDeleteError],
			})
			return
		}

		// 在redis中删除该评论
		commentJson, _ := json.Marshal(comment)
		common.Rdb.SRem(c, commentPrev+videoId, commentJson)
		c.JSON(http.StatusOK, model.Response{
			StatusCode: common.OK,
			StatusMsg:  common.RespMsg[common.OK],
		})
	}

}

func CommentList(c *gin.Context) {
	token := c.Query("token")
	videoId := c.Query("video_id")
	//userId := c.Query("user_id")
	//fmt.Println(userId)
	userJson := common.Rdb.Get(c, common.UserLoginPrefix+token).Val()
	if userJson == "" {
		c.JSON(http.StatusOK, model.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}
	var tUser model.TUser
	json.Unmarshal([]byte(userJson), &tUser)
	vId, _ := strconv.ParseInt(videoId, 10, 64)

	// 在redis中查找是否有该视频的评论
	commentCnt := common.Rdb.SCard(c, commentPrev+videoId).Val()
	var commentList []model.Comment
	if commentCnt != 0 {
		// 说明有数据
		log.Println("get comment list from redis")
		log.Print(commentPrev + videoId)
		comments := common.Rdb.SMembers(c, commentPrev+videoId).Val()
		var comment model.Comment
		for _, commentJson := range comments {
			json.Unmarshal([]byte(commentJson), &comment)
			commentList = append(commentList, comment)
		}
	} else {
		// 如果没有，则到mysql中拿
		log.Println("get comment list from mysql")
		commentList = service.GetCommentList(vId)

		// 将评论存储到redis中
		for _, comment := range commentList {
			commentJson, _ := json.Marshal(comment)
			common.Rdb.SAdd(c, commentPrev+videoId, commentJson)
		}
	}

	// 否则，直接返回
	c.JSON(http.StatusOK, CommentListResponse{
		Response: model.Response{
			StatusCode: common.OK,
			StatusMsg:  common.RespMsg[common.OK],
		},
		CommentList: commentList,
	})
}
