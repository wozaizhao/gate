package wechat

import (
	"github.com/gin-gonic/gin"
	// log "github.com/sirupsen/logrus"
	"net/http"
	"wozaizhao.com/gate/middlewares"
)

type code2SessionReq struct {
	Code string `json:"code"`
}

func Code2Session(c *gin.Context) {
	var req code2SessionReq
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	mini := middlewares.GetMiniProgram()
	auth := mini.GetAuth()
	res, err := auth.Code2Session(req.Code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "data": res, "errmsg": ""})
}

// type userInfoReq struct {
// 	SessionKey    string `json:"sessionKey"`
// 	EncryptedData string `json:"encryptedData"`
// 	IV            string `json:"iv"`
// }

// func DecryptUserInfo(c *gin.Context) {
// 	var req userInfoReq
// 	if err := c.BindJSON(&req); err != nil {
// 		c.JSON(http.StatusBadRequest, err.Error())
// 		return
// 	}
// 	mini := middlewares.GetMiniProgram()
// 	encryptor := mini.GetEncryptor()
// 	plainData, err := encryptor.Decrypt(req.SessionKey, req.EncryptedData, req.IV)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, err.Error())
// 		return
// 	}
// 	log.Debugf("plainData=%+v", plainData)
// 	c.JSON(http.StatusOK, gin.H{"code": 200, "data": plainData, "errmsg": ""})
// }
