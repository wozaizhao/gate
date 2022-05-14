package models

import (
	"time"

	"gorm.io/gorm"
	// "wozaizhao.com/gate/common"
)

// 图片
type FeImage struct {
	ID         uint           `json:"id" gorm:"primaryKey"`
	CreatedAt  time.Time      `json:"createdAt"`
	UpdatedAt  time.Time      `json:"updatedAt"`
	DeletedAt  gorm.DeletedAt `gorm:"index"`
	ImageName  string         `json:"imageName" gorm:"type:varchar(20);NOT NULL;comment:图片名称"`
	ImageURL   string         `json:"imageURL" gorm:"type:varchar(20);NOT NULL;comment:图片URL"`
	ImageDesc  string         `json:"imageDesc" gorm:"type:varchar(20);NOT NULL;comment:图片描述"`
	ResourceID uint           `json:"resourceID" gorm:"type:int(10);NOT NULL;comment:资源ID"`
}
