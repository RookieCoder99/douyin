package controller

import (
	"douyin/common"
	"github.com/gin-gonic/gin"
)

func Hello(c *gin.Context) {
	common.RespData(c, common.OK, common.RespMsg[common.OK])
}
