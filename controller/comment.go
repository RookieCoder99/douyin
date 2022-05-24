package controller

import (
	"douyin/common"
	"douyin/model"
	"douyin/service"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type CommentListResponse struct {
	model.Response
	CommentList []model.Comment `json:"comment_list,omitempty"`
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
	vId, _ := strconv.ParseInt(videoId, 10, 64)
	if actionType == "1" {
		res := service.CreateComment(tUser.ID, vId, commentText)
		if !res {
			c.JSON(http.StatusOK, model.Response{
				StatusCode: common.CommentActionError,
				StatusMsg:  common.RespMsg[common.CommentActionError],
			})
			return
		}
		c.JSON(http.StatusOK, model.Response{
			StatusCode: common.OK,
			StatusMsg:  common.RespMsg[common.OK],
		})
	} else if actionType == "2" {
		cId, _ := strconv.ParseInt(commentId, 10, 64)
		res := service.DeleteComment(cId)
		if !res {
			c.JSON(http.StatusOK, model.Response{
				StatusCode: common.CommentDeleteError,
				StatusMsg:  common.RespMsg[common.CommentDeleteError],
			})
			return
		}
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
	commentList := service.GetCommentList(vId)
	c.JSON(http.StatusOK, CommentListResponse{
		Response: model.Response{
			StatusCode: common.OK,
			StatusMsg:  common.RespMsg[common.OK],
		},
		CommentList: commentList,
	})
}
