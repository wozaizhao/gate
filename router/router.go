package router

import (
	"github.com/gin-gonic/gin"
	"wozaizhao.com/gate/controllers"
	"wozaizhao.com/gate/controllers/wechat"
)

// SetupRouter 设置路由
func SetupRouter() *gin.Engine {
	r := gin.Default()
	// r.SetTrustedProxies([]string{"192.168.1.2"})

	r.MaxMultipartMemory = 8 << 20 // 8 MiB
	// allows all origins
	// r.Use(cors.Default())

	// 任何人都可访问
	r.POST("/wechat/code2Session", wechat.Code2Session)
	// r.POST("/wechat/decryptUserInfo", wechat.DecryptUserInfo)
	r.GET("/wechat/getConfig", wechat.GetConfig)

	r.GET("/captcha", controllers.GeetestVerify)        // 极验验证
	r.POST("/login", controllers.Login)                 // 使用用户名密码登录
	r.POST("/loginByPhone", controllers.LoginByPhone)   // 手机号登录
	r.POST("/shortcutLogin", controllers.LoginByOpenID) // 快捷登录
	// r.GET("/checkPhone", controllers.CheckPhoneExist)   // 检测手机号
	// r.POST("/register", controllers.Register)            // 注册
	// r.GET("/captcha", controllers.GetCaptcha)            // 获取验证码

	// 注册用户可以访问 /user
	user := r.Group("/user", controllers.UserAuth())
	{
		user.GET("/currentUser", controllers.CurrentUser) // 当前帐户
		user.POST("/linkwechat", controllers.LinkWechat)  // 关联openID
		user.PUT("/edit", controllers.UpdateUser)         // 设置用户昵称、头像、性别、用户名、密码
		user.POST("/upload", controllers.Upload)          // 上传
	}

	// 管理员才能访问 /admin
	admin := r.Group("/admin", controllers.AdminAuth())
	{
		admin.GET("/users", controllers.AdminGetUsers)     // 获取所有用户
		admin.PUT("/users/:id", controllers.AdminEditUser) // 编辑用户状态、角色、积分等
	}

	return r
}
