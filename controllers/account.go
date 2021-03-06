package controllers

import (
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"wozaizhao.com/gate/common"
	"wozaizhao.com/gate/config"
	"wozaizhao.com/gate/middlewares"
	"wozaizhao.com/gate/models"
)

type normalLoginReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	var s normalLoginReq
	if err := c.ShouldBindJSON(&s); err != nil {
		RenderBadRequest(c, err)
		return
	}
	// todo 还需返回用户的角色和菜单
	user, err := models.VerifyUser(s.Username, s.Password)

	if err != nil {
		RenderFail(c, err.Error())
		return
	}

	// todo 判断是否是管理员
	// if common.ADMIN_ROLE == user.Role {
	var token, errorGenerateToken = middlewares.GenerateToken(user.ID, user.Phone)
	if errorGenerateToken != nil {
		RenderError(c, errorGenerateToken)
		return
	}
	roles, err := models.GetUserRole(user.ID)
	if err != nil {
		RenderError(c, err)
		return
	}
	var res loginRes
	res.Token = token
	res.User = basicUserInfo(user, roles)

	RenderSuccess(c, res, "登录成功")
	// return
	// } else {
	// 	RenderFail(c, "不是管理员")
	// }
}

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
	ID        uint                `json:"id"`
	Nickname  string              `json:"nickname"`  // 昵称
	Bio       string              `json:"bio"`       // 简介
	AvatarURL string              `json:"avatarUrl"` // 头像
	Gender    int                 `json:"gender"`    // 性别
	Phone     string              `json:"phone"`     // 手机号
	Username  string              `json:"username"`  // 用户名
	Status    uint                `json:"status"`    // 状态 1 初始值 正常 2 已失效
	Role      []models.RoleSimple `json:"role"`      // 角色 1 初始值 普通用户 2 管理员 3 vip
	// Credit    int    `json:"credit"`    // 用户积分
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
	user, err := models.GetUserByPhone(s.Phone)
	if err != nil {
		RenderError(c, err)
		return
	}
	var token, errorGenerateToken = middlewares.GenerateToken(user.ID, user.Phone)
	if errorGenerateToken != nil {
		RenderError(c, errorGenerateToken)
		return
	}
	roles, err := models.GetUserRole(user.ID)
	if err != nil {
		RenderError(c, err)
		return
	}
	var res loginRes
	res.Token = token
	res.User = basicUserInfo(user, roles)

	RenderSuccess(c, res, "登录成功")
	if s.OpenID != "" {
		models.UpsertUserLoginWithWechat(s.OpenID, user.ID)
	}
}

type shortcutLoginReq struct {
	OpenID string `json:"openID" binding:"required"`
}

func LoginByOpenID(c *gin.Context) {
	var s shortcutLoginReq
	if err := c.ShouldBindJSON(&s); err != nil {
		RenderBadRequest(c, err)
		return
	}
	record, exist, err := models.GetLoginRecordByOpenID(s.OpenID)
	if err != nil {
		RenderError(c, err)
		return
	}
	if !exist {
		RenderFail(c, "")
		return
	}
	// todo 还需返回用户的角色和功能
	user, err := models.GetUserByID(record.UserID)
	if err != nil {
		RenderError(c, err)
		return
	}
	var token, errorGenerateToken = middlewares.GenerateToken(user.ID, user.Phone)
	if errorGenerateToken != nil {
		RenderError(c, errorGenerateToken)
		return
	}
	roles, err := models.GetUserRole(user.ID)
	if err != nil {
		RenderError(c, err)
		return
	}
	var res loginRes
	res.Token = token
	res.User = basicUserInfo(user, roles)

	RenderSuccess(c, res, "")
}

func basicUserInfo(user models.User, roles []models.RoleSimple) (userInfo UserInfo) {
	return UserInfo{ID: user.ID, Nickname: user.Nickname, Bio: user.Bio, AvatarURL: user.AvatarURL, Gender: user.Gender, Phone: user.Phone, Username: user.Username, Status: user.Status, Role: roles}
}

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
	res.User = UserInfo{ID: user.ID, Nickname: user.Nickname, Bio: user.Bio, AvatarURL: user.AvatarURL, Gender: user.Gender, Phone: user.Phone, Username: user.Username, Status: user.Status}

	RenderSuccess(c, res, "")
}

type updateUserReq struct {
	Phone     string `json:"phone"`
	Nickname  string `json:"nickname"`
	AvatarURL string `json:"avatarUrl"`
	Gender    int    `json:"gender"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Bio       string `json:"bio"`
}

func UpdateUser(c *gin.Context) {
	userID := c.MustGet("userID").(uint)
	var s updateUserReq
	if err := c.ShouldBindJSON(&s); err != nil {
		RenderBadRequest(c, err)
		return
	}
	res, err := models.UpdateUser(userID, s.Gender, s.Phone, s.Nickname, s.AvatarURL, s.Username, s.Password, s.Bio)
	if err != nil {
		RenderError(c, err)
		return
	}
	RenderSuccess(c, res, "修改成功")

}

type UserData struct {
	ID        uint                `json:"id"`
	Username  string              `json:"username"`
	Nickname  string              `json:"nickname"`
	AvatarURL string              `json:"avatarUrl"`
	Phone     string              `json:"phone"`
	Gender    int                 `json:"gender"`
	Status    uint                `json:"status"`
	CreatedAt time.Time           `json:"created_at"`
	Bio       string              `json:"bio"`
	Roles     []models.RoleSimple `json:"roles"`
	// RoleNames    string    `json:"role_names"`
	// RoleKeys     string    `json:"role_keys"`
	// RoleStatuses string    `json:"role_statuses"`
}

type usersRes struct {
	List  []UserData `json:"list"`
	Total int64      `json:"total"`
}

func AdminGetUsers(c *gin.Context) {
	pageNumParam := c.DefaultQuery("pageNum", "1")
	pageSizeParam := c.DefaultQuery("pageSize", "10")
	pageNum, _ := common.ParseInt(pageNumParam)
	pageSize, _ := common.ParseInt(pageSizeParam)
	users, err := models.GetUsers(int(pageNum), int(pageSize))
	if err != nil {
		RenderError(c, err)
		return
	}
	vsm := make([]UserData, len(users))
	for i, v := range users {
		var roles []models.RoleSimple
		if len(v.RoleNames) > 0 {
			roleNames := strings.Split(v.RoleNames, ",")
			roleKeys := strings.Split(v.RoleKeys, ",")
			roleStatuses := strings.Split(v.RoleStatuses, ",")
			common.LogDebug("roleNames", len(roleNames))
			roles = make([]models.RoleSimple, len(roleNames))
			for j, role := range roleNames {
				roleStatus, _ := common.ParseInt(roleStatuses[j])
				roles[j] = models.RoleSimple{RoleName: role, RoleKey: roleKeys[j], RoleStatus: uint(roleStatus)}
			}
		}
		vsm[i] = UserData{
			ID:        v.ID,
			Nickname:  v.Nickname,
			Bio:       v.Bio,
			AvatarURL: v.AvatarURL,
			Gender:    v.Gender,
			Phone:     v.Phone,
			Username:  v.Username,
			Status:    v.Status,
			Roles:     roles,
			// Role:      v.Role,
			// Credit:    v.Credit,
		}
	}

	var total = models.GetUserCount()
	var pageCount = float64(total) / float64(pageSize)
	var res = usersRes{
		List:  vsm,
		Total: common.Round(pageCount),
	}
	RenderSuccess(c, res, "")
}

func AdminEditUser(c *gin.Context) {}
