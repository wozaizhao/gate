package models

import (
	"gorm.io/gorm"
	"time"
)

// Province 省级行政区
type Province struct {
	ID         uint           `json:"id" gorm:"primaryKey"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index"`
	Letter     string         `json:"letter" gorm:"type:varchar(2);NOT NULL;comment:首字母"`
	Name       string         `json:"name" gorm:"type:varchar(3);NOT NULL;comment:省级行政区名称"`
	Code       string         `json:"code" gorm:"type:varchar(6);NOT NULL;comment:省级行政区编码"`
	ParentCode string         `json:"parent_code" gorm:"type:varchar(6);NOT NULL;comment:父级编码"`
}

// 创建城市
func CreateProvince(letter, name, code, parentCode string) (err error) {
	var province = Province{Letter: letter, Name: name, Code: code, ParentCode: parentCode}
	result := DB.Create(&province)
	err = result.Error
	return
}

// 获取所有省级行政区
func GetProvinces() (provinces []Result, err error) {
	r := DB.Table("provinces").Scan(&provinces)
	err = r.Error
	return
}
