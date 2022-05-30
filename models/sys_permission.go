package models

import (
	"time"

	"gorm.io/gorm"
	"wozaizhao.com/gate/common"
)

// Permission 权限
type Permission struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"index"`
	Label     string         `json:"label" gorm:"type:varchar(40);not null;unique"`
	Value     string         `json:"value" gorm:"type:varchar(40);not null;unique"`
	ParentID  uint           `json:"parentId" gorm:"type:int(11);not null;default:0"`
	Status    uint           `json:"status" gorm:"type:tinyint(1);not null;default:1"`
}

func CreatePermission(label, value string, parentID uint) error {
	var permission = Permission{Label: label, Value: value, ParentID: parentID, Status: common.STATUS_NORMAL}
	return DB.Create(&permission).Error
}

func GetPermissions() (permissions []Permission, err error) {
	err = DB.Where("status = ?", common.STATUS_NORMAL).Find(&permissions).Error
	return
}

func GetFirstClassPermission() (permissions []Permission, err error) {
	err = DB.Where("status = ? AND parent_id = ?", common.STATUS_NORMAL, 0).Find(&permissions).Error
	return
}
