package models

import (
	"time"

	"gorm.io/gorm"
	// "wozaizhao.com/gate/common"
)

// 基础
type FeLink struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"index"`
	Title     string         `json:"title" gorm:"type:varchar(50);NOT NULL;comment:标题"`
	URL       string         `json:"url" gorm:"type:varchar(255);NOT NULL;comment:链接"`
	Status    uint           `json:"status" gorm:"type:tinyint(1);NOT NULL;DEFAULT '0';comment:状态"`
}
