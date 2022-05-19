package models

import (
	"time"

	"gorm.io/gorm"
	"wozaizhao.com/gate/common"
)

// 生态
type FeEcosystem struct {
	ID            uint           `json:"id" gorm:"primaryKey"`
	CreatedAt     time.Time      `json:"createdAt"`
	UpdatedAt     time.Time      `json:"updatedAt"`
	DeletedAt     gorm.DeletedAt `json:"deletedAt" gorm:"index"`
	EcosystemName string         `json:"ecosystemName" gorm:"type:varchar(30);NOT NULL;comment:生态名称"`
	EcosystemDesc string         `json:"ecosystemDesc" gorm:"type:varchar(255);NOT NULL;comment:生态描述"`
	Status        uint           `json:"status" gorm:"type:tinyint(1);NOT NULL;DEFAULT '0';comment:状态"`
}

func CreateFeEcosystem(ecosystemName, ecosystemDesc string) (ecosystem FeEcosystem, err error) {
	ecosystem = FeEcosystem{EcosystemName: ecosystemName, EcosystemDesc: ecosystemDesc, Status: common.STATUS_NORMAL}
	result := DB.Create(&ecosystem)
	err = result.Error
	return
}

func GetFeEcosystems() (ecosystems []FeEcosystem, err error) {
	err = DB.Where("status = ?", common.STATUS_NORMAL).Find(&ecosystems).Error
	return
}
