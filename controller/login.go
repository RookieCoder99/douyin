package controller

import (
	"douyin/common"
	"douyin/model"
	"github.com/gin-gonic/gin"
	"log"
)

func Login(c *gin.Context) {
	var json model.LoginRequest
	if err := c.ShouldBindJSON(&json); err != nil {
		common.RespData(c, common.ParamInvalid, common.RespMsg[common.ParamInvalid])
		return
	}
	log.Printf("login请求参数：%v", &json)

}
func Logout(c *gin.Context) {

}
func Register(c *gin.Context) {

}
