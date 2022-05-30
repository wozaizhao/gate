package controllers

import (
	"github.com/gin-gonic/gin"
	"wozaizhao.com/gate/common"
	"wozaizhao.com/gate/models"
)

// 新增角色
type AddRoleReq struct {
	RoleName string `json:"roleName" binding:"required"`
	RoleKey  string `json:"roleKey" binding:"required"`
	RoleDesc string `json:"roleDesc" binding:"required"`
}

func AddRole(c *gin.Context) {
	var req AddRoleReq
	if err := c.ShouldBindJSON(&req); err != nil {
		RenderBadRequest(c, err)
		return
	}
	role, err := models.CreateRole(req.RoleName, req.RoleKey, req.RoleDesc)
	if err != nil {
		RenderFail(c, err.Error())
		return
	}
	RenderSuccess(c, role, "添加成功")
}

type rolesRes struct {
	List  []models.Role `json:"list"`
	Total int64         `json:"total"`
}

// 查询角色
func GetRoles(c *gin.Context) {
	pageNumParam := c.DefaultQuery("pageNum", "1")
	pageSizeParam := c.DefaultQuery("pageSize", "10")
	pageNum, _ := common.ParseInt(pageNumParam)
	pageSize, _ := common.ParseInt(pageSizeParam)
	roles, err := models.GetRoles(int(pageNum), int(pageSize))
	if err != nil {
		RenderFail(c, err.Error())
		return
	}
	var total = models.GetUserCount()
	rolesRes := rolesRes{
		List:  roles,
		Total: GetTotal(int64(total), int64(pageSize)),
	}
	RenderSuccess(c, rolesRes, "查询成功")
}

// 新增菜单
type AddMenuReq struct {
	Name        string `json:"name" binding:"required"`
	Path        string `json:"path" binding:"required"`
	Redirect    string `json:"redirect"`
	Title       string `json:"title" binding:"required"`
	Icon        string `json:"icon"`
	Component   string `json:"component"`
	Permissions string `json:"permissions"`
	ExternalURL string `json:"externalURL"`
	Sort        int    `json:"sort"`
	ParentID    int    `json:"parentId"`
}

func AddMenu(c *gin.Context) {
	var req AddMenuReq
	if err := c.ShouldBindJSON(&req); err != nil {
		RenderBadRequest(c, err)
		return
	}
	menu, err := models.CreateMenu(req.Name, req.Path, req.Redirect, req.Title, req.Icon, req.Component, req.Permissions, req.ExternalURL, req.Sort, uint(req.ParentID))
	if err != nil {
		RenderFail(c, err.Error())
		return
	}
	RenderSuccess(c, menu, "添加成功")
}

func GetMenus(c *gin.Context) {
	menus, err := models.GetMenus()
	if err != nil {
		RenderFail(c, err.Error())
		return
	}
	RenderSuccess(c, gin.H{"list": menus}, "查询成功")
}

type addPermissionReq struct {
	Label    string `json:"label" binding:"required"`
	Value    string `json:"value" binding:"required"`
	ParentID int    `json:"parentId"`
}

func AddPermission(c *gin.Context) {
	var req addPermissionReq
	if err := c.ShouldBindJSON(&req); err != nil {
		RenderBadRequest(c, err)
		return
	}
	err := models.CreatePermission(req.Label, req.Value, uint(req.ParentID))
	if err != nil {
		RenderFail(c, err.Error())
		return
	}
	RenderSuccess(c, nil, "添加成功")
}

type permissionNode struct {
	ID       uint              `json:"id"`
	Label    string            `json:"label"`
	Value    string            `json:"value"`
	Children []*permissionNode `json:"children"`
}

func GetPermissions(c *gin.Context) {
	firstClasses, err := models.GetFirstClassPermission()
	if err != nil {
		RenderFail(c, err.Error())
		return
	}
	all, err := models.GetPermissions()
	if err != nil {
		RenderFail(c, err.Error())
		return
	}
	tree := make([]*permissionNode, 0)

	for _, firstClass := range firstClasses {
		node := &permissionNode{
			ID:       firstClass.ID,
			Label:    firstClass.Label,
			Value:    firstClass.Value,
			Children: make([]*permissionNode, 0),
		}
		firstChildren := make([]*permissionNode, 0)
		for _, permission := range all {
			if permission.ParentID == firstClass.ID {
				var child = &permissionNode{
					ID:       permission.ID,
					Label:    permission.Label,
					Value:    permission.Value,
					Children: make([]*permissionNode, 0),
				}
				firstChildren = append(firstChildren, child)
			}
		}
		if len(firstChildren) > 0 {
			node.Children = firstChildren
		}
		tree = append(tree, node)
	}
	// permissions, err := models.GetPermissions()
	// if err != nil {
	// 	RenderFail(c, err.Error())
	// 	return
	// }
	RenderSuccess(c, tree, "")
}

