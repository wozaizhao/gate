package wechat

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"wozaizhao.com/gate/middlewares"
)

func GetConfig(c *gin.Context) {
	url := c.Query("url")
	official := middlewares.GetOfficialAccount()
	js := official.GetJs()
	config, err := js.GetConfig(url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "data": config, "errmsg": ""})
}
