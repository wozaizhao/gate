package models

import (
	"gorm.io/gorm"
)

// RoleFeature 角色功能
type RoleFeature struct {
	RoleID    uint `json:"role_id" gorm:"primaryKey;autoIncrement:false;type:int(11);NOT NULL;DEFAULT '0'"`
	FeatureID uint `json:"feature_id" gorm:"primaryKey;autoIncrement:false;type:int(11);NOT NULL;"`
}

// 创建角色功能
func CreateRoleFeature(roleID, featureID uint) (roleFeature RoleFeature, err error) {
	roleFeature = RoleFeature{RoleID: roleID, FeatureID: featureID}
	result := DB.Create(&roleFeature)
	err = result.Error
	return
}

// 查询角色功能
func GetRoleFeature(roleID uint) (roleFeatures []RoleFeature, err error) {
	result := DB.Where("role_id = ?", roleID).Find(&roleFeatures)
	err = result.Error
	return
}

func ConfigRoleFeature(roleID uint, featureIDs []uint) (err error) {
	txErr := DB.Transaction(func(tx *gorm.DB) error {
		// 删除原有功能
		result := DB.Where("role_id = ?", roleID).Delete(RoleFeature{})
		err = result.Error
		if err != nil {
			return err
		}

		// 创建新功能
		var roleFeatures []RoleFeature
		for _, featureID := range featureIDs {
			roleFeatures = append(roleFeatures, RoleFeature{RoleID: roleID, FeatureID: featureID})
		}
		result = DB.Create(&roleFeatures)
		err = result.Error
		if err != nil {
			return err
		}
		return nil
	})
	return txErr
}