// 新增功能
type AddFeatureReq struct {
	FeatureName     string `json:"feature_name" binding:"required"`
	FeatureKey      string `json:"feature_key" binding:"required"`
	FeatureDesc     string `json:"feature_desc" binding:"required"`
	FeatureParentID string `json:"feature_parent_id" binding:"required"`
}

func AddFeature(c *gin.Context) {
	var req AddFeatureReq
	if err := c.ShouldBindJSON(&req); err != nil {
		RenderBadRequest(c, err)
		return
	}
	parentID, _ := common.ParseInt(req.FeatureParentID)
	menu, err := models.CreateFeature(req.FeatureName, req.FeatureKey, req.FeatureDesc, uint(parentID))
	if err != nil {
		RenderFail(c, err.Error())
		return
	}
	RenderSuccess(c, menu, "添加成功")
}

type configPermissionReq struct {
	RoleID        int    `json:"roleID" binding:"required"`
	PermissionIDs []uint `json:"permissionIDs" binding:"required"`
}

func ConfigRolePermission(c *gin.Context) {
	var req configPermissionReq
	if err := c.ShouldBindJSON(&req); err != nil {
		RenderBadRequest(c, err)
		return
	}
	err := models.ConfigRolePermission(uint(req.RoleID), req.PermissionIDs)
	if err != nil {
		RenderFail(c, err.Error())
		return
	}
	RenderSuccess(c, nil, "配置成功")
}

// 配置用户角色
// type ConfigRoleReq struct {
// 	UserID  string `json:"user_id" binding:"required"`
// 	RoleIDs []uint `json:"role_ids" binding:"required"`
// }

// func ConfigRole(c *gin.Context) {
// 	var req ConfigRoleReq
// 	if err := c.ShouldBindJSON(&req); err != nil {
// 		RenderBadRequest(c, err)
// 		return
// 	}
// 	userID, _ := common.ParseInt(req.UserID)
// 	err := models.ConfigUserRole(uint(userID), req.RoleIDs)
// 	if err != nil {
// 		RenderFail(c, err.Error())
// 		return
// 	}
// 	RenderSuccess(c, nil, "配置成功")
// }

// 配置角色菜单
// type ConfigRoleMenuReq struct {
// 	RoleID  string `json:"role_id" binding:"required"`
// 	MenuIDs []uint `json:"menu_ids" binding:"required"`
// }

// func ConfigRoleMenu(c *gin.Context) {
// 	var req ConfigRoleMenuReq
// 	if err := c.ShouldBindJSON(&req); err != nil {
// 		RenderBadRequest(c, err)
// 		return
// 	}
// 	roleID, _ := common.ParseInt(req.RoleID)
// 	err := models.ConfigRoleMenu(uint(roleID), req.MenuIDs)
// 	if err != nil {
// 		RenderFail(c, err.Error())
// 		return
// 	}
// 	RenderSuccess(c, nil, "配置成功")
// }

// 配置角色功能
// type ConfigRoleFeatureReq struct {
// 	RoleID     uint   `json:"role_id" binding:"required"`
// 	FeatureIDs []uint `json:"feature_ids" binding:"required"`
// }

// func ConfigRoleFeature(c *gin.Context) {
// 	var req ConfigRoleFeatureReq
// 	if err := c.ShouldBindJSON(&req); err != nil {
// 		RenderBadRequest(c, err)
// 		return
// 	}
// 	err := models.ConfigRoleFeature(req.RoleID, req.FeatureIDs)
// 	if err != nil {
// 		RenderFail(c, err.Error())
// 		return
// 	}
// 	RenderSuccess(c, nil, "配置成功")
// }

// 查询用户角色 优先级低

// 查询角色菜单
// type GetRoleMenuReq struct {
// 	RoleID uint `json:"role_id" binding:"required"`
// }

// func GetRoleMenu(c *gin.Context) {
// 	var req GetRoleMenuReq
// 	if err := c.ShouldBindJSON(&req); err != nil {
// 		RenderBadRequest(c, err)
// 		return
// 	}
// 	menus, err := models.GetRoleMenu(req.RoleID)
// 	if err != nil {
// 		RenderFail(c, err.Error())
// 		return
// 	}
// 	RenderSuccess(c, menus, "查询成功")
// }

// 查询角色功能
