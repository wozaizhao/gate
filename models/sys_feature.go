package models

import (
	"time"

	"gorm.io/gorm"
	"wozaizhao.com/gate/common"
)

// Feature 前台功能
type Feature struct {
	ID              uint           `json:"id" gorm:"primaryKey"`
	CreatedAt       time.Time      `json:"createdAt"`
	UpdatedAt       time.Time      `json:"updatedAt"`
	DeletedAt       gorm.DeletedAt `json:"deletedAt" gorm:"index"`
	FeatureName     string         `json:"feature_name" gorm:"type:varchar(20);NOT NULL;comment:功能名称"`
	FeatureKey      string         `json:"feature_key" gorm:"unique;type:varchar(20);NOT NULL;comment:功能标识"`
	FeatureDesc     string         `json:"feature_desc" gorm:"type:varchar(50);NOT NULL;comment:功能描述"`
	FeatureParentID uint           `json:"feature_parent_id" gorm:"type:int(11);NOT NULL;DEFAULT '0';comment:父级功能ID"`
	Status          uint           `json:"status" gorm:"type:tinyint(1);NOT NULL;DEFAULT '0';comment:状态"`
}

// 创建功能
func CreateFeature(featureName, featureKey, featureDesc string, featureParentID uint) (feature Feature, err error) {
	feature = Feature{FeatureName: featureName, FeatureKey: featureKey, FeatureDesc: featureDesc, FeatureParentID: featureParentID, Status: common.STATUS_NORMAL}
	result := DB.Create(&feature)
	err = result.Error
	return
}

// 查询功能树
func GetFeatureTree() (features []Feature, err error) {
	err = DB.Where("status = ?", common.STATUS_NORMAL).Find(&features).Error
	return
}
