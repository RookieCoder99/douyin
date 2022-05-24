package middleware

import (
	"douyin/common"
	"douyin/model"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Authority() func(c *gin.Context) {
	return func(c *gin.Context) {
		uuid := c.Query("token")
		userJson := common.Rdb.Get(c, common.UserLoginPrefix+uuid).Val()
		if userJson == "" {
			c.JSON(http.StatusOK, model.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
			c.Abort()
		}
		//dec := json.NewDecoder(strings.NewReader(userJson))\
		var m model.TUser
		json.Unmarshal([]byte(userJson), &m)
		c.Next()
	}
}
