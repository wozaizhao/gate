package models

import (
	"time"

	"gorm.io/gorm"
	"wozaizhao.com/gate/common"
)

// Role 角色
type Role struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	RoleName  string         `json:"role_name" gorm:"type:varchar(20);NOT NULL;comment:角色名称"`
	RoleKey   string         `json:"role_key" gorm:"type:varchar(20);NOT NULL;comment:角色标识"`
	RoleDesc  string         `json:"role_desc" gorm:"type:varchar(50);NOT NULL;comment:角色描述"`
	Status    uint           `json:"status" gorm:"type:tinyint(1);NOT NULL;DEFAULT '0';comment:状态"`
}

type RoleSimple struct {
	RoleName   string `json:"role_name"`
	RoleKey    string `json:"role_key"`
	RoleDesc   string `json:"role_desc"`
	RoleStatus uint   `json:"role_status"`
}

// 创建角色
func CreateRole(roleName, roleKey, roleDesc string) (role Role, err error) {
	role = Role{RoleName: roleName, RoleKey: roleKey, RoleDesc: roleDesc, Status: common.STATUS_NORMAL}
	result := DB.Create(&role)
	err = result.Error
	return
}

// 角色列表
func GetRoles(pageNum, pageSize int) (roles []Role, err error) {
	r := DB.Scopes(Paginate(pageNum, pageSize)).Find(&roles)
	err = r.Error
	return
}
