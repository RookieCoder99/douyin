package controller

import (
	"douyin/common"
	"github.com/gin-gonic/gin"
)

func Hello(c *gin.Context) {
	common.Resp0(c, common.OK, common.RespMsg[common.OK])
}
