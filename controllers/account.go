package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"wozaizhao.com/gate/config"
	"wozaizhao.com/gate/middlewares"
	"wozaizhao.com/gate/models"
)

func Login(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": 200, "errMsg": ""})
}

func LoginWeapp(c *gin.Context) {}

type loginReq struct {
	Phone  string `json:"phone" binding:"required"`
	Code   string `json:"code" binding:"required"`
	OpenID string `json:"openID"`
}

type loginRes struct {
	User  UserInfo `json:"user"`
	Token string   `json:"token"`
}

type UserInfo struct {
	ID        uint   `json:"id"`
	NickName  string `json:"nickname"`   // 昵称
	Bio       string `json:"bio"`        // 简介
	AvatarURL string `json:"avatar_url"` // 头像
	Gender    int    `json:"gender"`     // 性别
	Phone     string `json:"phone"`      // 手机号
	Username  string `json:"username"`   // 用户名
	Status    uint   `json:"status"`     // 状态 1 初始值 正常 2 已失效
	Role      uint   `json:"role"`       // 角色 1 初始值 普通用户 2 管理员 3 vip
	Credit    int    `json:"credit"`     // 用户积分
}

func LoginByPhone(c *gin.Context) {
	cfg := config.GetConfig()

	var s loginReq
	if err := c.ShouldBindJSON(&s); err != nil {
		RenderBadRequest(c, err)
		return
	}

	if cfg.Mode == "production" {
		captchaAvailable := models.CaptchaAvailable(s.Phone, s.Code)

		if !captchaAvailable {
			RenderFail(c, "验证码错误")
			return
		}
	}

	user, err := models.GetUserByPhone(s.Phone, s.OpenID)
	if err != nil {
		RenderError(c, err)
		return
	}
	var token, errorGenerateToken = middlewares.GenerateToken(user.ID, user.Phone)
	if errorGenerateToken != nil {
		RenderError(c, errorGenerateToken)
		return
	}
	var res loginRes
	res.Token = token
	res.User = UserInfo{ID: user.ID, NickName: user.NickName, Bio: user.Bio, AvatarURL: user.AvatarURL, Gender: user.Gender, Phone: user.Phone, Username: user.Username, Status: user.Status, Role: user.Role, Credit: user.Credit}

	RenderSuccess(c, res, "登录成功")
}

func LoginByOpenID(c *gin.Context) {}

func CheckPhoneExist(c *gin.Context) {}

// type captchaRes struct {
// 	Phone string `json:"phone"`
// 	From  string `json:"from"`
// }

func Register(c *gin.Context) {

}

func CurrentUser(c *gin.Context) {
	userID := c.MustGet("userID").(uint)
	user, err := models.GetUserByID(userID)
	if err != nil {
		RenderError(c, err)
		return
	}
	var res loginRes
	res.User = UserInfo{ID: user.ID, NickName: user.NickName, Bio: user.Bio, AvatarURL: user.AvatarURL, Gender: user.Gender, Phone: user.Phone, Username: user.Username, Status: user.Status, Role: user.Role, Credit: user.Credit}

	RenderSuccess(c, res, "")
}

func LinkWechat(c *gin.Context) {}

func UpdateUser(c *gin.Context) {}

func AdminGetUsers(c *gin.Context) {}

func AdminEditUser(c *gin.Context) {}
