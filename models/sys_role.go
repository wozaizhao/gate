package models

import (
	"time"

	"gorm.io/gorm"
	"wozaizhao.com/gate/common"
)

// Role 角色
type Role struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `json:"deletedAt" gorm:"index"`
	RoleName    string         `json:"roleName" gorm:"type:varchar(20);NOT NULL;comment:角色名称"`
	RoleKey     string         `json:"roleKey" gorm:"type:varchar(20);unique;NOT NULL;comment:角色标识"`
	RoleDesc    string         `json:"roleDesc" gorm:"type:varchar(50);NOT NULL;comment:角色描述"`
	Permissions []Permission   `json:"permissions" gorm:"many2many:role_permission;"`
	Features    []Feature      `json:"features" gorm:"many2many:role_feature;"`
	Status      uint           `json:"status" gorm:"type:tinyint(1);NOT NULL;DEFAULT '0';comment:状态"`
}

// type RoleSimple struct {
// 	RoleName   string `json:"role_name"`
// 	RoleKey    string `json:"role_key"`
// 	RoleDesc   string `json:"role_desc"`
// 	RoleStatus uint   `json:"role_status"`
// }

// 创建角色
func CreateRole(roleName, roleKey, roleDesc string) (role Role, err error) {
	role = Role{RoleName: roleName, RoleKey: roleKey, RoleDesc: roleDesc, Status: common.STATUS_NORMAL}
	result := DB.Create(&role)
	err = result.Error
	return
}

// 角色列表
func GetRoles(pageNum, pageSize int) (roles []Role, err error) {
	r := DB.Scopes(Paginate(pageNum, pageSize)).Preload("Permissions").Preload("Features").Find(&roles)
	err = r.Error
	return
}

func ConfigRolePermission(roleID uint, permissionIDs []uint) (err error) {
	var role Role
	role.ID = roleID
	var permissions []Permission
	txErr := DB.Transaction(func(tx *gorm.DB) error {
		err := tx.Preload("Permissions").First(&role).Error
		if err != nil {
			return err
		}
		err = tx.Model(&role).Association("Permissions").Delete(role.Permissions)
		if err != nil {
			return err
		}
		err = tx.Where("id in (?)", permissionIDs).Find(&permissions).Error
		if err != nil {
			return err
		}
		err = tx.Model(&role).Association("Permissions").Append(permissions)
		if err != nil {
			return err
		}
		return nil
	})
	return txErr
}

func GetRolesPermissions(roles []Role) (permissions []Permission, err error) {
	err = DB.Preload("Permissions").Find(&roles).Error
	if err != nil {
		return nil, err
	}
	for _, role := range roles {
		permissions = append(permissions, role.Permissions...)
	}
	permissions = uniquePermissions(permissions)
	return permissions, nil
}

func uniquePermissions(permissions []Permission) (uniquePermissions []Permission) {
	m := make(map[uint]Permission)
	for _, permission := range permissions {
		m[permission.ID] = permission
	}
	for _, permission := range m {
		uniquePermissions = append(uniquePermissions, permission)
	}
	return
}
