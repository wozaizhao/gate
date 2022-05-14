package models

import (
	"time"

	"gorm.io/gorm"
)

// City 城市
type City struct {
	ID         uint           `json:"id" gorm:"primaryKey"`
	CreatedAt  time.Time      `json:"createdAt"`
	UpdatedAt  time.Time      `json:"updatedAt"`
	DeletedAt  gorm.DeletedAt `gorm:"index"`
	Letter     string         `json:"letter" gorm:"type:varchar(2);NOT NULL;comment:首字母"`
	Name       string         `json:"name" gorm:"type:varchar(12);NOT NULL;comment:城市名称"`
	Code       string         `json:"code" gorm:"type:varchar(6);NOT NULL;comment:城市编码"`
	ParentCode string         `json:"parent_code" gorm:"type:varchar(6);NOT NULL;;comment:父级编码"`
}

// 创建城市
func CreateCity(letter, name, code, parentCode string) (err error) {
	var city = City{Letter: letter, Name: name, Code: code, ParentCode: parentCode}
	result := DB.Create(&city)
	err = result.Error
	return
}

type Result struct {
	Letter     string `json:"letter"`
	Name       string `json:"name"`
	Code       string `json:"code"`
	ParentCode string `json:"parent_code"`
}

// 获取所有城市
func GetCities() (cities []Result, err error) {
	r := DB.Table("cities").Scan(&cities)
	err = r.Error
	return
}
