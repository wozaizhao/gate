package models

import (
	"fmt"
	"gorm.io/gorm"
	"wozaizhao.com/gate/common"
)

// UserRole 用户角色
type UserRole struct {
	UserID uint `json:"user_id" gorm:"primaryKey;autoIncrement:false;type:int(11);NOT NULL;"`
	RoleID uint `json:"role_id" gorm:"primaryKey;autoIncrement:false;type:int(11);NOT NULL;DEFAULT '0'"`
}

// 创建用户角色
func CreateUserRole(userID, roleID uint) (userRole UserRole, err error) {
	userRole = UserRole{UserID: userID, RoleID: roleID}
	result := DB.Create(&userRole)
	err = result.Error
	return
}

// 查询用户角色
func GetUserRole(userID uint) (userRoles []RoleSimple, err error) {
	sqlstr := fmt.Sprintf("select ur.user_id, r.id as role_id, r.role_name, r.role_key, r.status as role_status from user_roles as ur left join roles as r on ur.role_id = r.id where user_id = %d and r.status = %d", userID, common.STATUS_NORMAL)
	r := DB.Raw(sqlstr).Scan(&userRoles)
	// result := DB.Where("user_id = ?", userID).Find(&userRoles)
	err = r.Error
	return
}

func ConfigUserRole(userID uint, roleIDs []uint) (err error) {
	txErr := DB.Transaction(func(tx *gorm.DB) error {
		// 删除原有角色
		result := DB.Where("user_id = ?", userID).Delete(UserRole{})
		err = result.Error
		if err != nil {
			return err
		}

		// 创建新角色
		var userRoles []UserRole
		for _, roleID := range roleIDs {
			userRoles = append(userRoles, UserRole{UserID: userID, RoleID: roleID})
		}
		result = DB.Create(&userRoles)
		err = result.Error
		if err != nil {
			return err
		}
		return nil
	})
	return txErr
}
