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
	// r.GET("/wechat/captcha", wechat.TencentCaptcha)

	// r.GET("/captcha", controllers.GeetestVerify)        // 极验验证
	// r.POST("/login", controllers.Login) // 使用用户名密码登录
	// r.POST("/loginByPhone", controllers.LoginByPhone)   // 手机号登录
	r.POST("/login", controllers.LoginByOpenID) // 快捷登录

	r.GET("/dict/cities", controllers.GetCities)
	r.GET("/dict/provinces", controllers.GetProvinces)
	r.GET("/dict/regions", controllers.GetRegions)
	r.GET("/dict/allRegions", controllers.GetAllRegions)
	// r.POST("/dict/addCity", controllers.AddCity)
	// r.POST("/dict/addProvince", controllers.AddProvince)
	// r.POST("/dict/addRegion", controllers.AddRegion)

	// 注册用户可以访问 /user
	user := r.Group("/user", controllers.UserAuth())
	{
		user.GET("/currentUser", controllers.CurrentUser) // 当前帐户
		user.PUT("/edit", controllers.UpdateUser)         // 设置用户昵称、头像、性别、用户名、密码
		user.POST("/upload", controllers.Upload)          // 上传
	}

	// 管理员才能访问 /admin
	admin := r.Group("/admin", controllers.AdminAuth())
	{
		admin.GET("/users", controllers.AdminGetUsers)          // 获取所有用户
		admin.PUT("/users/:id", controllers.AdminEditUser)      // 编辑用户状态、角色、积分等
		admin.POST("/users/configRole", controllers.ConfigRole) // 配置用户角色
		admin.POST("/roles", controllers.AddRole)               // 添加角色
		admin.GET("/roles", controllers.GetRoles)               // 获取所有角色
		admin.POST("/menus", controllers.AddMenu)               // 添加菜单
		admin.POST("/features", controllers.AddFeature)         // 添加功能
		// admin.PUT("/roles/:id", controllers.EditRole) // 编辑角色
		admin.POST("/roles/configMenu", controllers.ConfigRoleMenu)       // 配置角色菜单
		admin.POST("/roles/configFeature", controllers.ConfigRoleFeature) // 配置角色功能
	}

	return r
}
