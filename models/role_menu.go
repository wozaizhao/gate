package models

import (
	"gorm.io/gorm"
)

// RoleMenu 角色菜单
type RoleMenu struct {
	RoleID uint `json:"role_id" gorm:"primaryKey;autoIncrement:false;type:int(11);NOT NULL;DEFAULT '0'"`
	MenuID uint `json:"menu_id" gorm:"primaryKey;autoIncrement:false;type:int(11);NOT NULL;"`
}

// 创建角色菜单
func CreateRoleMenu(roleID, menuID uint) (roleMenu RoleMenu, err error) {
	roleMenu = RoleMenu{RoleID: roleID, MenuID: menuID}
	result := DB.Create(&roleMenu)
	err = result.Error
	return
}

// 查询角色菜单
func GetRoleMenu(roleID uint) (roleMenus []RoleMenu, err error) {
	result := DB.Where("role_id = ?", roleID).Find(&roleMenus)
	err = result.Error
	return
}

func ConfigRoleMenu(roleID uint, menuIDs []uint) (err error) {
	txErr := DB.Transaction(func(tx *gorm.DB) error {
		// 删除原有菜单
		result := DB.Where("role_id = ?", roleID).Delete(RoleMenu{})
		err = result.Error
		if err != nil {
			return err
		}

		// 创建新菜单
		var roleMenus []RoleMenu
		for _, menuID := range menuIDs {
			roleMenus = append(roleMenus, RoleMenu{RoleID: roleID, MenuID: menuID})
		}
		result = DB.Create(&roleMenus)
		err = result.Error
		if err != nil {
			return err
		}
		return nil
	})
	return txErr
}
